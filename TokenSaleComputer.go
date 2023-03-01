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

var (
	SnakePath = "_TSC_Snapshots\\"
	SnakeName = "DEMIOU-6d1b5c-top-holders-omniscient-tools.csv"

	DT1 = p.NFS("500")
	DT2 = p.NFS("1000")
	DT3 = p.NFS("2000")
	DT4 = p.NFS("5000")
	DT5 = p.NFS("10000")
	DT6 = p.NFS("20000")
	DT7 = p.NFS("50000")
)

func MakeSnakeChain(SnakePathVar, SnakeNameVar string) ([]mvx.BalanceESDT, *p.Decimal) {
	Snakes := mvx.ReadOmniscientSS(SnakePathVar, SnakeNameVar)
	SortedSnakes := mvx.SortBalanceIntegerChain(Snakes)
	SnakeSum := mvx.AddBalanceDecimalChain(SortedSnakes)
	return SortedSnakes, SnakeSum
}

func MakeSnakeTenChain(SnakePathVar, SnakeNameVar string) ([]mvx.BalanceESDT, *p.Decimal) {
	var (
		SnakesTenUnit   mvx.BalanceSFT
		SnakesTenOutput []mvx.BalanceSFT
	)
	Snakes := mvx.ReadOmniscientSS(SnakePathVar, SnakeNameVar)
	for i := 0; i < len(Snakes); i++ {
		if sm.DecimalGreaterThanOrEqual(p.NFS(Snakes[i].Balance), p.NFS("10")) == true {
			SnakesTenUnit = Snakes[i]
			SnakesTenOutput = append(SnakesTenOutput, SnakesTenUnit)
		}
	}

	SortedAll := mvx.SortBalanceIntegerChain(SnakesTenOutput)
	SnakeSum := mvx.AddBalanceDecimalChain(SortedAll)
	return SortedAll, SnakeSum
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

func MakeGoldenVestaChain() ([]mvx.BalanceESDT, *p.Decimal) {
	GoldenVesta, _ := bloom.CreateVestaChain(mvx.VestaGold)
	GoldenVestaTrim := bloom.CreateVestaAmountChains(GoldenVesta)
	GoldenVestaTrimSorted := mvx.SortBalanceIntegerChain(GoldenVestaTrim)
	GoldenVestaTrimSortedSum := mvx.AddBalanceDecimalChain(GoldenVestaTrimSorted)
	return GoldenVestaTrimSorted, GoldenVestaTrimSortedSum
}

func MakeSilverVestaChain() ([]mvx.BalanceESDT, *p.Decimal) {
	SilverVesta, _ := bloom.CreateVestaChain(mvx.VestaSilver)
	SilverVestaTrim := bloom.CreateVestaAmountChains(SilverVesta)
	SilverVestaTrimSorted := mvx.SortBalanceIntegerChain(SilverVestaTrim)
	SilverVestaTrimSortedSum := mvx.AddBalanceDecimalChain(SilverVestaTrimSorted)
	return SilverVestaTrimSorted, SilverVestaTrimSortedSum
}

func GetDemiourgosTier(Input *p.Decimal) string {
	var Tier string
	if sm.DecimalLessThan(Input, DT1) == true {
		Tier = "Tier0 = MORTAL"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT1) && sm.DecimalLessThan(Input, DT2) == true {
		Tier = "Tier1 = PARTNER"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT2) && sm.DecimalLessThan(Input, DT3) == true {
		Tier = "Tier2 = INVESTOR"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT3) && sm.DecimalLessThan(Input, DT4) == true {
		Tier = "Tier3 = ENTREPRENEUR"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT4) && sm.DecimalLessThan(Input, DT5) == true {
		Tier = "Tier4 = BARON"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT5) && sm.DecimalLessThan(Input, DT6) == true {
		Tier = "Tier5 = MOGUL"
	} else if sm.DecimalGreaterThanOrEqual(Input, DT6) && sm.DecimalLessThan(Input, DT7) == true {
		Tier = "Tier6 = MAGNATE"
	} else if sm.DecimalGreaterThan(Input, DT7) == true {
		Tier = "Tier7 = DEMIURG"
	}

	return Tier
}

func AmountForNextTier(Input *p.Decimal) *p.Decimal {
	var Result *p.Decimal
	if sm.DecimalLessThan(Input, DT1) == true {
		Result = sm.SUBxc(DT1, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT1) && sm.DecimalLessThan(Input, DT2) == true {
		Result = sm.SUBxc(DT2, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT2) && sm.DecimalLessThan(Input, DT3) == true {
		Result = sm.SUBxc(DT3, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT3) && sm.DecimalLessThan(Input, DT4) == true {
		Result = sm.SUBxc(DT4, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT4) && sm.DecimalLessThan(Input, DT5) == true {
		Result = sm.SUBxc(DT5, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT5) && sm.DecimalLessThan(Input, DT6) == true {
		Result = sm.SUBxc(DT6, Input)
	} else if sm.DecimalGreaterThanOrEqual(Input, DT6) && sm.DecimalLessThan(Input, DT7) == true {
		Result = sm.SUBxc(DT7, Input)
	} else if sm.DecimalGreaterThan(Input, DT7) == true {
		Result = p.NFS("0")
	}
	return Result
}

func TokenComputer(SnakePathVar, SnakeNameVar string, Addy mvx.MvxAddress) {
	//Snake Chain
	var (
		Price1 = p.NFS("2")
		Price2 = p.NFS("4")
		Price3 = p.NFS("5.6")
		Price4 = p.NFS("6.8")
		Price5 = p.NFS("7.2")
		Price6 = p.NFS("8")

		Round1 = p.NFS("18375")
		Round2 = p.NFS("25725")
		Round3 = p.NFS("36750")
		Round4 = p.NFS("41650")
		//Round5 = p.NFS("61250")
	)

	SnakesTen, SnakesTenSum := MakeSnakeTenChain(SnakePathVar, SnakeNameVar)
	Snakes, SnakesSum := MakeSnakeChain(SnakePathVar, SnakeNameVar)
	CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()
	GoldenVesta, GoldenVestaSum := MakeGoldenVestaChain()
	SilverVesta, SilverVestaSum := MakeSilverVestaChain()
	DemiourgosAssets := sm.SUMxc(SnakesSum, CodingDivisionSum, GoldenVestaSum, SilverVestaSum)

	Round1Amount := sm.TruncateCustom(sm.DIVxc(Round1, SnakesTenSum), 18)
	Round2Amount := sm.TruncateCustom(sm.DIVxc(Round2, SnakesSum), 18)
	Round3Amount := sm.TruncateCustom(sm.DIVxc(Round3, CodingDivisionSum), 18)
	Round4Amount := sm.TruncateCustom(sm.DIVxc(Round4, DemiourgosAssets), 18)

	UserTenSnake := p.NFS(mvx.GetAddressDecimalAmount(Addy, SnakesTen))
	UserSnake := p.NFS(mvx.GetAddressDecimalAmount(Addy, Snakes))
	UserCodingDivision := p.NFS(mvx.GetAddressDecimalAmount(Addy, CodingDivision))
	UserGoldenVesta := p.NFS(mvx.GetAddressDecimalAmount(Addy, GoldenVesta))
	UserSilverVesta := p.NFS(mvx.GetAddressDecimalAmount(Addy, SilverVesta))
	UserAssets := sm.SUMxc(UserSnake, UserCodingDivision, UserGoldenVesta, UserSilverVesta)

	//Token Amounts and Prices
	UserRound1TokenAmount := sm.MULxc(UserTenSnake, Round1Amount)
	UserRound1Price := sm.TruncateCustom(sm.MULxc(UserRound1TokenAmount, Price1), 2)

	UserRound2TokenAmount := sm.MULxc(UserSnake, Round2Amount)
	UserRound2Price := sm.TruncateCustom(sm.MULxc(UserRound2TokenAmount, Price2), 2)

	UserRound3TokenAmount := sm.MULxc(UserCodingDivision, Round3Amount)
	UserRound3Price := sm.TruncateCustom(sm.MULxc(UserRound3TokenAmount, Price3), 2)

	UserRound4TokenAmount := sm.MULxc(UserAssets, Round4Amount)
	UserRound4Price := sm.TruncateCustom(sm.MULxc(UserRound4TokenAmount, Price4), 2)

	TotalTokenAmount := sm.SUMxc(UserRound1TokenAmount, UserRound2TokenAmount, UserRound3TokenAmount, UserRound4TokenAmount)
	TotalPrice := sm.SUMxc(UserRound1Price, UserRound2Price, UserRound3Price, UserRound4Price)

	CurrentTier := GetDemiourgosTier(TotalTokenAmount)
	MissingAmount := AmountForNextTier(TotalTokenAmount)
	MissingAmountPrice := sm.TruncateCustom(sm.MULxc(MissingAmount, Price5), 2)
	MissingAmountPriceMax := sm.TruncateCustom(sm.MULxc(MissingAmount, Price6), 2)

	//Write on Screen
	fmt.Println("Address has ", UserSnake, "Snakes")
	fmt.Println("            ", UserCodingDivision, "Coding Division SFTs")
	fmt.Println("            ", UserGoldenVesta, "Golden Vesta SFTs")
	fmt.Println("            ", UserSilverVesta, "Golden Vesta SFTs")
	fmt.Println("============================================================")
	fmt.Println("The Following NFTs/SFTs amounts are taken into computation")
	fmt.Println("ROUND 1: ", SnakesTenSum, "Snakes NFTs")
	fmt.Println("ROUND 2: ", SnakesSum, "Snakes NFTs")
	fmt.Println("ROUND 3: ", CodingDivisionSum, "CD SFTs")
	fmt.Println("Round 4: ", DemiourgosAssets, "Total NFTs/SFTs (GoldenVESTA SFTs = ", GoldenVestaSum, "and, SilverVestaSFTs = ", SilverVestaSum, ")")
	fmt.Println("============================================================")
	fmt.Println("You will be able to buy the following amount of tokens")
	fmt.Println("ROUND 1: ", UserRound1TokenAmount, " Tokens, for ", UserRound1Price, " USDC")
	fmt.Println("ROUND 2: ", UserRound2TokenAmount, " Tokens, for ", UserRound2Price, " USDC")
	fmt.Println("ROUND 3: ", UserRound3TokenAmount, " Tokens, for ", UserRound3Price, " USDC")
	fmt.Println("ROUND 4: ", UserRound4TokenAmount, " Tokens, for ", UserRound4Price, " USDC")
	fmt.Println("This AMOUNTS to a TOTAL of ", TotalTokenAmount, " Tokens, for ", TotalPrice, " USDC")
	fmt.Println("This would make you a ", CurrentTier)
	fmt.Println("============================================================")
	fmt.Println("To reach the next tier you need to buy ", MissingAmount, " TOKENS in the following rounds, 5 or 6")
	fmt.Println("Round 5 would cost you ", MissingAmountPrice, " USDC, while")
	fmt.Println("Round 6 would cost you ", MissingAmountPriceMax, " USDC.")

}

func main() {

	var (
		TokenComputerStandard = `--tcs  <> as string;
Computes Token Sale possibilities of the given ERD Address using the built in
Snake Snapshot File
Usage: ".exe -tcs erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s"
`
		TokenComputerInput = `--tci  <> as string;
Computes Token Sale possibilities of the given ERD Address using the given
Snake Snapshot File
    Usage: ".exe -tcs SnakeList.csv erd1h6lh2tqjscs4n69c4w4wunu4qw2mz708qn8mqk4quzsyz2syn0aq5gu64s"
`
	)

	const (
		ConstTCS = "tcs" // string
		ConstTCI = "tci" // string
	)

	FlagConstTCS := flag.String(ConstTCS, "", TokenComputerStandard)
	FlagConstTCI := flag.String(ConstTCI, "", TokenComputerInput)

	flag.Parse()

	//Option 1
	if *FlagConstTCS != "" {
		Address := mvx.MvxAddress(*FlagConstTCS)
		TokenComputer(SnakePath, SnakeName, Address)
	}

	//Option 2
	if *FlagConstTCI != "" {
		if strings.Contains(*FlagConstTCI, " ") == true {
			SplitString := strings.Split(*FlagConstTCI, " ")
			SnakeNameInput := SplitString[0]
			Address := mvx.MvxAddress(SplitString[1])
			TokenComputer(SnakePath, SnakeNameInput, Address)
		}
	}
}
