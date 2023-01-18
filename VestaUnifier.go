package main

import (
	"Demiourgos/Blooming"
	vt "Demiourgos/Vesta"
	p "Firefly-APD"
	mt "SuperMath"
	"strconv"
	"strings"
)

// function necessary in the SnapshotScanner function
func ProcessScannedLine(Line string) Blooming.BalanceESDT {
	var (
		ProcessedString string
		Output          Blooming.BalanceESDT
	)
	//Remove the { and } character
	ProcessedString = strings.ReplaceAll(Line, "{", "")
	ProcessedString = strings.ReplaceAll(ProcessedString, "}", "")
	Parts := strings.Split(ProcessedString, " ")
	Output.Address = Blooming.ElrondAddress(Parts[0])
	Output.Balance = Parts[1]
	return Output
}

// Scans BalanceESDT SnapshotFiles from HDD and saves them in a program Slice to be used further
func SnapshotScanner(WeekNumber, DayNumber, PoolPosition int, Type string) []Blooming.BalanceESDT {
	var (
		Unit   Blooming.BalanceESDT
		Output []Blooming.BalanceESDT
	)

	ImportName := vt.MakeImportName(WeekNumber, DayNumber, PoolPosition, Type)
	//Path := "d:\\.GO_workspace\\src\\Demiourgos\\_VESTA-Snapshots\\" + ImportName
	Path := "_VESTA-Snapshots\\" + ImportName
	ReadStringSlice := Blooming.ReadFile(Path)
	for i := 0; i < len(ReadStringSlice); i++ {
		Unit = ProcessScannedLine(ReadStringSlice[i])
		Output = append(Output, Unit)
	}
	return Output
}

// Adds two BalanceESDT Files, removing duplicate addresses and summing their balances.
func SnapshotAdder(S1, S2 []Blooming.BalanceESDT) []Blooming.BalanceESDT {
	var (
		ValueS1, ValueS2, TotalValue string
	)
	AllSlice := append(S1, S2...)

	//2.    Make a slice with all Elrond Address (will contain duplicate Elrond Addresses)
	//      basically removes the balance value.
	ElrondSlice := make([]Blooming.ElrondAddress, len(AllSlice))

	for i := 0; i < len(AllSlice); i++ {
		ElrondSlice[i] = AllSlice[i].Address
	}

	Unique := Blooming.RemoveDuplicateElrondAddresses(ElrondSlice)
	Output := make([]Blooming.BalanceESDT, len(Unique))

	for i := 0; i < len(Unique); i++ {
		for j := 0; j < len(S1); j++ {
			if Unique[i] == S1[j].Address {
				ValueS1 = S1[j].Balance
				break
			} else {
				ValueS1 = "0"
			}
		}

		for k := 0; k < len(S2); k++ {
			if Unique[i] == S2[k].Address {
				ValueS2 = S2[k].Balance
				break
			} else {
				ValueS2 = "0"
			}
		}

		TotalValue = mt.DTS(mt.ADDxc(p.NFS(ValueS1), p.NFS(ValueS2)))

		Output[i].Address = Unique[i]
		Output[i].Balance = TotalValue
	}

	return Output
}

// Separates a weekly BalanaceESDT List into 2, using a blacklist.
// Blacklist addreses are saved in the second output with their balances intact.
func SeparateBlacklist(Week int, Input, Blacklist []Blooming.BalanceESDT) ([]Blooming.BalanceESDT, []Blooming.BalanceESDT) {
	//Function that removes given index from a Slice
	RemovePosition := func(s []Blooming.BalanceESDT, index int) []Blooming.BalanceESDT {
		return append(s[:index], s[index+1:]...)
	}

	VestaAirdrop := Input
	BlacklistOutput := make([]Blooming.BalanceESDT, len(Blacklist))
	for i := 0; i < len(Blacklist); i++ {
		for j := 0; j < len(Input); j++ {
			if Input[j].Address == Blacklist[i].Address {
				BlacklistOutput[i] = Input[j]
				VestaAirdrop = RemovePosition(VestaAirdrop, j)
			}
		}
	}
	Vesta := Blooming.SortBalanceESDTChain(VestaAirdrop)
	Black := Blooming.SortBalanceESDTChain(BlacklistOutput)
	WeekString := strconv.Itoa(Week)
	V := WeekString + "_VESTA-Airdrop"
	B := WeekString + "_BLACK-Cumulus"
	Blooming.WriteListOneByOneB(V, Vesta)
	Blooming.WriteListOneByOneB(B, Black)
	return Vesta, Black
}

func MakeMeanCutOfVESTA(Week int) []vt.VestaSplit {
	var Output []vt.VestaSplit
	return Output
}

func main() {
	//ReadSnapshots
	SS1 := SnapshotScanner(1, 1, 0, "VESTA")
	SS2 := SnapshotScanner(1, 2, 0, "VESTA")
	SS3 := SnapshotScanner(1, 3, 0, "VESTA")
	SS4 := SnapshotScanner(1, 4, 0, "VESTA")
	SS5 := SnapshotScanner(1, 5, 0, "VESTA")
	SS6 := SnapshotScanner(1, 6, 0, "VESTA")
	SS7 := SnapshotScanner(1, 7, 0, "VESTA")

	//Add Snapshots and sorts them
	SUM1 := Blooming.SortBalanceESDTChain(SnapshotAdder(SS1, SS2))
	SUM2 := Blooming.SortBalanceESDTChain(SnapshotAdder(SUM1, SS3))
	SUM3 := Blooming.SortBalanceESDTChain(SnapshotAdder(SUM2, SS4))
	SUM4 := Blooming.SortBalanceESDTChain(SnapshotAdder(SUM3, SS5))
	SUM5 := Blooming.SortBalanceESDTChain(SnapshotAdder(SUM4, SS6))
	SUM6 := Blooming.SortBalanceESDTChain(SnapshotAdder(SUM5, SS7))

	//Read BlackList
	BLK := SnapshotScanner(1, 0, 0, "BLACKLIST")

	//Separate Blacklist Values from Snapshot Values and outputs Vesta Airdrop Values.
	SeparateBlacklist(1, SUM6, BLK)
}
