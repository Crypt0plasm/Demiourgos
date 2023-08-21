package main

import (
	vst "Demiourgos/Vesta"
	p "Firefly-APD"
	"fmt"
)

func main() {
	//Raw Vesta de Baza e 12.7
	RawVesta := p.NFS("12.700")
	InputUM := p.NFS("4.3525")

	//At 61563.105380652538219620 VLP we have
	//12.7101 Raw Daily.
	//14.771$ represent aproxx 24.625 native LP
	//1.000$ would mean 1667 LP

	//1667LP la Daniel pe bronze produce 63230.105380652538219620 VLP
	//Raw cu care trebe Testat e 13.0438

	//1667LP la Daniel pe silver produce 64230.305380652538219620 VLP
	//Raw cu care trebe Testat e 13.2502

	//1667LP la Daniel pe gold produce 65730.605380652538219620 VLP
	//Raw cu care trebe Testat e 13.5597

	//V1 = Total VLP
	//V2 = VLPSplit
	//V3 = AncientAmount
	//V4 = TotalVestaRewardChain

	V1, V2, V3, V4 := vst.MultipleAbsoluteSplitWithVesta(RawVesta, InputUM, vst.UserChain, vst.LiquidityUserChain)
	//fmt.Println("Total Vesta Amounts are: ", VestaAmountsWithAll)
	fmt.Println("=====================================")
	fmt.Println("Total VLP is: ", V1)
	fmt.Println("**********")
	fmt.Println("Total VLP Split is: ", V2)
	fmt.Println("**********")
	fmt.Println("Ancient Amount is, ", V3)
	fmt.Println("**********")
	fmt.Println("Vesta Rewards Chain for Export is: ", V4)

	vst.ComputeMintPercent(V3, V4)
	vst.ExportOutgoingVestas(vst.UserChain, V4)
}
