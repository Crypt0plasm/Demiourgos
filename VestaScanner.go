package main

import (
    "Demiourgos/Blooming"
    vt "Demiourgos/Vesta"
    p "Firefly-APD"
    mt "SuperMath"
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func WriteListOneByOneC(Name string, List []vt.VestaSplit) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}

// PART 1
// Function 1 ScanPool with IPM
// Function 2 MakeMean from Multiple Pool Scans
// Function 3 MakePoolVesta Split (from Pool Scan info, either single, or mean)

// PART 2
// Function 4 ScanLPs
// Function 5 Make Individual Vesta Split (Input: ScanLP Info, and PoolVesta Split info)

// ScanPools Function 1
// Scans the Liquidity Pools for Token amounts, and applies IPM (Intrinsic Partner Multiplier)
// Pool Number is updated once per day.
func ScanPools() (Output []vt.VestaPool) {
	//Scan Pool Number 1, Super Pool
	SuperPool := vt.ScanLiquidity(vt.EGLD_Super_LP, vt.SuperIdentifier)
	WeighedSuperPool := vt.WeightPool(SuperPool, p.NFS("10"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedSuperPool.VEGLD, p.NFS("0")) == true {
		Output = append(Output, WeighedSuperPool)
	}

	//Scan Pool Number 2, Crust
	CrustPool := vt.ScanLiquidity(vt.EGLD_Crust_LP, vt.CrustIdentifier)
	WeighedCrustPool := vt.WeightPool(CrustPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedCrustPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedCrustPool)
	}

	//Scan Pool Number 3 Aero
	AeroPool := vt.ScanLiquidity(vt.EGLD_Aero_LP, vt.AeroIdentifier)
	WeighedAeroPool := vt.WeightPool(AeroPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedAeroPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedAeroPool)
	}

	return
}

func ComputePoolVestaSplit(TotalVesta *p.Decimal, Input []vt.VestaPool) []vt.VestaSplit {
	Output := make([]vt.VestaSplit, len(Input))
	Sum := p.NFS("0")
	var (
		Remainder    *p.Decimal
		Last         *p.Decimal
		VestaForPool *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Sum = mt.ADDxc(Sum, Input[i].VEGLD)
	}

	for i := 0; i < len(Input); i++ {

		Output[i].Pool = Input[i]
		VestaForPool = mt.TruncateCustom(mt.MULxc(mt.DIVxc(Input[i].VEGLD, Sum), TotalVesta), 18)
		if i == 0 {
			Remainder = mt.SUBxc(TotalVesta, VestaForPool)
		} else {
			Remainder = mt.SUBxc(Remainder, VestaForPool)
		}
		if i == len(Input)-2 {
			Last = Remainder
		}

		if i == len(Input)-1 {
			Output[i].Vesta = Last
		} else {
			Output[i].Vesta = VestaForPool
		}
	}

	return Output
}

func ComputeIndividualVestaSplit(TotalVesta *p.Decimal, Input []Blooming.BalanceESDT) []Blooming.BalanceESDT {
	Output := make([]Blooming.BalanceESDT, len(Input))
	Sum := p.NFS("0")
	var (
		Remainder          *p.Decimal
		Last               *p.Decimal
		VestaForIndividual *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Input[i].Balance))
	}

	for i := 0; i < len(Input); i++ {

		VestaForIndividual = mt.TruncateCustom(mt.MULxc(mt.DIVxc(p.NFS(Input[i].Balance), Sum), TotalVesta), 18)
		if i == 0 {
			Remainder = mt.SUBxc(TotalVesta, VestaForIndividual)
		} else {
			Remainder = mt.SUBxc(Remainder, VestaForIndividual)
		}
		if i == len(Input)-2 {
			Last = Remainder
		}

		if i == len(Input)-1 {
			Output[i].Balance = mt.DTS(Last)
		} else {
			Output[i].Balance = mt.DTS(VestaForIndividual)
		}

		Output[i].Address = Input[i].Address

	}

	return Output
}

func ScanLPOwners(LPToken vt.ESDT, ExportName string) (Output []Blooming.BalanceESDT) {
	LPScan := Blooming.CreateBalanceChain(vt.MakeESDTSnapshotLink(LPToken))
	SortedLPScan := Blooming.SortBalanceSFTChain(LPScan)
	SortedLPScanESDT := Blooming.ConvertSFTtoESDTChain(SortedLPScan)
	Sum := Blooming.AtomicUnitsDecimalToDecimalString(Blooming.AddBalanceSFTChain(SortedLPScan))

	fmt.Println("LP-Token ", string(LPToken), " sum is ", Sum, " on ", len(SortedLPScan), " addresses.")

	fmt.Println("File ", ExportName, "won't be printed.")
	//LP List won't be printed
	//Blooming.WriteListOneByOneB(ExportName, SortedLPScanESDT)
	return SortedLPScanESDT
}

func ScanPoolsMakeVestaSplit(WeekNumber int, Daily bool) {
	//Scan Liquidity Pools and Determine VESTA Split
	var VestaFileName string
	LiquidityPools := ScanPools()
	PoolVestaSplit := ComputePoolVestaSplit(p.NFS("800"), LiquidityPools)
	for i := 0; i < len(PoolVestaSplit); i++ {
		fmt.Println("Pool ", i, "has ", LiquidityPools[i].VEGLD, "vEGLD equivalent")
		fmt.Println("Pool ", i, "gets ", PoolVestaSplit[i].Vesta, "Vesta")
	}

	if Daily == true {
		VestaFileName = vt.MakeExportName(WeekNumber, 0, "CutOfVESTA", true)
	} else {
		VestaFileName = vt.MakeExportName(WeekNumber, 0, "CutOfVESTA", false)
	}
	WriteListOneByOneC(VestaFileName, PoolVestaSplit)

	//Copy Created File to another Folder
	B, _ := vt.MyCopy(VestaFileName, "_VESTA-Snapshots\\"+VestaFileName)
	fmt.Println(B, " bytes copied!")
}

func ScanLPsMakeIndividualVestaSplit(WeekNumber int) {
	//Scan LP Owners, needed to compute VESTA amount for each
	//As More Pools are added, more LP types need to be scanned
	LP00 := ScanLPOwners(vt.SUPEREGLD, vt.MakeExportName(0, 0, "LP", false))
	LP01 := ScanLPOwners(vt.CRUSTEGLD, vt.MakeExportName(0, 1, "LP", false))
	LP02 := ScanLPOwners(vt.AEROEGLD, vt.MakeExportName(0, 2, "LP", false))

	//Read Each Pool Vesta Split from the already existing file
	VestaValues := VestaSplitScanner(WeekNumber)

	//Compute Individual Vesta Split - Compute Individual Vesta Split for each LP.
	LP00VS := ComputeIndividualVestaSplit(VestaValues[0], LP00)
	ExportNamePool00 := vt.MakeExportName(WeekNumber, 0, "VESTA", false)
	Blooming.WriteListOneByOneB(ExportNamePool00, LP00VS)
	B00, _ := vt.MyCopy(ExportNamePool00, "_VESTA-Snapshots\\"+ExportNamePool00)
	fmt.Println(B00, " bytes copied!")

	LP01VS := ComputeIndividualVestaSplit(VestaValues[1], LP01)
	ExportNamePool01 := vt.MakeExportName(WeekNumber, 1, "VESTA", false)
	Blooming.WriteListOneByOneB(ExportNamePool01, LP01VS)
	B01, _ := vt.MyCopy(ExportNamePool01, "_VESTA-Snapshots\\"+ExportNamePool01)
	fmt.Println(B01, " bytes copied!")

	LP02VS := ComputeIndividualVestaSplit(VestaValues[2], LP02)
	ExportNamePool02 := vt.MakeExportName(WeekNumber, 2, "VESTA", false)
	Blooming.WriteListOneByOneB(ExportNamePool02, LP02VS)
	B02, _ := vt.MyCopy(ExportNamePool02, "_VESTA-Snapshots\\"+ExportNamePool02)
	fmt.Println(B02, " bytes copied!")
}

func VestaSplitScanner(WeekNumber int) []*p.Decimal {
	var (
		Unit   *p.Decimal
		Output []*p.Decimal
	)

	GetValueFromLine := func(Line string) *p.Decimal {
		var (
			ProcessedString string
		)
		//Remove the { and } character
		ProcessedString = strings.ReplaceAll(Line, "{", "")
		ProcessedString = strings.ReplaceAll(ProcessedString, "}", "")
		Parts := strings.Split(ProcessedString, " ")
		return p.NFS(Parts[2])
	}

	ImportName := vt.MakeImportName(WeekNumber, 0, 0, "CutOfVESTA")
	Path := "_VESTA-Snapshots\\" + ImportName
	ReadStringSlice := Blooming.ReadFile(Path)
	for i := 0; i < len(ReadStringSlice); i++ {
		Unit = GetValueFromLine(ReadStringSlice[i])
		Output = append(Output, Unit)
	}
	return Output
}

func main() {
	var (
		ScanLiquidity = `--sl  WeekNumber as string;
WeekNumber followed by point character makes a daily snapshot file.
Scan the Inputed Pools, and outputs the VESTA Split for the Pools for the Week
`
		CIVS = `--vs  WeekNumber;
Computes the Individual Vesta Splits for all Pools for the Week. The VestaPoolSplit
must exist in its designated folder.
`
	)

	const (
		ScnLiquidity = "sl" // number
		CmpVesta     = "vs" // number
	)

	FlagScnLiquidity := flag.String(ScnLiquidity, "0", ScanLiquidity)
	FlagCmpVesta := flag.Int(CmpVesta, 0, CIVS)

	flag.Parse()

	//
	// First Option
	if *FlagScnLiquidity != "0" {
		// String "." signal to make a daily split
		if strings.Contains(*FlagScnLiquidity, ".") == true {
			SplitString := strings.Split(*FlagScnLiquidity, ".")
			Number, _ := strconv.Atoi(SplitString[0])
			ScanPoolsMakeVestaSplit(Number, true)
		} else {
			Number, _ := strconv.Atoi(*FlagScnLiquidity)
			ScanPoolsMakeVestaSplit(Number, false)
		}
	}

	// Second Option
	if *FlagCmpVesta != 0 {
		ScanLPsMakeIndividualVestaSplit(*FlagCmpVesta)
	}

}
