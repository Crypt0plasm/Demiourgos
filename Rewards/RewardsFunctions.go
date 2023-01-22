package Rewards

import (
	bloom "Demiourgos/Blooming"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	sm "SuperMath"
	"fmt"
	"strings"
	"time"
)

var (
	RewardPath = "_REWRD-Snapshots\\"
	OMNsctPath = "_SNAKE-Snapshots\\DEMIOU-6d1b5c-top-holders-omniscient-tools.csv"
)

func ReadOmniscientSnakeSnapshot() []mvx.BalanceSFT {
	var (
		MainUnit mvx.BalanceSFT
		Output   []mvx.BalanceSFT
	)

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

	StringSlice := mvx.ReadFile(OMNsctPath)
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
	AllExceptionSorted := mvx.SortBalanceIntegerChain(AllException)
	AllExceptionSortedSum := mvx.AddBalanceDecimalChain(AllExceptionSorted)
	return AllExceptionSorted, AllExceptionSortedSum
}

// Compute Rewards
func ComputeCodingDivisionRewards(Amount *p.Decimal) []mvx.BalanceESDT {

	//Make Snakes Chain
	Snakes, SnakesSum := MakeSnakeChain()

	//Make Coding Division Chain
	CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()

	DemiourgosDistribution := sm.MULxc(Amount, p.NFS("0.35"))
	SnakesDistribution := sm.MULxc(Amount, p.NFS("0.15"))
	CodingDivisionDistribution := sm.MULxc(Amount, p.NFS("0.5"))

	fmt.Println("")
	fmt.Println("A TOTAL of ", Amount, " Tokens will be distributed as follows:")
	fmt.Println("	1)A total of ", SnakesSum, "Snakes NFTs have been read from Omniscient file, on", len(Snakes), " addresses.")
	fmt.Println("	  They will receive a total of ", SnakesDistribution, " Tokens")
	fmt.Println("	2)A total of ", CodingDivisionSum, "CodingDivision SFTs have been snapshotted, on", len(CodingDivision), " addresses.")
	fmt.Println("	  They will receive a total of ", CodingDivisionDistribution, " Tokens")
	fmt.Println("	3)Demiourgos.Holdings™ retains ", DemiourgosDistribution, " Tokens")
	fmt.Println("")

	SnakeRewardsChain := mvx.ExactPercentualIntegerRewardSplitter(SnakesDistribution, Snakes)
	CodingDivisionRewardsChain := mvx.ExactPercentualDecimalRewardSplitter(CodingDivisionDistribution, CodingDivision)
	TotalRewards := mvx.DecimalChainAdder(SnakeRewardsChain, CodingDivisionRewardsChain)
	TotalRewardsSorted := mvx.SortBalanceDecimalChain(TotalRewards)

	for i := 0; i < len(TotalRewardsSorted); i++ {
		fmt.Println(TotalRewardsSorted[i])
	}

	//Export Files
	EVDName := RewardExport(TotalRewardsSorted, "tCDr", sm.DTS(Amount))

	//Make Evidence and Export it
	Evidence := MakeTotalCDEvidence(DistributionType1, DistributionMode1, Payee3, Payee2, CodingDivisionSum, SnakesSum, Amount, mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceCD(EVDName, Snakes, CodingDivision, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")

	return TotalRewards
}

func ComputeSnakeRewardsByMultiplication(MultiplicationAmount string) []mvx.BalanceESDT {
	Snakes, _ := MakeSnakeChain()
	SnakesRewards := mvx.RewardsComputerIntegerChainMultiplication(Snakes, MultiplicationAmount)
	RewardExport(SnakesRewards, "sSr", MultiplicationAmount)
	return SnakesRewards
}

func ComputeCodingDivisionRewardsByMultiplication(MultiplicationAmount string) []mvx.BalanceESDT {
	CodingDivision, _ := MakeCodingDivisionChain()
	CodingDivisionRewards := mvx.RewardsComputerDecimalChainMultiplication(CodingDivision, MultiplicationAmount)
	RewardExport(CodingDivisionRewards, "sCDr", MultiplicationAmount)
	return CodingDivisionRewards
}

// If Type is true extension is csv, else is txt
func OutputRewardsName(RewardName, Amount string) (R1, R2 string) {
	currentTime := time.Now()
	T := currentTime.Format("2006;(01-January);(02-Day)_(T-15h;04m;05s)")
	AmountToPrint := "[" + Amount + "]"
	R1 = RewardName + "_" + AmountToPrint + "_" + T + ".csv"
	R2 = RewardName + "_" + AmountToPrint + "_" + T + ".evidence"
	return
}

func OutputDualRewardsName(RewardName, Amount string) (CSV, EVD string) {
	CSV, EVD = OutputRewardsName(RewardName, Amount)
	return
}

func RewardExport(InputChain []mvx.BalanceESDT, RewardName, Amount string) string {
	OutputNameCSV, OutputNameEVD := OutputDualRewardsName(RewardName, Amount)
	mvx.ConvertToBulkCSV(OutputNameCSV, InputChain)

	//Copy Files in their folders
	B2, _ := mvx.MyCopy(OutputNameCSV, RewardPath+OutputNameCSV)
	fmt.Println(B2, " bytes copied for the CSV!")
	return OutputNameEVD
}
