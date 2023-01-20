package main

import (
    bloom "Demiourgos/Blooming"
    p "Firefly-APD"
    mvx "MvxApiScanner"
    sm "SuperMath"
    "flag"
    "fmt"
    "strings"
)

func ReadOmniscientSnakeSnapshot() []mvx.BalanceSFT {
    var (
	MainUnit mvx.BalanceSFT
	Output   []mvx.BalanceSFT
    )
    Path := "_SNAKE-Snapshots\\DEMIOU-6d1b5c-top-holders-omniscient-tools.csv"

    ProcessLine := func(Line string) mvx.BalanceSFT {
	var (
	    Unit mvx.BalanceSFT
	)
	LineString := strings.ReplaceAll(Line, "\"", "")
	SeparatedLineStringSlice := strings.Split(LineString, ",")
	Unit.Address = mvx.MvxAddress(SeparatedLineStringSlice[0])
	Unit.Balance = SeparatedLineStringSlice[1]
	return Unit
    }

    StringSlice := mvx.ReadFile(Path)
    for i := 0; i < len(StringSlice); i++ {
	MainUnit = ProcessLine(StringSlice[i])
	Output = append(Output, MainUnit)
    }
    return Output
}

func MakeSnakeChain() ([]mvx.BalanceSFT, *p.Decimal) {
    //Make Snakes Chain
    Snakes := ReadOmniscientSnakeSnapshot()
    SnakesSum := mvx.AddBalanceIntegerChain(Snakes)
    fmt.Println("A total of ", SnakesSum, "have been read from Omniscient file, on", len(Snakes), " addresses.")
    return Snakes, SnakesSum
}

func MakeCodingDivisionChain() ([]mvx.BalanceESDT, *p.Decimal) {
    //Make Coding Division Chain
    //Uses Amount Exception (Paul holds 50 Company SFTs that aren't Include in the computation)
    MultiChain, _, _ := bloom.CreateCodingDivisionChain()
    Owners := bloom.CreateCodingDivisionOwners(MultiChain)

    All := bloom.CreateCodingDivisionAmountChain(Owners, MultiChain)
    AllException := bloom.CreateCodingDivisionAmountExceptionChain(All, true)
    AllExceptionESDT := mvx.ConvertIntegerSFTtoESDTChain(AllException)
    AllExceptionESDTSorted := mvx.SortBalanceDecimalChain(AllExceptionESDT)
    AllExceptionESDTSortedSum := mvx.AddBalanceDecimalChain(AllExceptionESDTSorted)
    return AllExceptionESDTSorted, AllExceptionESDTSortedSum
}

func ComputeCodingDivisionRewards(Amount *p.Decimal) []mvx.BalanceESDT {

    //Make Snakes Chain
    Snakes,SnakesSum := MakeSnakeChain()

    //Make Coding Division Chain
    CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()

    SnakesDistribution := sm.MULxc(Amount, p.NFS("0.15"))
    CodingDivisionDistribution := sm.MULxc(Amount, p.NFS("0.5"))
    
    fmt.Println("A total of ", SnakesSum, "have been read from Omniscient file, on", len(Snakes), " addresses.")
    fmt.Println("They will receive a total of ", SnakesDistribution, " Tokens")
    fmt.Println("A total of ", CodingDivisionSum, "have been snapshotted, on", len(CodingDivision), " addresses.")
    fmt.Println("They will receive a total of ", CodingDivisionDistribution, " Tokens")
    fmt.Println("")

    SnakeRewardsChain := mvx.ExactPercentualIntegerRewardSplitter(SnakesDistribution, Snakes)
    CodingDivisionRewardsChain := mvx.ExactPercentualDecimalRewardSplitter(CodingDivisionDistribution, CodingDivision)
    TotalRewards := mvx.DecimalChainAdder(SnakeRewardsChain, CodingDivisionRewardsChain)
    TotalRewardsSorted := mvx.SortBalanceDecimalChain(TotalRewards)

    for i := 0; i < len(TotalRewardsSorted); i++ {
	fmt.Println(TotalRewardsSorted[i])
    }

    //mvx.WriteChainBalanceESDT("SnakesRewards.txt", SnakeRewardsChain)
    //mvx.WriteChainBalanceESDT("CDRewards.txt", CodingDivisionRewardsChain)
    mvx.WriteChainBalanceESDT("SummedRewards.txt", TotalRewardsSorted)

    return TotalRewards
}

func main() {
    var (
	CDRewards = `--cd  <Token Amount> as string;
Computes rewards for Coding Division Distribution
50% go to all CD SFTs on user addresses
15% go to all Snakes on user addresses (60% of 50% of 50%)
35% remains at Demiourgos.Holdings™
`
    )

    const (
	CDR = "cd" //string
    )

    FlagCDR := flag.String(CDR, "0", CDRewards)

    flag.Parse()

    //First Option
    if *FlagCDR != "0" {
	ComputeCodingDivisionRewards(p.NFS(*FlagCDR))
    }
}
