package main

import (
	bloom "Demiourgos/Blooming"
	"Demiourgos/Vesta"
	mvx "MvxApiScanner"
	"fmt"
)

func main() {
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
	SortedAll := mvx.SortBalanceIntegerChain(All)
	AllSum := mvx.AddBalanceDecimalChain(SortedAll)
	fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", AllSum, "SFTs on", len(All), "Addresses")
	fmt.Println("")

	//Net Amount Computations (with single exceptions)
	AllException := bloom.CreateCodingDivisionAmountExceptionChain(All, true)
	SortedAllException := mvx.SortBalanceIntegerChain(AllException)
	AllExceptionSum := mvx.AddBalanceDecimalChain(SortedAllException)
	fmt.Println("Excluding SC, Net (excluding SC, with single exceptions) SFTs sum is ", AllExceptionSum, "SFTs on", len(AllException), "Addresses")
	fmt.Println("")

	//Raw Set Computations
	fmt.Println("====SET Computations====")
	Set := bloom.CreateCodingDivisionSetChain(Owners, MultiChain)
	SortedSet := mvx.SortBalanceIntegerChain(Set)
	SetSum := mvx.AddBalanceDecimalChain(SortedSet)
	fmt.Println("Excluding SC, RAW (excluding SC, no set exceptions) SFTs-SET sum is ", SetSum, "on ", len(SortedSet), " Addresses")
	fmt.Println("")

	//Net Set Computations
	SetException := bloom.CreateCodingDivisionSetExceptionChain(Set, true)
	SortedSetException := mvx.SortBalanceIntegerChain(SetException)
	SetExceptionSum := mvx.AddBalanceDecimalChain(SortedSetException)
	fmt.Println("Excluding SC, Net (excluding SC, with set exceptions) SFTs-SET sum is ", SetExceptionSum, "on ", len(SortedSetException), " Addresses")
	fmt.Println("")

	//Final Stats and file Outputs
	fmt.Println("====Total Stats====")
	fmt.Println("SFTs on Blockchain are ", Sum0, "on ", len(Owners), " Addresses")

	mvx.WriteChainBalanceESDT("Output_All_Raw.txt", SortedAll)
	mvx.WriteChainBalanceESDT("Output_All_Net.txt", SortedAllException)
	mvx.WriteChainBalanceESDT("Output_Set_Raw.txt", SortedSet)
	mvx.WriteChainBalanceESDT("Output_Set_Net.txt", SortedSetException)

	//KYC Comparison
	Path := "d:\\_Crypto\\Demiourgos\\Subsidiaries\\Coding Division\\Blooming\\export_416055238_02-2023.csv"
	KycERDs := bloom.KycScanner(Path)
	SortedSetExceptionKYC := bloom.MakeSetKYC(SortedSetException, KycERDs)
	mvx.WriteChainTrueBalanceSFT("Output_Set_Net_KYC.txt", SortedSetExceptionKYC)
	//Set to Draw
	DrawSET := bloom.CleanKycSet(SortedSetExceptionKYC)
	mvx.WriteChainBalanceESDT("Output_Set_Net_KYC_Draw.txt", DrawSET)
	DrawSUM := mvx.AddBalanceDecimalChain(DrawSET)
	fmt.Println("Drawing will be executed for ", DrawSUM, " units")
	fmt.Println("")

	//Vesta Gold
	fmt.Println("====Vesta GOLD====")
	VestaGoldFull := Vesta.ScanVestaGoldChain()
	VestaGoldException := Vesta.CreateVestaGoldAmounts(VestaGoldFull)

	SortedVestaExceptionGold := mvx.SortBalanceIntegerChain(VestaGoldException)
	SortedVestaExceptionGoldSum := mvx.AddBalanceDecimalChain(SortedVestaExceptionGold)
	fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", SortedVestaExceptionGoldSum, "SFTs on", len(SortedVestaExceptionGold), "Addresses")
	fmt.Println("")
	mvx.WriteChainBalanceESDT("Output_VestaGOLD_Raw.txt", SortedVestaExceptionGold)

}
