package main

import (
	vst "Demiourgos/Vesta"
	p "Firefly-APD"
	"flag"
	"fmt"
	"strings"
)

func VestaComputer(RawVesta, VestaUM *p.Decimal, SFTs []vst.VestaHoldings, LPs []vst.VestaLPHoldings) {
	V1, V2, V3, V4 := vst.MultipleAbsoluteSplitWithVesta(RawVesta, VestaUM, SFTs, LPs)

	//Variant that uses hardcoded Values
	//V1, V2, V3, V4 := vst.MultipleAbsoluteSplitWithVesta(RawVesta, VestaUM, vst.UserChain, vst.LiquidityUserChain)
	fmt.Println("=====================================")
	fmt.Println("Total VLP is: ", V1)
	fmt.Println("**********")
	fmt.Println("Total VLP Split is: ", V2)
	fmt.Println("**********")
	fmt.Println("Ancient Amount is, ", V3)
	fmt.Println("**********")
	fmt.Println("Vesta Rewards Chain for Export is: ", V4)
	fmt.Println("=====================================")
	vst.ComputeMintPercent(V3, V4)
	fmt.Println("Exporting Rewards File:")
	vst.ExportOutgoingVestas(SFTs, V4)
}

func main() {
	var (
		ComputeVST = `--vst  <> as string;
Computes Vesta Split. Enter the following DATA:
(Raw Vesta Amount)/(vEGLD-VST Universal Multiplier UM)/(Data File)
File must be in the same folder with the executable
Example without the quotes: 
"--vst 14.4199:4.3525:Participants.csv"
`
		ComputeHVST = `--hvst  <> as string;
Computes Vesta Split using Hardcoded Asset Values. Enter the following DATA:
(Raw Vesta Amount)/(vEGLD-VST Universal Multiplier UM)
File must be in the same folder with the executable
Example without the quotes: 
"--hvst 14.4199:4.3525""
`
		ExportHardcodedAssets = `--eha  <> as bool;
Computes Vesta Split using Hardcoded Asset Values. Enter the following DATA:
(Raw Vesta Amount)/(vEGLD-VST Universal Multiplier UM)
File must be in the same folder with the executable
Example without the quotes: 
"--hvst 14.4199:4.3525""

Hardcoded Assets Order is:
Name;ERD;GoldSFT;SilverSFT;BronzeSFT;GoldLiq;SilverLiq;BronzeLiq;UGoldLiq;USilverLiq;UBronzeLiq
`
	)

	const (
		ConstComputeVesta          = "vst"  //string
		ConstComputeHVesta         = "hvst" //string
		ConstExportHardcodedAssets = "eha"  //bool
	)

	FlagComputeVesta := flag.String(ConstComputeVesta, "", ComputeVST)
	FlagComputeHVesta := flag.String(ConstComputeHVesta, "", ComputeHVST)
	FlagExportHardcodedAssets := flag.Bool(ConstExportHardcodedAssets, false, ExportHardcodedAssets)

	flag.Parse()

	//Option 1
	if *FlagComputeVesta != "" {
		ReadString := *FlagComputeVesta
		RawVesta := p.NFS(strings.Split(ReadString, ":")[0])
		VestaUM := p.NFS(strings.Split(ReadString, ":")[1])
		FileName := strings.Split(ReadString, ":")[2]
		VestaSFTs, VestaLPs := vst.ImportGroupData(FileName)

		VestaComputer(RawVesta, VestaUM, VestaSFTs, VestaLPs)
	}

	//Option 2
	if *FlagComputeHVesta != "" {
		ReadString := *FlagComputeHVesta
		fmt.Println("ReadString este: ", ReadString)
		RawVesta := p.NFS(strings.Split(ReadString, ":")[0])
		VestaUM := p.NFS(strings.Split(ReadString, ":")[1])
		//FileName := strings.Split(ReadString, ";")[2]
		//VestaSFTs, VestaLPs := vst.ImportGroupData(FileName)

		VestaComputer(RawVesta, VestaUM, vst.UserChain, vst.LiquidityUserChain)
	}

	//Option 3
	if *FlagExportHardcodedAssets == true {
		vst.ExportGroupData("HardcodedAssets.txt", vst.UserNameList, vst.UserChain, vst.LiquidityUserChain)
	}

}
