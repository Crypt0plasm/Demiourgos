package main

import (
    vst "Demiourgos/Vesta"
    p "Firefly-APD"
    mvx "MvxApiScanner"
    sm "SuperMath"
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

    Slip_TrDaniel := p.NFS("15123.508834720582170251")
    Slip_DRX := p.NFS("5284.675775354580154297")
    Slip_Patryx := p.NFS("3520.559097374999979661")
    Slip_Cuciorva := p.NFS("5233.263523124999959499")
    Slip_Laruentiu := p.NFS("4376.911310249999928825")
    Slip_MakeAStep := p.NFS("9645.841586504894510732")
    Slip_Buguletu := p.NFS("1903.004917499999937687")
    Slip_Bail := p.NFS("12940.433438999999906962")
    Slip_FrostedK9 := p.NFS("7231.418686499999920720")
    Slip_Ancient := p.NFS("285866.786823133593750844")

    TotalSum := sm.SUMxc(Slip_TrDaniel, Slip_DRX, Slip_Patryx, Slip_Cuciorva, Slip_Laruentiu, Slip_MakeAStep, Slip_Buguletu, Slip_Bail, Slip_FrostedK9, Slip_Ancient)
    fmt.Println("Total Veste Slipped is: ", TotalSum)

    SlipChain := []*p.Decimal{Slip_TrDaniel, Slip_DRX, Slip_Patryx, Slip_Cuciorva, Slip_Laruentiu, Slip_MakeAStep, Slip_Buguletu, Slip_Bail, Slip_FrostedK9, Slip_Ancient}

    Liquidity := p.NFS("438908.004991829562774")
    SendKeep := sm.TruncateCustom(sm.MULxc(Liquidity, p.NFS("0.48")), 18)
    fmt.Println("Amount to send is: ", SendKeep)

    Values := mvx.PercentualSplitter(SendKeep, SlipChain)
    fmt.Println("Values are: ", Values)

    A := p.NFS("171520.072093880156250376")
    B := p.NFS("204800")
    AplusB := sm.ADDxc(A, B)
    fmt.Println("A plus B is: ", AplusB)

    //MakeAStep01 := p.NFS("587.487237387816262791")
    //Rogojan01 := p.NFS("20.256852240155188848")
    //Bloodshed := p.NFS("24452.66380850879469346")
    //Ancient := p.NFS("-24452.66380850879469346")

    //Buy := p.NFS("249559.999999999999999001")
    //Owned := p.NFS("110759.999999999999999001")
    //M1 := p.NFS("397.181340846612475575")
    //M2 := p.NFS("587.487237387816262791")
    //Total de adaugat la Make a step dupa purge: 984.668578234428738366
    //Sum := sm.SUBxc(Buy, Owned)
    //Sum2 := sm.ADDxc(M1, M2)
    //fmt.Println("Total is", Sum)
    //fmt.Println("Total is", Sum2)
}
