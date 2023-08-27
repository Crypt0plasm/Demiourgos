package main

import (
    vst "Demiourgos/Vesta"
    p "Firefly-APD"
    "flag"
    "fmt"
    "strings"
)

func VestaComputer(Variant string, RawVesta, VestaUM *p.Decimal, SFTs []vst.VestaHoldings, LPs []vst.LpHoldings) {
	var (
		GP int
	)
	V1, V2, V3, V4 := vst.MultipleAbsoluteSplitWithVesta(Variant, RawVesta, VestaUM, SFTs, LPs)

	//Variant that uses hardcoded Values
	//V1, V2, V3, V4 := vst.MultipleAbsoluteSplitWithVesta(RawVesta, VestaUM, vst.UserChain, vst.LiquidityUserChain)
	fmt.Println("=====================================")
	fmt.Println("Total VLP is: ", V1)
	fmt.Println("**********")
	fmt.Println("Total VLP Split is: ", V2)
	fmt.Println("**********")
	fmt.Println("Ancient Amount as guest is, ", V3)
	fmt.Println("**********")
	fmt.Println("Vesta Rewards Chain for Export is: ", V4)
	fmt.Println("=====================================")
	//Ancient Vesta Boost: Guest Position is 0
	//Bloodshed Vesta Boost: Guest Position is 19
	if Variant == "vesta" || Variant == "koson" {
		GP = 0
	} else if Variant == "blood" {
		GP = 19
	}
	vst.ComputeMintPercent(GP, V3, V4)
	fmt.Println("Exporting Rewards File:")
	vst.ExportOutgoingVestas(GP, SFTs, V4)
}

func main() {
	var (
		ComputeVST = `--vst  <> as string;
Computes Vesta Split. Enter the following DATA:
(type)/(Raw Vesta Amount)/(vEGLD-VST Universal Multiplier UM)/(Data File)
File must be in the same folder with the executable
Types accepted: vesta, koson
Example without the quotes: 
"--vst vesta:14.4199:4.3525:HardcodedAssets.txt"
`
		ComputeHVST = `--hvst  <> as string;
Computes Vesta Split using Hardcoded Asset Values. Enter the following DATA:
(type)/(Raw Vesta Amount)/(vEGLD-VST Universal Multiplier UM)
File must be in the same folder with the executable
Types accepted: vesta, koson
Example without the quotes: 
"--hvst koson:14.4199:4.3525""
`
		ExportHardcodedAssets = `--eha  <> as bool;
Exports Hardcoded Values into HardcodedAssets.txt; simply run the flag.

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
		Type := strings.Split(ReadString, ":")[0]
		RawVesta := p.NFS(strings.Split(ReadString, ":")[1])
		VestaUM := p.NFS(strings.Split(ReadString, ":")[2])
		FileName := strings.Split(ReadString, ":")[3]
		VestaSFTs, VestaLPs := vst.ImportGroupData(FileName)

		VestaComputer(Type, RawVesta, VestaUM, VestaSFTs, VestaLPs)
	}

	//Option 2
	if *FlagComputeHVesta != "" {
		ReadString := *FlagComputeHVesta
		Type := strings.Split(ReadString, ":")[0]
		RawVesta := p.NFS(strings.Split(ReadString, ":")[1])
		VestaUM := p.NFS(strings.Split(ReadString, ":")[2])
		//FileName := strings.Split(ReadString, ";")[2]
		//VestaSFTs, VestaDexLPs := vst.ImportGroupData(FileName)

		VestaComputer(Type, RawVesta, VestaUM, vst.UserChain, vst.LiquidityUserChain)
	}

	//Option 3
	if *FlagExportHardcodedAssets == true {
		vst.ExportGroupData("HardcodedAssets.txt", vst.UserNameList, vst.UserChain, vst.LiquidityUserChain)
	}

	//Buy := p.NFS("249559.999999999999999001")
	//Owned := p.NFS("110759.999999999999999001")
    	//M1 := p.NFS("397.181340846612475575")
    	//M2 := p.NFS("587.487237387816262791")
    	//Sum := sm.SUBxc(Buy, Owned)
    	//Sum2 := sm.ADDxc(M1, M2)
    	//fmt.Println("Total is", Sum)
    	//fmt.Println("Total is", Sum2)
}
