package Vesta

import (
	"Demiourgos/Blooming"
	p "Firefly-APD"
	mt "SuperMath"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type VestaSplit struct {
	Pool  VestaPool
	Vesta *p.Decimal
}

type VestaPool struct {
	VEGLD *p.Decimal
	Token *p.Decimal
}

type ESDT string

var (
	VestaGold = "https://api.multiversx.com/nfts/VESTAXDAO-e6c48c-01/accounts?size=10000"

	//TokenIdentifiers
	WEGLDIdentifier = ESDT("WEGLD-bd4d79")
	VEGLDIdentifier = ESDT("VEGLD-2b9319")

	SuperIdentifier = ESDT("SUPER-507aa6")
	XLHIdentifier   = ESDT("XLH-8daa50")
	CrustIdentifier = ESDT("CRU-a5f4aa")
	AeroIdentifier  = ESDT("AERO-458bbf")

	//LP Tokens
	SUPEREGLD = ESDT("SUPEREGLD-a793b9")
	CRUSTEGLD = ESDT("CRUWEGLD-76c269")
	AEROEGLD  = ESDT("AEROWEGLD-81cc37")

	//Liquidity Pools
	EGLD_Super_LP = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	EGLD_Crust_LP = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqj6daemefdk5kjgy9rs4zsng03kezgxdm2jps3h5n07")
	EGLD_Aero_LP  = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqzjctu8xrgn8jmfp503tajjvzz2zq60v92jpsslkh5a")
)

func MakeESDTSnapshotLink(Token ESDT) string {
	String1 := "https://api.multiversx.com/tokens/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

func VestaGoldSnapshot() []Blooming.BalanceSFT {
	var OutputChain []Blooming.BalanceSFT
	SS := Blooming.OnPage(VestaGold)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func CreateVestaGoldChain() []Blooming.BalanceSFT {
	fmt.Println("Snapshotting VestaGold Addresses and Amounts")
	VestaGoldChain := VestaGoldSnapshot()
	S01 := Blooming.AddBalanceSFTChain(VestaGoldChain)
	fmt.Println(len(VestaGoldChain), "addresses snapshotted with Vesta GOLD SFTs", S01)
	fmt.Println("")
	return VestaGoldChain
}

// Remove Exceptions
func CreateVestaGoldAmounts(Input []Blooming.BalanceSFT) []Blooming.BalanceSFT {
	var (
		Result []Blooming.BalanceSFT
		Unit   Blooming.BalanceSFT
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
//   - Relevant for all 3 Types
//
// DayNumber		- Day Number is only relevant for Import (when ExportImport Variable is false)
//   - Otherwise 0 should be inputted as variable
func MakeFileName(ExportImport, BlackList bool, Type string, WeekNumber, PoolNumber, PoolPosition, DayNumber int) string {
	var (
		Output string
	)
	Extension := ".txt"
	BlackListPrefix := "BL_"
	PoolScanPrefix := "PS_"
	PoolVESTASplitPrefix := "PVS_"
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
			case n == "PVS":
				Output = PoolVESTASplitPrefix + WeekDesignation + Extension
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
			case n == "PVS":
				Output = PoolVESTASplitPrefix + WeekDesignation + Extension
			case n == "IVS":
				Output = IndividualVESTASplitPrefix + WeekDesignation + "_" + ReverseDayString(DayNumber) + "_" + PoolNumberDesignation + "_" + PoolPositionDesignation + Extension
			}
		}
	}
	return Output
}

func MakeExportName(WeekNumber, PoolPosition int, Type string, Daily bool) string {
	var (
		WeekDesignation string
		DayDesignation  string
		PoolDesignation string
		FinalOutputName string
		DayNumberS      string
	)
	if WeekNumber < 10 {
		WeekDesignation = "W0" + strconv.Itoa(WeekNumber)
	} else {
		WeekDesignation = "W" + strconv.Itoa(WeekNumber)
	}

	if PoolPosition < 10 {
		PoolDesignation = "P0" + strconv.Itoa(PoolPosition)
	} else {
		PoolDesignation = "P" + strconv.Itoa(PoolPosition)
	}

	Day := time.Now().Weekday()
	if int(Day) == 0 {
		DayNumberS = strconv.Itoa(int(Day) + 7)
	} else {
		DayNumberS = strconv.Itoa(int(Day))
	}

	DayNameS := Day.String()
	DayDesignation = DayNumberS + "-" + DayNameS

	// Daily bool is only made for "CutOfVESTA", for other Types is set to false and it doesnt matter
	if Daily == true {
		if Type == "BLACKLIST" || Type == "WHITELIST" {
			FinalOutputName = WeekDesignation + "_" + Type + ".txt"
		} else if Type == "CutOfVESTA" {
			FinalOutputName = WeekDesignation + "_" + DayDesignation + "_" + Type + ".txt"
		} else {
			FinalOutputName = WeekDesignation + "_" + DayDesignation + "_" + Type + "_" + PoolDesignation + ".txt"
		}
	} else {
		if Type == "CutOfVESTA" || Type == "BLACKLIST" || Type == "WHITELIST" {
			FinalOutputName = WeekDesignation + "_" + Type + ".txt"
		} else {
			FinalOutputName = WeekDesignation + "_" + DayDesignation + "_" + Type + "_" + PoolDesignation + ".txt"
		}
	}

	return FinalOutputName
}

func MakeImportName(WeekNumber, DayNumber, PoolPosition int, Type string) string {
	var (
		WeekDesignation string
		DayDesignation  string
		PoolDesignation string
		FinalOutputName string
	)
	if WeekNumber < 10 {
		WeekDesignation = "W0" + strconv.Itoa(WeekNumber)
	} else {
		WeekDesignation = "W" + strconv.Itoa(WeekNumber)
	}

	if PoolPosition < 10 {
		PoolDesignation = "P0" + strconv.Itoa(PoolPosition)
	} else {
		PoolDesignation = "P" + strconv.Itoa(PoolPosition)
	}

	switch i := DayNumber; {
	case i == 1:
		DayDesignation = "1-Monday"
	case i == 2:
		DayDesignation = "2-Tuesday"
	case i == 3:
		DayDesignation = "3-Wednesday"
	case i == 4:
		DayDesignation = "4-Thursday"
	case i == 5:
		DayDesignation = "5-Friday"
	case i == 6:
		DayDesignation = "6-Saturday"
	case i == 7:
		DayDesignation = "7-Sunday"
	}

	if Type == "CutOfVESTA" || Type == "BLACKLIST" || Type == "WHITELIST" {
		FinalOutputName = WeekDesignation + "_" + Type + ".txt"
	} else {
		FinalOutputName = WeekDesignation + "_" + DayDesignation + "_" + Type + "_" + PoolDesignation + ".txt"
	}

	return FinalOutputName
}

func ReadESDTAmount(Addy Blooming.ElrondAddress, Token ESDT) string {
	var (
		String1 = "https://api.multiversx.com/accounts/"
		String2 = "/tokens/"

		ScannedJSON ESDTSuperStructure
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(Token)

	Snapshot := Blooming.OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON.Balance
	}

	Result := Blooming.AtomicUnitsStringToDecimalString(Balance)

	return Result
}

func ScanLiquidity(PoolSC Blooming.ElrondAddress, Token ESDT) VestaPool {
	var Result VestaPool
	Result.VEGLD = p.NFS(ReadESDTAmount(PoolSC, WEGLDIdentifier))
	Result.Token = p.NFS(ReadESDTAmount(PoolSC, Token))
	return Result
}

func WeightPool(Pool VestaPool, Weight *p.Decimal) (Output VestaPool) {
	Output.VEGLD = mt.MULxc(Pool.VEGLD, Weight)
	Output.Token = Pool.Token
	return
}

// Copy Function
func MyCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
