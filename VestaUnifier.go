package main

import (
	"Demiourgos/Blooming"
	vt "Demiourgos/Vesta"
	"fmt"
	"strings"
)

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

func SnapshotScanner(WeekNumber, DayNumber, PoolPosition int, Type string) []Blooming.BalanceESDT {
	var (
		Unit   Blooming.BalanceESDT
		Output []Blooming.BalanceESDT
	)

	ImportName := vt.MakeImportName(WeekNumber, DayNumber, PoolPosition, Type)
	fmt.Println("ImportName Iz ", ImportName)
	Path := "d:\\.GO_workspace\\src\\Demiourgos\\_VESTA-Snapshots\\" + ImportName
	fmt.Println("PATH iz ", Path)
	ReadStringSlice := Blooming.ReadFile(Path)
	for i := 0; i < len(ReadStringSlice); i++ {
		Unit = ProcessScannedLine(ReadStringSlice[i])
		Output = append(Output, Unit)
	}
	return Output
}

func main() {
	SS1 := SnapshotScanner(2, 6, 0, "VESTA")
	SS2 := SnapshotScanner(2, 7, 0, "VESTA")
	Blooming.WriteListOneByOneB("Test1.txt", SS1)
	Blooming.WriteListOneByOneB("Test2.txt", SS2)
}
