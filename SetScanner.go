package main

import (
	bloom "Demiourgos/Blooming"
	"Demiourgos/Rewards"
	mvx "MvxApiScanner"
	"flag"
	"fmt"
)

func SetScanner() (SortedAll, SortedAllException, SortedSet, SortedSetException []mvx.BalanceESDT) {
	//CodingDivisionSetScanner()
	fmt.Println("")
	fmt.Println("====Snapshotting====")
	MultiChain, Len, Sum0 := bloom.CreateCodingDivisionChain()
	Owners := bloom.CreateCodingDivisionOwners(MultiChain)
	fmt.Println("A total of", Len, "have been snapshotted")
	fmt.Println("")
	fmt.Println("")

	//Raw Amount Computations
	fmt.Println("====Amount Computations====")
	All := bloom.CreateCodingDivisionAmountChain(Owners, MultiChain)
	SortedAll = mvx.SortBalanceIntegerChain(All)
	AllSum := mvx.AddBalanceDecimalChain(SortedAll)
	fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", AllSum, "SFTs on", len(All), "Addresses")
	fmt.Println("")

	//Net Amount Computations (with single exceptions)
	AllException := bloom.CreateCodingDivisionAmountExceptionChain(All, true)
	SortedAllException = mvx.SortBalanceIntegerChain(AllException)
	AllExceptionSum := mvx.AddBalanceDecimalChain(SortedAllException)
	fmt.Println("Excluding SC, Net (excluding SC, with single exceptions) SFTs sum is ", AllExceptionSum, "SFTs on", len(AllException), "Addresses")
	fmt.Println("")

	//Raw Set Computations
	fmt.Println("====SET Computations====")
	Set := bloom.CreateCodingDivisionSetChain(Owners, MultiChain)
	SortedSet = mvx.SortBalanceIntegerChain(Set)
	SetSum := mvx.AddBalanceDecimalChain(SortedSet)
	fmt.Println("Excluding SC, RAW (excluding SC, no set exceptions) SFTs-SET sum is ", SetSum, "on ", len(SortedSet), " Addresses")
	fmt.Println("")

	//Net Set Computations
	SetException := bloom.CreateCodingDivisionSetExceptionChain(Set, true)
	SortedSetException = mvx.SortBalanceIntegerChain(SetException)
	SetExceptionSum := mvx.AddBalanceDecimalChain(SortedSetException)
	fmt.Println("Excluding SC, Net (excluding SC, with set exceptions) SFTs-SET sum is ", SetExceptionSum, "on ", len(SortedSetException), " Addresses")
	fmt.Println("")

	//Final Stats and file Outputs
	fmt.Println("====Total Stats====")
	fmt.Println("SFTs on Blockchain are ", Sum0, "on ", len(Owners), " Addresses")

	return
}

func KycSetComputer(KYC string) (SortedSetExceptionKYC []mvx.TrueBalanceESDT, DrawSET []mvx.BalanceESDT) {
	//SortedAll, SortedAllException, SortedSet, SortedSetException := SetScanner()
	_, _, _, SortedSetException := SetScanner()

	//KYC Comparison
	//KYC := "KYC_29.01.2023_06.08.csv"

	Path := bloom.KYCDirectory + KYC
	KycERDs := bloom.KycScanner(Path)
	mvx.WriteChainMvxAddresses("AllKYCAddresses.txt", KycERDs)
	SortedSetExceptionKYC = bloom.MakeSetKYC(SortedSetException, KycERDs)
	mvx.WriteChainTrueBalanceSFT("Output_Set_Net_KYC.txt", SortedSetExceptionKYC)

	//Set to Draw
	DrawSET = bloom.CleanKycSet(SortedSetExceptionKYC)
	mvx.WriteChainBalanceESDT("Output_Set_Net_KYC_Draw.txt", DrawSET)
	DrawSUM := mvx.AddBalanceDecimalChain(DrawSET)
	fmt.Println("Drawing will be executed for ", DrawSUM, " units")
	fmt.Println("")
	return
}

func InvestorScanner() int {
	//Investor Computation
	Snakes, _ := Rewards.MakeSnakeChain()
	CD, _ := Rewards.MakeCodingDivisionChain()
	GoldenVesta, _ := Rewards.MakeGoldenVestaChain()
	SilverVesta, _ := Rewards.MakeSilverVestaChain()
	Total := mvx.MultipleDecimalChainAdder(Snakes, CD, GoldenVesta, SilverVesta)

	fmt.Println("There are ", len(Total), " unique Demiourgos NFTs/SFTs holders, excluding SC")
	return len(Total)
}

func SFTScanner(SFTDesignation string) []mvx.BalanceESDT {
	ScanSFT := mvx.SFT(SFTDesignation)
	SFTChain := mvx.SnapshotSFTChain(ScanSFT)
	SortedAll := mvx.SortBalanceIntegerChain(SFTChain)
	return SortedAll
}

func main() {
	var (
		SetScannerPrint = `--ssp  <> as bool;
Scans Coding Division set Numbers and prints information on screen
`
		SetScannerPrintOutput = `--sso  <> as bool;
Scans Coding Division set Numbers, prints information on screen,
and outputs files locally
`
		KycBloomComputer = `--kycbloom  <> as string;
Scans Coding Division Sets, compares them to a KYC file,
and outputs the Draw list.
`
		InvestorComputer = `--ivs  <> bool;
Computes the number of Unique ERD Addresses holding Demiourgos Assets
`
		SFTScannerVar = `--sft  <> bool;
Scans an SFT Chain and outputs it in a txt file.
`
	)

	const (
		ConstSetScannerPrint       = "ssp"      // bool
		ConstSetScannerPrintOutput = "sso"      // bool
		ConstKycBloomComputer      = "kycbloom" // string
		ConstInvestorComputer      = "ivs"      // bool
		ConstSFTScanner            = "sft"
	)

	FlagSetScannerPrint := flag.Bool(ConstSetScannerPrint, false, SetScannerPrint)
	FlagSetScannerPrintOutput := flag.Bool(ConstSetScannerPrintOutput, false, SetScannerPrintOutput)
	FlagKycBloomComputer := flag.String(ConstKycBloomComputer, "", KycBloomComputer)
	FlagInvestorComputer := flag.Bool(ConstInvestorComputer, false, InvestorComputer)
	FlagConstSFTScanner := flag.String(ConstSFTScanner, "", SFTScannerVar)

	flag.Parse()

	//Option 1
	if *FlagSetScannerPrint != false {
		SetScanner()
	}

	//Option 2
	if *FlagSetScannerPrintOutput != false {
		S1, S2, S3, S4 := SetScanner()

		mvx.WriteChainBalanceESDT("Output_All_Raw.txt", S1)
		mvx.WriteChainBalanceESDT("Output_All_Net.txt", S2)
		mvx.WriteChainBalanceESDT("Output_Set_Raw.txt", S3)
		mvx.WriteChainBalanceESDT("Output_Set_Net.txt", S4)
	}

	//Option 3
	if *FlagKycBloomComputer != "" {
		KycSetComputer(*FlagKycBloomComputer)
	}

	//Option 4
	if *FlagInvestorComputer != false {
		InvestorScanner()
	}

	//Option 5
	if *FlagConstSFTScanner != "" {
		SFTChain := SFTScanner(*FlagConstSFTScanner)
		Name := *FlagConstSFTScanner + ".txt"
		mvx.WriteChainBalanceESDT(Name, SFTChain)
	}

}
