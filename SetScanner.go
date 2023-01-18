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
	AllSum := mvx.AddBalanceIntegerChain(SortedAll)
	fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", AllSum, "SFTs on", len(All), "Addresses")
	fmt.Println("")

	//Net Amount Computations (with single exceptions)
	AllException := bloom.CreateCodingDivisionAmountExceptionChain(SortedAll, true)
	SortedAllException := mvx.SortBalanceIntegerChain(AllException)
	AllExceptionSum := mvx.AddBalanceIntegerChain(SortedAllException)
	fmt.Println("Excluding SC, Net (excluding SC, with single exceptions) SFTs sum is ", AllExceptionSum, "SFTs on", len(AllException), "Addresses")
	fmt.Println("")

	//Raw Set Computations
	fmt.Println("====SET Computations====")
	Set := bloom.CreateCodingDivisionSetChain(Owners, MultiChain)
	SortedSet := mvx.SortBalanceIntegerChain(Set)
	SetSum := mvx.AddBalanceIntegerChain(SortedSet)
	fmt.Println("Excluding SC, RAW (excluding SC, no set exceptions) SFTs-SET sum is ", SetSum, "on ", len(SortedSet), " Addresses")
	fmt.Println("")

	//Net Set Computations
	SetException := bloom.CreateCodingDivisionSetExceptionChain(SortedSet, true)
	SortedSetException := mvx.SortBalanceIntegerChain(SetException)
	SetExceptionSum := mvx.AddBalanceIntegerChain(SortedSetException)
	fmt.Println("Excluding SC, Net (excluding SC, with set exceptions) SFTs-SET sum is ", SetExceptionSum, "on ", len(SortedSetException), " Addresses")
	fmt.Println("")

	//Reward Chain Computations (Amount Exception Set multiplied with reward 0.025)
	RewardChain := mvx.RewardsComputerIntegerChain(SortedAllException, "0.025")

	//Final Stats and file Outputs
	fmt.Println("====Total Stats====")
	fmt.Println("SFTs on Blockchain are ", Sum0, "on ", len(Owners), " Addresses")

	mvx.WriteChainBalanceSFT("Output_All_Raw.txt", SortedAll)
	mvx.WriteChainBalanceSFT("Output_All_Net.txt", SortedAllException)
	mvx.WriteChainBalanceSFT("Output_Set_Raw.txt", SortedSet)
	mvx.WriteChainBalanceSFT("Output_Set_Net.txt", SortedSetException)

	mvx.WriteChainBalanceESDT("Output_All_Net_Reward.txt", RewardChain)

	//KYC Comparison
	Path := "d:\\_Crypto\\Demiourgos\\Subsidiaries\\Coding Division\\Blooming\\export_416055238_02-2023.csv"
	KycERDs := bloom.KycScanner(Path)
	SortedSetExceptionKYC := bloom.MakeSetKYC(SortedSetException, KycERDs)
	mvx.WriteChainTrueBalanceSFT("Output_Set_Net_KYC.txt", SortedSetExceptionKYC)
	//Set to Draw
	DrawSET := bloom.CleanKycSet(SortedSetExceptionKYC)
	mvx.WriteChainBalanceSFT("Output_Set_Net_KYC_Draw.txt", DrawSET)
	DrawSUM := mvx.AddBalanceIntegerChain(DrawSET)
	fmt.Println("Drawing will be executed for ", DrawSUM, " units")
	fmt.Println("")

	//Vesta Gold
	fmt.Println("====Vesta GOLD====")
	VestaGoldFull := Vesta.ScanVestaGoldChain()
	VestaGoldException := Vesta.CreateVestaGoldAmounts(VestaGoldFull)

	SortedVestaExceptionGold := mvx.SortBalanceIntegerChain(VestaGoldException)
	SortedVestaExceptionGoldSum := mvx.AddBalanceIntegerChain(SortedVestaExceptionGold)
	fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", SortedVestaExceptionGoldSum, "SFTs on", len(SortedVestaExceptionGold), "Addresses")
	fmt.Println("")
	mvx.WriteChainBalanceSFT("Output_VestaGOLD_Raw.txt", SortedVestaExceptionGold)

}
