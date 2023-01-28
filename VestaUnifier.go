package main

import (
    vt "Demiourgos/Vesta"
    mvx "MvxApiScanner"
    "fmt"
    "strconv"
)

// SeparateBlacklist
// Separates a weekly BalanceESDT List into 2, using a blacklist.
// Blacklist addresses are saved in the second output with their balances intact.
func SeparateBlacklist(Week int, Input, Blacklist []mvx.BalanceESDT) ([]mvx.BalanceESDT, []mvx.BalanceESDT) {
    //Function that removes given index from a Slice
    RemovePosition := func(s []mvx.BalanceESDT, index int) []mvx.BalanceESDT {
	return append(s[:index], s[index+1:]...)
    }

    VestaAirdrop := Input
    BlacklistOutput := make([]mvx.BalanceESDT, len(Blacklist))
    for i := 0; i < len(Blacklist); i++ {
	for j := 0; j < len(Input); j++ {
	    if Input[j].Address == Blacklist[i].Address {
		BlacklistOutput[i] = Input[j]
		VestaAirdrop = RemovePosition(VestaAirdrop, j)
	    }
	}
    }
    Vesta := mvx.SortBalanceDecimalChain(VestaAirdrop)
    Black := mvx.SortBalanceDecimalChain(BlacklistOutput)
    WeekString := strconv.Itoa(Week)
    V := WeekString + "_VESTA-Airdrop"
    B := WeekString + "_BLACK-Cumulus"
    mvx.WriteChainBalanceESDT(V, Vesta)
    mvx.WriteChainBalanceESDT(B, Black)
    return Vesta, Black
}

func UnifyPoolWeeklyValues(WeekNumber, PoolPosition int) []mvx.BalanceESDT {
    N1 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 1)
    N2 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 2)
    N3 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 3)
    N4 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 4)
    N5 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 5)
    N6 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 6)
    N7 := vt.MakeFileName(false, false, "IVS", WeekNumber, PoolPosition, 7)

    C1 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N1)
    fmt.Println(C1)
    C2 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N2)
    C3 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N3)
    C4 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N4)
    C5 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N5)
    C6 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N6)
    C7 := mvx.ReadBalanceChain(vt.VestaSnapshotDirectory, N7)

    Sum := mvx.MultipleDecimalChainAdder(C1, C2, C3, C4, C5, C6, C7)
    return Sum
}

func main() {
    //ReadSnapshots
    U0 := UnifyPoolWeeklyValues(1, 0)
    U1 := UnifyPoolWeeklyValues(1, 1)
    U2 := UnifyPoolWeeklyValues(1, 2)
    EndList := mvx.MultipleDecimalChainAdder(U0, U1, U2)
    SumSorted := mvx.SortBalanceDecimalChain(EndList)
    fmt.Println("Lungimea is", len(SumSorted))
    mvx.WriteChainBalanceESDT("Testu-pulii.txt", SumSorted)

    //Read BlackList
    //BLK := vt.HDDSnapshotScanner(true, "", 1, -1, -1, -1)

    //Separate Blacklist Values from Snapshot Values and outputs Vesta Airdrop Values.
    //SeparateBlacklist(1, SUM6, BLK)
}
