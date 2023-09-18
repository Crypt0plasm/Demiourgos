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

func SlipComputer(IFTA *p.Decimal, Ratio *p.Decimal) *p.Decimal {
	Output := sm.TruncateCustom(sm.DIVxc(IFTA, Ratio), 18)
	return Output
}

func SlipComputerSender(Input *p.Decimal) (*p.Decimal, *p.Decimal) {
	A := sm.TruncateCustom(sm.MULxc(Input, p.NFS("0.48")), 18)
	B := sm.SUBxc(Input, A)
	return A, B
}

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
	//Pulecs Vesta Boos: Guest Position is 8
	//Emma Vault Boost: Guest Position is 29
	if Variant == "vesta" || Variant == "koson" {
		GP = 0
	} else if Variant == "blood" {
		GP = 19
	} else if Variant == "vault" {
		GP = 8
	} else if Variant == "vaultbtc" {
		GP = 0
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
	fmt.Println("**************************************************")
	Slip_TrDaniel := p.NFS("6567.794414772907381575")
	Slip_DRX := p.NFS("2405.210727046111560132")
	Slip_Patryx := p.NFS("1705.824485929687472430")
	Slip_Cuciorva := p.NFS("2405.649916054687433843")
	Slip_Pulecs := p.NFS("8241.373867748426204446")
	Slip_Laruentiu := p.NFS("3523.385233989615742560")
	Slip_FrostedK9 := p.NFS("4808.289620185144129115")
	Slip_TheKid := p.NFS("1968.259022226562439904")
	Slip_Buguletu := p.NFS("1477.042344404771491866")
	Slip_Bail := p.NFS("5948.516156062499934074")

	Slip_Ancient := p.NFS("0")
	Rest_Vesta := p.NFS("184332.344036201919718421")

	OutgoingSum := sm.SUMxc(Slip_TrDaniel, Slip_DRX, Slip_Patryx, Slip_Cuciorva, Slip_Pulecs, Slip_Laruentiu, Slip_FrostedK9, Slip_TheKid, Slip_Buguletu, Slip_Bail)

	Slip_Ancient = sm.SUBxc(Rest_Vesta, OutgoingSum)

	TotalSum := sm.SUMxc(Slip_TrDaniel, Slip_DRX, Slip_Patryx, Slip_Cuciorva, Slip_Pulecs, Slip_Laruentiu, Slip_TheKid, Slip_Buguletu, Slip_Bail, Slip_FrostedK9, Slip_Ancient)
	//TotalSum := OutgoingSum

	fmt.Println("Total Vesta Slipped is: ", TotalSum)
	OUROn := sm.TruncateCustom(sm.DIVxc(TotalSum, p.NFS("800")), 18)
	fmt.Println("Ouro needed is: ", OUROn)

	SlipChain := []*p.Decimal{Slip_TrDaniel, Slip_DRX, Slip_Patryx, Slip_Cuciorva, Slip_Pulecs, Slip_Laruentiu, Slip_FrostedK9, Slip_TheKid, Slip_Buguletu, Slip_Bail, Slip_Ancient}

	Liquidity := p.NFS("230415.430045252399648") //Must Update after add Liqudity
	// erd1qqqqqqqqqqqqqpgqkpcm6xqrsfr8gzlh7pdlc2d6gfvfe2j827rstr8jwd
	SendKeep := sm.TruncateCustom(sm.MULxc(Liquidity, p.NFS("0.48")), 18)
	SendToVault := sm.SUBxc(Liquidity, SendKeep)
	fmt.Println("Amount to keep is: ", SendKeep)
	fmt.Println("Amount to send is: ", SendToVault)

	Values := mvx.PercentualSplitter(SendKeep, SlipChain)
	fmt.Println("Made SLIP Values are: ", Values)

	New_TrDaniel := sm.ADDxc(Values[0], vst.LiquidityUserChain[1].SnakeLiq.Frozen)
	New_DRX := sm.ADDxc(Values[1], vst.LiquidityUserChain[2].SnakeLiq.Frozen)
	New_Patryx := sm.ADDxc(Values[2], vst.LiquidityUserChain[3].SnakeLiq.Frozen)
	New_Cuciorva := sm.ADDxc(Values[3], vst.LiquidityUserChain[6].SnakeLiq.Frozen)
	New_Pulecs := sm.ADDxc(Values[3], vst.LiquidityUserChain[8].SnakeLiq.Frozen)
	New_Laurentiu := sm.ADDxc(Values[5], vst.LiquidityUserChain[9].SnakeLiq.Frozen)
	New_Frostedk9 := sm.ADDxc(Values[6], vst.LiquidityUserChain[10].SnakeLiq.Frozen)
	New_TheKid := sm.ADDxc(Values[7], vst.LiquidityUserChain[13].SnakeLiq.Frozen)
	New_Buguletu := sm.ADDxc(Values[8], vst.LiquidityUserChain[21].SnakeLiq.Frozen)
	New_Bail := sm.ADDxc(Values[9], vst.LiquidityUserChain[24].SnakeLiq.Frozen)
	New_Ancient := sm.ADDxc(Values[10], vst.LiquidityUserChain[0].SnakeLiq.Frozen)

	fmt.Println("============")
	fmt.Println("New_TrDaniel: ", New_TrDaniel)
	fmt.Println("New_DRX: ", New_DRX)
	fmt.Println("New_Patryx: ", New_Patryx)
	fmt.Println("New_Cuciorva: ", New_Cuciorva)
	fmt.Println("New_Pulecs: ", New_Pulecs)
	fmt.Println("New_Laurentiu: ", New_Laurentiu)
	fmt.Println("New_Frostedk9: ", New_Frostedk9)
	fmt.Println("New_TheKid: ", New_TheKid)
	fmt.Println("New_Buguletu: ", New_Buguletu)
	fmt.Println("New_Bail: ", New_Bail)
	fmt.Println("New_Ancient: ", New_Ancient)

	fmt.Println("**************************************************")
	Z := sm.MULxc(p.NFS("73"), p.NFS("22.222222222222222222"))
	Z2 := p.NFS("1631.419497294048156625")
	O := sm.SUBxc(Z2, Z)
	fmt.Println("Send GSC is ", O)
	fmt.Println("Z is", Z)

	A := p.NFS("137273.654097714850028660")
	B := p.NFS("125.718665624999987825")
	AplusB := sm.ADDxc(A, B)
	fmt.Println("A plus B: ", AplusB)

	A2 := p.NFS("10000")
	B2 := p.NFS("66")
	A2plusB2 := sm.TruncateCustom(sm.DIVxc(A2, B2), 18)
	fmt.Println("A2 divided B2: ", A2plusB2)

	fmt.Println("**************************************************")
	OURO4SLIP := SlipComputer(p.NFS("87.478178765624967204"), p.NFS("800"))
	fmt.Println("Ouro4Slip is: ", OURO4SLIP)
	SlipLiquidity := p.NFS("16005.43009642456681")
	C, D := SlipComputerSender(SlipLiquidity)
	fmt.Println("Slip for ", SlipLiquidity, "is: ", C, D)

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
