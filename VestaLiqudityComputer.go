package main

import (
	vst "Demiourgos/Vesta"
	p "Firefly-APD"
	"fmt"
)

func main() {
	RawVesta := p.NFS("12.7116")
	InputUM := p.NFS("4.3475")

	AH, VestaAmountsWithAll := vst.MultipleAbsoluteSplitWithVesta(RawVesta, InputUM, vst.UserChain, vst.LiquidityUserChain)
	//fmt.Println("Total Vesta Amounts are: ", VestaAmountsWithAll)
	fmt.Println("======")
	fmt.Println("AH Amount is, ", AH)

	vst.ComputeMintPercent(AH, VestaAmountsWithAll)
	vst.ExportOutgoingVestas(vst.UserChain, VestaAmountsWithAll)
}
