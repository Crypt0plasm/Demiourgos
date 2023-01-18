package Vesta

import (
	"Demiourgos/Blooming"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	mt "SuperMath"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CreateVestaGoldAmounts ========================================================================================
//
// [B]01         CreateVestaGoldAmounts
//
//	Create a Chain of all Addresses Containing All VestaGold SFTs and their Amounts and applies SC Exception
func ScanVestaGoldChain() []mvx.BalanceSFT {
	fmt.Println("Snapshotting VestaGold Addresses and Amounts")
	VestaGoldChain := mvx.SnapshotSFTChain(mvx.VestaGold)
	S01 := mvx.AddBalanceIntegerChain(VestaGoldChain)
	fmt.Println(len(VestaGoldChain), "addresses snapshotted with Vesta GOLD SFTs", S01)
	fmt.Println("")
	return VestaGoldChain
}

// CreateVestaGoldAmounts ========================================================================================
//
// [B]02         CreateVestaGoldAmounts
//
//	Create a Chain of all Addresses Containing All VestaGold SFTs and their Amounts and applies SC Exception
func CreateVestaGoldAmounts(Input []mvx.BalanceSFT) []mvx.BalanceSFT {
	var (
		Result []mvx.BalanceSFT
		Unit   mvx.BalanceSFT
	)

	for i := 0; i < len(Input); i++ {
		if Blooming.ComputeExceptionAddress(Input[i].Address) == false {
			Unit.Address = Input[i].Address
			Unit.Balance = Input[i].Balance
			Result = append(Result, Unit)
		}
	}
	return Result
}

// VESTA Distribution Functions

// MakeFileName Function
// ExportImport bool 	- true = Export
// BlackList bool	- true = BlackList (weekly)
// Type			- "PS" = PoolScan	(daily)
//   - "PVS"= PoolVESTASplit (weekly)
//   - "IVS"= IndividualVESTASplit (daily)
//
// WeekNumber		- Number of Week the Snapshot was taken
// PoolNumber		- Number of Pools the Snapshot was executed on
//   - Relevant only for "PS" Type
//
// PoolPosition		- Pool Position the Snapshot was done for
//   - Relevant for only for "IVS"
//
// DayNumber		- Day Number is only relevant for Import (when ExportImport Variable is false)
//   - Otherwise 0 should be inputted as variable
//
// -1 indicates that the number won't be in use.
func MakeFileName(ExportImport, BlackList bool, Type string, WeekNumber, PoolNumber, PoolPosition, DayNumber int) string {
	var (
		Output string
	)
	Extension := ".txt"
	BlackListPrefix := "BL_"
	PoolScanPrefix := "PS_"
	PoolVESTASplitPrefix := "PVSm_"
	PoolVESTASplitInstantaneousPrefix := "PVSi_"
	IndividualVESTASplitPrefix := "IVS_"

	TripleDigitDesignation := func(Number int, Designation string) (StringName string) {
		if Number < 10 {
			StringName = Designation + "00" + strconv.Itoa(Number)
		} else if Number >= 10 && Number < 100 {
			StringName = Designation + "0" + strconv.Itoa(Number)
		} else {
			StringName = Designation + strconv.Itoa(Number)
		}
		return
	}

	WeekDesignation := TripleDigitDesignation(WeekNumber, "W")
	PoolNumberDesignation := TripleDigitDesignation(PoolNumber, "PN")
	PoolPositionDesignation := TripleDigitDesignation(PoolPosition, "POOL")

	DayDesignation := func() (StringName string) {
		var DayNumberToString string
		Day := time.Now().Weekday()
		if int(Day) == 0 {
			DayNumberToString = strconv.Itoa(int(Day) + 7)
		} else {
			DayNumberToString = strconv.Itoa(int(Day))
		}

		DayToString := Day.String()
		StringName = DayNumberToString + "-" + DayToString
		return
	}

	ReverseDayString := func(Number int) (StringName string) {
		switch i := Number; {
		case i == 1:
			StringName = "1-Monday"
		case i == 2:
			StringName = "2-Tuesday"
		case i == 3:
			StringName = "3-Wednesday"
		case i == 4:
			StringName = "4-Thursday"
		case i == 5:
			StringName = "5-Friday"
		case i == 6:
			StringName = "6-Saturday"
		case i == 7:
			StringName = "7-Sunday"
		}
		return
	}

	if ExportImport == true { //Export
		if BlackList == true { //Blacklist
			Output = BlackListPrefix + WeekDesignation + Extension
		} else {
			switch n := Type; {
			case n == "PS":
				Output = PoolScanPrefix + WeekDesignation + "_" + DayDesignation() + "_" + PoolNumberDesignation + Extension
			case n == "PVSm":
				Output = PoolVESTASplitPrefix + WeekDesignation + Extension
			case n == "PVSi":
				Output = PoolVESTASplitInstantaneousPrefix + WeekDesignation + Extension
			case n == "IVS":
				Output = IndividualVESTASplitPrefix + WeekDesignation + "_" + DayDesignation() + "_" + PoolNumberDesignation + "_" + PoolPositionDesignation + Extension
			}
		}
	} else { //Import
		if BlackList == true { //Blacklist
			Output = BlackListPrefix + WeekDesignation + Extension
		} else {
			switch n := Type; {
			case n == "PS":
				Output = PoolScanPrefix + WeekDesignation + "_" + ReverseDayString(DayNumber) + "_" + PoolNumberDesignation + Extension
			case n == "PVSm":
				Output = PoolVESTASplitPrefix + WeekDesignation + Extension
			case n == "PVSi":
				Output = PoolVESTASplitInstantaneousPrefix + WeekDesignation + Extension
			case n == "IVS":
				Output = IndividualVESTASplitPrefix + WeekDesignation + "_" + ReverseDayString(DayNumber) + "_" + PoolNumberDesignation + "_" + PoolPositionDesignation + Extension
			}
		}
	}
	return Output
}

// Scans BalanceESDT SnapshotFiles from HDD and saves them in a program Slice to be used further
func HDDSnapshotScanner(Blacklist bool, Type string, WeekNumber, PoolNumber, PoolPosition, DayNumber int) []mvx.BalanceESDT {
	var (
		Unit   mvx.BalanceESDT
		Output []mvx.BalanceESDT
	)

	ProcessScannedLine := func(Line string) mvx.BalanceESDT {
		var (
			ProcessedString string
			Result          mvx.BalanceESDT
		)
		//Remove the { and } character
		ProcessedString = strings.ReplaceAll(Line, "{", "")
		ProcessedString = strings.ReplaceAll(ProcessedString, "}", "")
		Parts := strings.Split(ProcessedString, " ")
		Result.Address = mvx.MvxAddress(Parts[0])
		Result.Balance = Parts[1]
		return Result
	}

	ImportName := MakeFileName(false, Blacklist, Type, WeekNumber, PoolNumber, PoolPosition, DayNumber)
	//Path := "d:\\.GO_workspace\\src\\Demiourgos\\_VESTA-Snapshots\\" + ImportName
	Path := "_VESTA-Snapshots\\" + ImportName
	ReadStringSlice := mvx.ReadFile(Path)
	for i := 0; i < len(ReadStringSlice); i++ {
		Unit = ProcessScannedLine(ReadStringSlice[i])
		Output = append(Output, Unit)
	}
	return Output
}

// Vesta Pools Liquidity Scans
// wEGLD to be replaced with vEGLD Pools are live

func ScanLiquidity(PoolSC mvx.MvxAddress, Token mvx.ESDT) mvx.VestaPool {
	var Result mvx.VestaPool
	Result.VEGLD = mvx.GetAddressESDTAmount(PoolSC, mvx.WrappedEGLD)
	Result.Token = mvx.GetAddressESDTAmount(PoolSC, Token)
	return Result
}

func WeightPool(Pool mvx.VestaPool, Weight *p.Decimal) (Output mvx.VestaPool) {
	Output.VEGLD = mt.MULxc(Pool.VEGLD, Weight)
	Output.Token = Pool.Token
	return
}
