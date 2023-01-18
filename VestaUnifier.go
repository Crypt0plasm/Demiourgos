package main

import (
	vt "Demiourgos/Vesta"
	mvx "MvxApiScanner"
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

func main() {
	//ReadSnapshots
	SS1 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 1)
	SS2 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 2)
	SS3 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 3)
	SS4 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 4)
	SS5 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 5)
	SS6 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 6)
	SS7 := vt.HDDSnapshotScanner(false, "IVS", 1, 3, 0, 7)

	//Add Snapshots and sorts them
	SUM1 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SS1, SS2))
	SUM2 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SUM1, SS3))
	SUM3 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SUM2, SS4))
	SUM4 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SUM3, SS5))
	SUM5 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SUM4, SS6))
	SUM6 := mvx.SortBalanceDecimalChain(mvx.DecimalChainAdder(SUM5, SS7))

	//Read BlackList
	BLK := vt.HDDSnapshotScanner(true, "", 1, -1, -1, -1)

	//Separate Blacklist Values from Snapshot Values and outputs Vesta Airdrop Values.
	SeparateBlacklist(1, SUM6, BLK)
}
