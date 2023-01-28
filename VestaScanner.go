package main

import (
	vt "Demiourgos/Vesta"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	mt "SuperMath"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// PART 1
// Function 				1 ScanPoolsAndWeighThem with IPM (internal Partner Multiplier aka Pool Weight)
// Function 				2 ScanPoolsAndOutputThem, outputs scanned Pool.

// PART 2
// Function 				3 SplitWeeklyVestaPool
// Function 				3 ScanPoolsMakeVestaSplit(WeekNumber)
// Function 				4 ScanWeeklyPoolsMakeMedianVestaSplit(Week Number) Makes PoolVesta Split from weekly PoolScans
// Function 				5 ComputePoolVestaSplit, computes the PoolVestaSplit

// PART 3
// Function 				6 ScanIndividualLPOwners; Scans all addresses owning a given LP Token
// ScanAllLPs 				7 Scan all LPs (the number is defined in the variable above)
// PoolVestaSplitScanner 		8 Reads the PoolVestaSplit from hdd.
// MakeIndividualVestaSplitAllPools 	9 Computes Vesta Split for all Pool, for each of the individuals.
//					Each File is printed out externally.
//					These files will have to be added for each pool resulting in Vesta amounts for each pool
//					And these could be added together for a single Big Vesta Airdrop file.

// ScanPoolsAndWeighThem Function 1
// Scans the Liquidity Pools for Token amounts, and applies IPM (Intrinsic Partner Multiplier)
// Pool Number is updated once per day.
func ScanPoolsAndWeighThem() (Output []mvx.VestaPool) {
	//Scan Pool Number 1, Super Pool
	SuperPool := vt.ScanLiquidity(mvx.EGLDSuperLP, mvx.Super)
	WeighedSuperPool := vt.WeightPool(SuperPool, p.NFS("10"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedSuperPool.VEGLD, p.NFS("0")) == true {
		Output = append(Output, WeighedSuperPool)
	}

	//Scan Pool Number 2, Crust
	CrustPool := vt.ScanLiquidity(mvx.EGLDCrustLP, mvx.Crust)
	WeighedCrustPool := vt.WeightPool(CrustPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedCrustPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedCrustPool)
	}

	//Scan Pool Number 3 Aero
	AeroPool := vt.ScanLiquidity(mvx.EGLDAeroLP, mvx.Aero)
	WeighedAeroPool := vt.WeightPool(AeroPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedAeroPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedAeroPool)
	}
	return
}

// ScanPoolsAndOutputThem Function 2
// OutPuts Pools
func ScanPoolsAndOutputThem(WeekNumber, PoolNumber int, Export bool) []mvx.VestaPool {
	DailyVestaPoolChain := ScanPoolsAndWeighThem()
	OutputName := vt.MakeFileName(true, false, "PS", WeekNumber, -1, -1)
	if Export == true {
		mvx.WriteChainVestaPool(OutputName, DailyVestaPoolChain)
	}

	//Move Export to Folder
	B, _ := mvx.MyCopy(OutputName, vt.VestaSnapshotDirectory+OutputName)
	fmt.Println(B, " bytes copied!")

	return DailyVestaPoolChain
}

// SplitWeeklyVestaPool 3
func SplitWeeklyVestaPool(TotalVesta *p.Decimal, Input []mvx.VestaPool, WeekNumber int, Type string) {
	var OutputName string
	//Compute Pool Split using 800 VESTA (TotalVesta) for all Pools. It splits 800 daily VESTA to multiple Pools.
	PoolVestaSplit := ComputePoolVestaSplit(TotalVesta, Input)
	//Display Information about LiquidityPool Scan and PoolVestaSplit
	for i := 0; i < len(PoolVestaSplit); i++ {
		fmt.Println("Pool ", i, "has ", Input[i].VEGLD, "vEGLD equivalent")
		fmt.Println("Pool ", i, "gets ", PoolVestaSplit[i].Vesta, "Vesta")
	}

	if Type == "i" {
		OutputName = vt.MakeFileName(true, false, "PVSi", WeekNumber, -1, -1)
	} else if Type == "m" {
		OutputName = vt.MakeFileName(true, false, "PVSm", WeekNumber, -1, -1)
	}
	fmt.Println("OutputName IZ ", OutputName)
	mvx.WriteChainVestaSplit(OutputName, PoolVestaSplit)

	//Copy Created File to another Folder
	B, _ := mvx.MyCopy(OutputName, vt.VestaSnapshotDirectory+OutputName)
	fmt.Println(B, " bytes copied!")
	return
}

// ScanPoolsMakeVestaSplit 4
// Makes a PoolVestaSplit based on instantaneous Liquidity-Pool Snapshot
func ScanPoolsMakeVestaSplit(WeekNumber int) {
	InstantaneousLiquidityPoolSnapshot := ScanPoolsAndOutputThem(WeekNumber, -1, false)
	SplitWeeklyVestaPool(vt.WeeklyLiquidityPoolVestaShare, InstantaneousLiquidityPoolSnapshot, WeekNumber, "i")
	return
}

// ScanWeeklyPoolsMakeMedianVestaSplit 4
// Makes a PoolVestaSplit based on daily PoolsScan done over a week period, averaging the values for each Pool
func ScanWeeklyPoolsMakeMedianVestaSplit(WeekNumber int) []mvx.VestaPool {
	//After a few days of pool Scans, I can attempt to build the function
	//var (
	//	S1, S2, S3, S4, S5, S6, S7 []mvx.VestaPool
	//)
	Path := vt.VestaSnapshotDirectory
	WeekDesignation := mvx.TripleDigitDesignation(WeekNumber, "W")
	D1 := "PS_" + WeekDesignation + "_1-Monday.txt"
	D2 := "PS_" + WeekDesignation + "_2-Tuesday.txt"
	D3 := "PS_" + WeekDesignation + "_3-Wednesday.txt"
	D4 := "PS_" + WeekDesignation + "_4-Thursday.txt"
	D5 := "PS_" + WeekDesignation + "_5-Friday.txt"
	D6 := "PS_" + WeekDesignation + "_6-Saturday.txt"
	D7 := "PS_" + WeekDesignation + "_7-Sunday.txt"

	S1 := mvx.ReadVestaPoolSnapshot(Path, D1)
	S2 := mvx.ReadVestaPoolSnapshot(Path, D2)
	S3 := mvx.ReadVestaPoolSnapshot(Path, D3)
	S4 := mvx.ReadVestaPoolSnapshot(Path, D4)
	S5 := mvx.ReadVestaPoolSnapshot(Path, D5)
	S6 := mvx.ReadVestaPoolSnapshot(Path, D6)
	S7 := mvx.ReadVestaPoolSnapshot(Path, D7)

	VestaPoolSum := mvx.MultipleVestaPoolAdder(S1, S2, S3, S4, S5, S6, S7)
	SplitWeeklyVestaPool(vt.WeeklyLiquidityPoolVestaShare, VestaPoolSum, WeekNumber, "m")
	//fmt.Println(VestaPoolSum)

	return VestaPoolSum
	//MakeTotalWeekly Pool Data
}

// ComputePoolVestaSplit 5
// Computes PoolVestaSplit using a total VESTA Value spread over a slice of VestaPools
func ComputePoolVestaSplit(TotalVesta *p.Decimal, Input []mvx.VestaPool) []mvx.VestaSplit {
	InputDecimalSlice := make([]*p.Decimal, len(Input))
	for i := 0; i < len(Input); i++ {
		InputDecimalSlice[i] = Input[i].VEGLD
	}
	OutputDecimalSlice := mvx.PercentualSplitter(TotalVesta, InputDecimalSlice)

	Output := make([]mvx.VestaSplit, len(Input))
	for i := 0; i < len(Input); i++ {
		Output[i].Pool = Input[i]
		Output[i].Vesta = OutputDecimalSlice[i]
	}
	return Output
}

//Part 3

// ScanIndividualLPOwners 6
// Scans a single LPToken for all addresses that own it
func ScanIndividualLPOwners(LPToken mvx.ESDT) (Output []mvx.BalanceESDT) {
	LPScan := mvx.SnapshotIntegerChain(mvx.MakeAddressESDTSnapshotLink(LPToken))
	LPScanESDT := mvx.ConvertSFTAUStoESDTChain(LPScan)
	SortedLPScanESDT := mvx.SortBalanceDecimalChain(LPScanESDT)
	Sum := mvx.AddBalanceDecimalChain(SortedLPScanESDT)
	fmt.Println("LP-Token ", string(LPToken), " sum is ", Sum, " on ", len(SortedLPScanESDT), " addresses.")
	//LP List won't be printed
	return SortedLPScanESDT
}

// ScanAllLPs 7
// Scans All LPs and makes a chain of chains
func ScanAllLPs(PoolNumber int) [][]mvx.BalanceESDT {
	var (
		ScannedLP     []mvx.BalanceESDT
		AllScannedLPs [][]mvx.BalanceESDT
	)
	for i := 0; i < PoolNumber; i++ {
		ScannedLP = ScanIndividualLPOwners(vt.LiquidityPoolOrder[i])
		AllScannedLPs = append(AllScannedLPs, ScannedLP)
	}
	return AllScannedLPs
}

// PoolVestaSplitScanner 8
// Reads The PoolVestaSplit File from hdd
// Only two string types allowed
//
//	Type "i", reads instantaneous PoolVestSplitFile of the given Week.
//	Type "m", reads the median    PoolVestSplitFile of the given Week.
//
// Outputs a slice of Decimals, that is used to compute individual Vesta Splits for each Pool, in their order
// Int Output is the number of lines in the Scanned Vesta Split File
func PoolVestaSplitScanner(WeekNumber int, Type string) ([]*p.Decimal, int) {
	var (
		Unit            *p.Decimal
		Output          []*p.Decimal
		TypeDesignation string
	)

	//Custom function that specifically processes the Content of the PoolVestaSplit file
	GetValueFromLine := func(Line string) *p.Decimal {
		var (
			ProcessedString string
		)
		//Remove the { and } character
		ProcessedString = strings.ReplaceAll(Line, "{", "")
		ProcessedString = strings.ReplaceAll(ProcessedString, "}", "")
		Parts := strings.Split(ProcessedString, " ")
		//The last split string, represents the VestaSplit of the given Pool
		return p.NFS(Parts[2])
	}

	if Type == "i" {
		TypeDesignation = "PVSi"
	} else if Type == "m" {
		TypeDesignation = "PVSm"
	}

	ImportName := vt.MakeFileName(false, false, TypeDesignation, WeekNumber, -1, -1)
	fmt.Println("INIZ ", ImportName)
	Path := vt.VestaSnapshotDirectory + ImportName
	ReadStringSlice := mvx.ReadFile(Path)
	for i := 0; i < len(ReadStringSlice); i++ {
		Unit = GetValueFromLine(ReadStringSlice[i])
		Output = append(Output, Unit)
	}
	return Output, len(ReadStringSlice)
}

// MakeIndividualVestaSplitAllPools 9
// Computes IndividualVestaSplit, for all Pools. Their number is defined at the beginning of the file.
// in the variable LiquidityPoolOrder. This is also their scanned order.
func MakeIndividualVestaSplitAllPools(WeekNumber int, AllScannedLPs [][]mvx.BalanceESDT, VestaValues []*p.Decimal) {
	for i := 0; i < len(VestaValues); i++ {
		LiquidityPoolVestaSplit := mvx.ExactPercentualDecimalRewardSplitter(VestaValues[i], AllScannedLPs[i])
		ExportName := vt.MakeFileName(true, false, "IVS", WeekNumber, i, -1)
		//Exporting POOL
		mvx.WriteChainBalanceESDT(ExportName, LiquidityPoolVestaSplit)
		//Move Export to Folder
		B, _ := mvx.MyCopy(ExportName, vt.VestaSnapshotDirectory+ExportName)
		fmt.Println(B, " bytes copied!")
	}

}

func main() {
	var (
		VarPoolScan = `--ps  <WeekNumber.PoolNumber> as string;
OPTION NO 0: DAILY POOL SCAN
Input Method: WeekNumber followed by point representing the number of pools snapshotted.
Input Example: <12.3> means Week12, 3 Pools to be scanned. Until the PoolNumber can be retrieved automatically,
it must be inserted manually. Code only scans Pools and Outputs the result in their respective files.
`
		VarPoolVestaSplit = `--vs  <WeekNumber(.)> as string;
OPTION NO 1: SNAPSHOT/COMPUTE instantaneous/computed PoolVESTASplit;
Input Variant1: <3.>WeekNumber (with point) makes and Unique one time only Split, after Scanning the Pools in the moment.
		Can be done regardless of week day
Input Variant2: <2> WeekNumber (no point) reads all Pool-Scan-Outputs for the Week, computes mean, and outputs result
		Should be done after all PoolScanSnapshots are done for the Week, that is Sunday.

`
		VarIndividualVestaSplit = `--ivs  <WeekNumber.(i/m)>;
OPTION NO 2: Once PoolVestaSplit Exists, that is used for the whole of next week to compute Individual Vesta Splits
Computes the Individual Vesta Splits for all Pools eligible to mint VESTA for the ongoing Week.
InputMethod: PoolNumbers.(i)
Input Variant1: <2.i> (Week 2 file, and that the "i"nstantaneous snapshot should be used.
Input Variant1: <5.m> (Week 5 file, and that the "m"edian snapshot should be used.
A correct "PoolVestaSplit" must exist(either instantaneous or media), otherwise the IndividualVestaSplit cannot be calculated
The Number is known before hand, from the scans of the past week.

`
	)

	const (
		PoolScan             = "ps"  // string
		PoolVestaSplit       = "pvs" // string
		IndividualVestaSplit = "ivs" // number
	)

	FlagPoolScan := flag.String(PoolScan, "0", VarPoolScan)
	FlagPoolVestaSplit := flag.String(PoolVestaSplit, "0", VarPoolVestaSplit)
	FlagIndividualVestaSplit := flag.String(IndividualVestaSplit, "0", VarIndividualVestaSplit)

	flag.Parse()

	//
	// First Option
	if *FlagPoolScan != "0" {
		//String is split by the "." character
		if strings.Contains(*FlagPoolScan, ".") == true {
			SplitString := strings.Split(*FlagPoolScan, ".")
			WeekNumber, _ := strconv.Atoi(SplitString[0])
			PoolNumber, _ := strconv.Atoi(SplitString[1])
			ScanPoolsAndOutputThem(WeekNumber, PoolNumber, true)
		} else {
			fmt.Println("No Pool Number inputted, please restart program and try again")
		}
	}

	// Second Option
	if *FlagPoolVestaSplit != "0" {
		//String is split by the "." character
		if strings.Contains(*FlagPoolVestaSplit, ".") == true {
			//OneShot Scan and PoolVESTASplit Compute
			SplitString := strings.Split(*FlagPoolVestaSplit, ".")
			WeekNumber, _ := strconv.Atoi(SplitString[0])
			ScanPoolsMakeVestaSplit(WeekNumber) // i option
		} else {
			//Scan All PoolScans for the week and computes mean.
			WeekNumber, _ := strconv.Atoi(*FlagPoolVestaSplit)
			ScanWeeklyPoolsMakeMedianVestaSplit(WeekNumber) // m option
		}
	}

	//Third Option
	if *FlagIndividualVestaSplit != "0" {
		var (
			VestaValues            []*p.Decimal
			PoolNumber, WeekNumber int
		)
		//String is split by the "." character
		if strings.Contains(*FlagIndividualVestaSplit, ".") == true {
			SplitString := strings.Split(*FlagIndividualVestaSplit, ".")
			if len(SplitString) == 2 {
				WeekNumber, _ = strconv.Atoi(SplitString[0])
				VestaType := SplitString[1]
				if VestaType == "i" {
					//Read Individual VestaSplit
					VestaValues, PoolNumber = PoolVestaSplitScanner(WeekNumber, VestaType)
				} else if VestaType == "m" {
					//Read Median Vesta Split
					VestaValues, PoolNumber = PoolVestaSplitScanner(WeekNumber, VestaType)
				} else {
					fmt.Println("PoolVestaSplit Type designation not properly entered.")
				}
			} else {
				fmt.Println("Incorrect format length, please restart program and try again")
			}

			//Read LPs of "PoolNumber" Pools from BlockchainScan
			//Pool Number is inputted by the user, and must be exactly the same with the number of pools
			//that are written in LiquidityPoolOrder, until these can be retrieved from the Master SC
			//That governs all dex pools

			AllScannedLPs := ScanAllLPs(PoolNumber)

			MakeIndividualVestaSplitAllPools(WeekNumber, AllScannedLPs, VestaValues)
		} else {
			fmt.Println("No Points detected, please restart program and try again")
		}
	}

}
