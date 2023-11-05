package Bloodshed

import (
	bm "Demiourgos/Blooming"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	sm "SuperMath"
	"fmt"
)

//Exception Addresses

var (
	BloodshedExceptions = []mvx.MvxAddress{
		bm.ExA3, bm.ExA4, bm.ExA5, bm.ExA6, bm.ExA7, bm.ExA8, bm.ExA9,
		bm.ExA10, bm.ExA11, bm.ExA12, bm.ExA13, bm.ExA14, bm.ExA17}
)

//======================================================================================================================
//
// CAPITAL FUNCTIONS
//
//	0.1(GetNFTType)
//	0.2(GetOrderMultiplier)
//	0.3(ReadSetComposition)
//
//=============================
//
//	Using an Input Nonce as int64, outputs the NFT Type
//

func GetNFTType(InputNonce int64) []TTV {
	var (
		Position int64
		NFTType  []TTV
	)

	if InputNonce >= 1 && InputNonce <= 152 {
		Position = InputNonce % 8
		if Position == 0 {
			NFTType = LT[len(LT)-1]
			//Position = GetPositionStringCustom(8, 1)
		} else {
			NFTType = LT[Position-1]
			//Position = GetPositionStringCustom(PositionC, 1)
		}
	} else if InputNonce >= 153 && InputNonce <= 554 {
		Position = (InputNonce - 152) % 48
		if Position == 0 {
			NFTType = ET[len(ET)-1]
			//Position = GetPositionStringCustom(8, 1)
		} else {
			NFTType = ET[Position-1]
			//Position = GetPositionStringCustom(PositionC, 1)
		}
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		Position = (InputNonce - 554) % 72
		if Position == 0 {
			NFTType = RT[len(RT)-1]
			//Position = GetPositionStringCustom(8, 1)
		} else {
			NFTType = RT[Position-1]
			//Position = GetPositionStringCustom(PositionC, 1)
		}
	} else if InputNonce >= 1158 && InputNonce <= 8861 {
		Position = (InputNonce - 1157) % 144
		if Position == 0 {
			NFTType = CT[len(CT)-1]
			//Position = GetPositionStringCustom(8, 1)
		} else {
			NFTType = CT[Position-1]
			//Position = GetPositionStringCustom(PositionC, 1)
		}
	}
	return NFTType
}

//
//	Using an Input Nonce as int64, outputs the Order Multiplier
//

func GetOrderMultiplier(InputNonce int64) *p.Decimal {
	var (
		RarityRange     *p.Decimal
		RarityNumber    *p.Decimal
		OrderMultiplier *p.Decimal
		Position        int64
	)
	One := p.NFS("1")
	if InputNonce >= 1 && InputNonce <= 152 {
		RarityRange = p.NFS("100")
		RarityNumber = p.NFS("152")
		Position = InputNonce
	} else if InputNonce >= 153 && InputNonce <= 554 {
		RarityRange = p.NFS("200")
		RarityNumber = p.NFS("402")
		Position = (InputNonce - 152)
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		RarityRange = p.NFS("300")
		RarityNumber = p.NFS("603")
		Position = (InputNonce - 554)
	} else if InputNonce >= 1158 && InputNonce <= 8861 {
		RarityRange = p.NFS("400")
		RarityNumber = p.NFS("7704")
		Position = (InputNonce - 1157)
	}
	if InputNonce == 152 || InputNonce == 554 || InputNonce == 1157 || InputNonce == 8861 {
		OrderMultiplier = One
	} else if InputNonce == 1 {
		OrderMultiplier = p.NFS("2")
	} else if InputNonce == 153 {
		OrderMultiplier = p.NFS("3")
	} else if InputNonce == 555 {
		OrderMultiplier = p.NFS("4")
	} else if InputNonce == 1158 {
		OrderMultiplier = p.NFS("5")
	} else {
		V1 := sm.SUBxc(RarityNumber, One)
		V2 := sm.SUBxc(p.NFI(Position), One)
		V3 := sm.TruncateCustom(sm.DIVxc(RarityRange, V1), 18)
		V4 := sm.TruncateCustom(sm.MULxc(V3, V2), 18)
		OrderPercent := sm.TruncateCustom(sm.SUBxc(RarityRange, V4), 18)
		V5 := sm.TruncateCustom(sm.DIVxc(OrderPercent, p.NFS("100")), 18)
		OrderMultiplier = sm.TruncateCustom(sm.ADDxc(One, V5), 18)
	}
	return OrderMultiplier
}

//
//	Decodes the ResultSlice []int  created with the previous function.
//

func ReadSetComposition(Input []int) {
	fmt.Println("==========================================")
	fmt.Println(Input[0], " T4 Sets")
	fmt.Println("==========================================")
	fmt.Println(Input[1], " T3 Legendary Sets")
	fmt.Println(Input[2], " T3 Epic Sets")
	fmt.Println(Input[3], " T3 Rare Sets")
	fmt.Println(Input[4], " T3 Common Sets")
	fmt.Println("==========================================")
	fmt.Println(Input[5], " T2 Comati Sets")
	fmt.Println(Input[6], " T2 Ursoi Sets")
	fmt.Println(Input[7], " T2 Pileati Sets")
	fmt.Println(Input[8], " T2 Smardoi Sets")
	fmt.Println(Input[9], " T2 Carpian Sets")
	fmt.Println(Input[10], " T2 Tarabostes Sets")
	fmt.Println(Input[11], " T2 Costoboc Sets")
	fmt.Println(Input[12], " T2 Buridavens Sets")
	fmt.Println("==========================================")
	fmt.Println(Input[13], " T1 Epic Comati Sets")
	fmt.Println(Input[14], " T1 Epic Ursoi Sets")
	fmt.Println(Input[15], " T1 Epic Pileati Sets")
	fmt.Println(Input[16], " T1 Epic Smardoi Sets")
	fmt.Println(Input[17], " T1 Epic Carpian Sets")
	fmt.Println(Input[18], " T1 Epic Tarabostes Sets")
	fmt.Println(Input[19], " T1 Epic Costoboc Sets")
	fmt.Println(Input[20], " T1 Epic Buridavens Sets")
	fmt.Println("==========================================")
	fmt.Println(Input[21], " T1 Rare Comati Sets")
	fmt.Println(Input[22], " T1 Rare Ursoi Sets")
	fmt.Println(Input[23], " T1 Rare Pileati Sets")
	fmt.Println(Input[24], " T1 Rare Smardoi Sets")
	fmt.Println(Input[25], " T1 Rare Carpian Sets")
	fmt.Println(Input[26], " T1 Rare Tarabostes Sets")
	fmt.Println(Input[27], " T1 Rare Costoboc Sets")
	fmt.Println(Input[28], " T1 Rare Buridavens Sets")
	fmt.Println("==========================================")
	fmt.Println(Input[29], " T1 Common Comati Sets")
	fmt.Println(Input[30], " T1 Common Ursoi Sets")
	fmt.Println(Input[31], " T1 Common Pileati Sets")
	fmt.Println(Input[32], " T1 Common Smardoi Sets")
	fmt.Println(Input[33], " T1 Common Carpian Sets")
	fmt.Println(Input[34], " T1 Common Tarabostes Sets")
	fmt.Println(Input[35], " T1 Common Costoboc Sets")
	fmt.Println(Input[36], " T1 Common Buridavens Sets")

}

// CONVERTERS

//Snapshot and Conversion Order:
//	1)	Make Snapshot:
//			Function(MakeNFTSnapshotChain) or Function(ImportNFTSnapshotChain)
//			Output-Type([]EstarIndividualNFTChain)
//			Output-Type: mvx.MvxAddress, []int64
//	2)	Convert Snapshot from an []int64 ID to a []TTVScore ID:
//			Function(ConvertChainToScoreChain),
//			Output-Type([]EstarIndividualNFTScoreChain)
//			Output-Type: mvx.MvxAddress, []TTVScore
//				(Here the Order Multiplier is computed and added within the []TTVScore Structure)
//				[]TTVScore is created by default with 1.0 Set Multiplier and 1.0 Week Multiplier
//	3)	Adding the Set Multiplier by scanning for sets:
//			Function(AddSetMultiplier)
//			Output-Type([]EstarIndividualNFTScoreChain)
//			Output-Type: mvx.MvxAddress, []TTVScore
//	4)	Adding the Week Multiplier by using a string as the weekly trait Input
//			String "NONE", means the multiplier is kept neutral (1.0x)
//
//	5)	Creates a Slice with individual NFT Value, in order to have individual NFT Scores for export.
//			Function(MakeIndividualScoreChain)
//			Output-Type(EstarIndividualNFTValueChain)
//			Output-Type: mvx.MvxAddress, []*p.Decimal
//	6)	Creates a Slice with the added individual NFT Value comprising a total score per ERD.
//			Function(ConvertFinal)
//			Output-Type(mvx.BalanceESDT)
//	7)	Trims the exceptions from the Slice created above. An exception Chain is used.
//			Function(mvx.MakeExChainFromBalanceESDT)
//			Output-Type(mvx.BalanceESDT)
//	8)	Using this final Slice of mvx.BalanceESDT which contains the ERD Scores, a reward Slice is computed.
//			Function(mvx.ExactPercentualDecimalRewardSplitter)
//			Output-Type(mvx.BalanceESDT) representing the Token Amounts for each Address.

//============================Reward Functions==========================================================================
//
//	1	Function(MakeNFTSnapshotChain) or Function(ImportNFTSnapshotChain)
//
//		Function(MakeNFTSnapshotChain) is in BloodshedSnapshot.go
//			Requires 3 more functions found there:
//			Function(GetNFTOwner), Function(AddNFTOwner), Function(AddNFTOwnerToChain)
//			An auxiliary function, Function(ExportNFTSnapshotChain) is also there
//
//		Function(ImportNFTSnapshotChain) is in BloodshedSnapshot.go
//			Function(GetAddyPosition) returns the position of an input ERD in the Snapshot Slice.
//			This is used for the SetScan Flag no. 2 in the Main Bloodshed Program.
//
//======================================================================================================================
//
//	2)	Convert Snapshot from an []int64 ID to a []TTVScore ID:
//			Function(ConvertChainToScoreChain),
//			2 More Auxiliary Functions are needed:
//				2.1(ConvertNonceListToScoreList)
//				2.2(ConvertNonceToScore) (This functions requires the 2 capital functions:
//					GetNFTType
//					GetOrderMultiplier
//
//	2.0)Converts an []EstarIndividualNFTChain to a []EstarIndividualNFTScoreChain
//

func ConvertChainToScoreChain(Input []EstarIndividualNFTChain) []EstarIndividualNFTScoreChain {
	var (
		Unit   EstarIndividualNFTScoreChain
		Output []EstarIndividualNFTScoreChain
	)

	for i := 0; i < len(Input); i++ {
		Unit.Address = Input[i].Address
		Unit.ID = ConvertNonceListToScoreList(Input[i].ID)

		Output = append(Output, Unit)
	}
	return Output
}

//
//	2.1)Converts an Input []int64 (a chain of nonces), to a []TTVScore (a chain of reward scores)
//

func ConvertNonceListToScoreList(Input []int64) []TTVScore {
	var (
		Unit   TTVScore
		Output []TTVScore
	)

	for i := 0; i < len(Input); i++ {
		Unit = ConvertNonceToScore(Input[i])
		Output = append(Output, Unit)
	}
	return Output
}

//
//	2.2)Converts an Input []int64 (a chain of nonces), to a []TTVScore (a chain of reward scores)
//

func ConvertNonceToScore(Input int64) TTVScore {
	var (
		ScoreUnit               TTVScore
		ComputedBase            int
		ComputedOrderMultiplier *p.Decimal
	)
	//Compute the TTV
	ScoreUnit.TTV = GetNFTType(Input)

	//Compute Base Bloodshed Score
	if ScoreUnit.TTV[0] == R1 {
		ComputedBase = CommonBS
	} else if ScoreUnit.TTV[0] == R2 {
		ComputedBase = RareBS
	} else if ScoreUnit.TTV[0] == R3 {
		ComputedBase = EpicBS
	} else if ScoreUnit.TTV[0] == R4 {
		ComputedBase = LegendaryBS
	}
	ScoreUnit.Base = ComputedBase

	//Compute OrderMultiplier
	ComputedOrderMultiplier = GetOrderMultiplier(Input)
	ScoreUnit.OM = ComputedOrderMultiplier

	//Set Default values for SetMultiplier and WeekMultiplier to 1.0x
	//Set Default value as false for the SetBoolean (will be set in a later function as true, when sets are checked)
	ScoreUnit.SetBoolean = false
	ScoreUnit.SM = p.NFS("1")
	ScoreUnit.WM = p.NFS("1")

	return ScoreUnit
}

//
//======================================================================================================================
//
//	3)	Adding the Set Multiplier by scanning for sets:
//			Function(AddSetMultiplier)
//			3 More Auxiliary Functions are needed:
//				3.1(BloodshedSetSorting)
//				3.2(SetSorter)
//				3.3(IzEqualTTV)
//
//	Adds Set Multiplier by performing Set Sorting on all NFTs for all ERDs

func AddSetMultiplier(Input []EstarIndividualNFTScoreChain) []EstarIndividualNFTScoreChain {
	var (
		Unit   EstarIndividualNFTScoreChain
		Output []EstarIndividualNFTScoreChain
		Sorted []TTVScore
	)

	for i := 0; i < len(Input); i++ {
		Sorted, _ = BloodshedSetSorting(Input[i].ID)
		Unit.Address = Input[i].Address
		Unit.ID = Sorted
		Output = append(Output, Unit)
	}
	return Output
}

// 3.1)The Bloodshed Set Sorter, sorts a []TTVScore against all possible sets.
//
//	SortedOutput []TTV represents the sorted Score List taking in account the Sets that have been detected
//	ResultSlice []int, is a slice of integers that show how much of each set was detected.
func BloodshedSetSorting(Input []TTVScore) (SortedOutput []TTVScore, ResultSlice []int) {

	var (
		V01                                    int //T4 Sets
		V02, V03, V04, V05                     int //T3 Sets
		V06, V07, V08, V09, V10, V11, V12, V13 int //T2 Sets
		V14, V15, V16, V17, V18, V19, V20, V21 int //T1 Epic Set
		V22, V23, V24, V25, V26, V27, V28, V29 int //T1 Rare Set
		V30, V31, V32, V33, V34, V35, V36, V37 int //T1 Epic Set
	)

	//Check 1
	//Checks if a Tier 4 Set is possible, until no further checks are possible
	//Output1: Number of T4 Sets
	SortedOutput, V01 = SetSorter(p.NFS("2"), T4, Input)
	ResultSlice = append(ResultSlice, V01)

	//Check 2
	//Checks if a Legendary T3 Set is possible, until no further checks are possible
	//Output2: Number of Legendary T3 Sets
	SortedOutput, V02 = SetSorter(p.NFS("1.6"), T3Legendary, SortedOutput)
	ResultSlice = append(ResultSlice, V02)

	//Check 3
	//Checks if an Epic T3 Set is possible, until no further checks are possible
	//Output3: Number of Epic T3 Sets
	SortedOutput, V03 = SetSorter(p.NFS("1.6"), T3Epic, SortedOutput)
	ResultSlice = append(ResultSlice, V03)

	//Check 4
	//Checks if a Rare T3 Set is possible, until no further checks are possible
	//Output4: Number of Rare T3 Sets
	SortedOutput, V04 = SetSorter(p.NFS("1.6"), T3Rare, SortedOutput)
	ResultSlice = append(ResultSlice, V04)

	//Check 5
	//Checks if a Common T3 Set is possible, until no further checks are possible
	//Output5: Number of Common T3 Sets
	SortedOutput, V05 = SetSorter(p.NFS("1.6"), T3Common, SortedOutput)
	ResultSlice = append(ResultSlice, V05)

	//Check 6
	//Checks if a T2 Set is possible, until no further checks are possible.
	//Output6: Number of T2 Sets
	SortedOutput, V06 = SetSorter(p.NFS("1.3"), T2Comati, SortedOutput)
	ResultSlice = append(ResultSlice, V06)
	SortedOutput, V07 = SetSorter(p.NFS("1.3"), T2Ursoi, SortedOutput)
	ResultSlice = append(ResultSlice, V07)
	SortedOutput, V08 = SetSorter(p.NFS("1.3"), T2Pileati, SortedOutput)
	ResultSlice = append(ResultSlice, V08)
	SortedOutput, V09 = SetSorter(p.NFS("1.3"), T2Smardoi, SortedOutput)
	ResultSlice = append(ResultSlice, V09)
	SortedOutput, V10 = SetSorter(p.NFS("1.3"), T2Carpian, SortedOutput)
	ResultSlice = append(ResultSlice, V10)
	SortedOutput, V11 = SetSorter(p.NFS("1.3"), T2Tarabostes, SortedOutput)
	ResultSlice = append(ResultSlice, V11)
	SortedOutput, V12 = SetSorter(p.NFS("1.3"), T2Costoboc, SortedOutput)
	ResultSlice = append(ResultSlice, V12)
	SortedOutput, V13 = SetSorter(p.NFS("1.3"), T2Buridavens, SortedOutput)
	ResultSlice = append(ResultSlice, V13)

	//Check 7
	//Checks if an Epic T1 Set is possible, until no further checks are possible.
	//Output7: Number of T1 Epic Sets for each Tribe
	//	Example Output [1 1 2 0 0 0 3 4] means the respective number of T1 Epic Sets for each Tribe
	//	in this order [Comati, Ursoi, Pileati, Smardoi, Carpian, Tarabostes, Costoboc, Buridavens]
	SortedOutput, V14 = SetSorter(p.NFS("1.1"), ET1Comati, SortedOutput)
	ResultSlice = append(ResultSlice, V14)
	SortedOutput, V15 = SetSorter(p.NFS("1.1"), ET1Ursoi, SortedOutput)
	ResultSlice = append(ResultSlice, V15)
	SortedOutput, V16 = SetSorter(p.NFS("1.1"), ET1Pileati, SortedOutput)
	ResultSlice = append(ResultSlice, V16)
	SortedOutput, V17 = SetSorter(p.NFS("1.1"), ET1Smardoi, SortedOutput)
	ResultSlice = append(ResultSlice, V17)
	SortedOutput, V18 = SetSorter(p.NFS("1.1"), ET1Carpian, SortedOutput)
	ResultSlice = append(ResultSlice, V18)
	SortedOutput, V19 = SetSorter(p.NFS("1.1"), ET1Tarabostes, SortedOutput)
	ResultSlice = append(ResultSlice, V19)
	SortedOutput, V20 = SetSorter(p.NFS("1.1"), ET1Costoboc, SortedOutput)
	ResultSlice = append(ResultSlice, V20)
	SortedOutput, V21 = SetSorter(p.NFS("1.1"), ET1Buridavens, SortedOutput)
	ResultSlice = append(ResultSlice, V21)

	//Check 8
	//Checks if a Rare T1 Set is possible, until no further checks are possible.
	//Output8: Number of T1 Rare Sets for each Tribe, see example above.
	SortedOutput, V22 = SetSorter(p.NFS("1.1"), RT1Comati, SortedOutput)
	ResultSlice = append(ResultSlice, V22)
	SortedOutput, V23 = SetSorter(p.NFS("1.1"), RT1Ursoi, SortedOutput)
	ResultSlice = append(ResultSlice, V23)
	SortedOutput, V24 = SetSorter(p.NFS("1.1"), RT1Pileati, SortedOutput)
	ResultSlice = append(ResultSlice, V24)
	SortedOutput, V25 = SetSorter(p.NFS("1.1"), RT1Smardoi, SortedOutput)
	ResultSlice = append(ResultSlice, V25)
	SortedOutput, V26 = SetSorter(p.NFS("1.1"), RT1Carpian, SortedOutput)
	ResultSlice = append(ResultSlice, V26)
	SortedOutput, V27 = SetSorter(p.NFS("1.1"), RT1Tarabostes, SortedOutput)
	ResultSlice = append(ResultSlice, V27)
	SortedOutput, V28 = SetSorter(p.NFS("1.1"), RT1Costoboc, SortedOutput)
	ResultSlice = append(ResultSlice, V28)
	SortedOutput, V29 = SetSorter(p.NFS("1.1"), RT1Buridavens, SortedOutput)
	ResultSlice = append(ResultSlice, V29)

	//Check 9
	//Checks if a Common T1 Set is possible, until no further checks are possible.
	//Output9: Number of T1 Rare Sets for each Tribe, see example above.
	SortedOutput, V30 = SetSorter(p.NFS("1.1"), CT1Comati, SortedOutput)
	ResultSlice = append(ResultSlice, V30)
	SortedOutput, V31 = SetSorter(p.NFS("1.1"), CT1Ursoi, SortedOutput)
	ResultSlice = append(ResultSlice, V31)
	SortedOutput, V32 = SetSorter(p.NFS("1.1"), CT1Pileati, SortedOutput)
	ResultSlice = append(ResultSlice, V32)
	SortedOutput, V33 = SetSorter(p.NFS("1.1"), CT1Smardoi, SortedOutput)
	ResultSlice = append(ResultSlice, V33)
	SortedOutput, V34 = SetSorter(p.NFS("1.1"), CT1Carpian, SortedOutput)
	ResultSlice = append(ResultSlice, V34)
	SortedOutput, V35 = SetSorter(p.NFS("1.1"), CT1Tarabostes, SortedOutput)
	ResultSlice = append(ResultSlice, V35)
	SortedOutput, V36 = SetSorter(p.NFS("1.1"), CT1Costoboc, SortedOutput)
	ResultSlice = append(ResultSlice, V36)
	SortedOutput, V37 = SetSorter(p.NFS("1.1"), CT1Buridavens, SortedOutput)
	ResultSlice = append(ResultSlice, V37)

	return SortedOutput, ResultSlice
}

// 3.2)Function sorts a []TTVScore comparing it with a set.
//
//	Returns the sorted []TTVScore and an int representing how many times the Set is included.
func SetSorter(SetMultiplier *p.Decimal, Set [][]TTV, Input []TTVScore) ([]TTVScore, int) {
	var (
		Output        []TTVScore
		OutputCounter int
	)

	IzIncluded := func(InputA [][]TTV, InputB []TTVScore) bool {
		var (
			Counter int
			Result  bool
		)
		for i := 0; i < len(InputA); i++ {
			for j := 0; j < len(InputB); j++ {
				if IzEqualTTV(InputA[i], InputB[j].TTV) == true && InputB[j].SetBoolean == false {
					Counter++
					break
				}
			}
		}
		//fmt.Println("Counter iz: ", Counter)
		if Counter >= len(InputA) {
			Result = true
		} else {
			Result = false
		}
		return Result
	}

	Truth := IzIncluded(Set, Input)
	//fmt.Println("Truth iz: ", Truth)

	InputToBeOperatedUpon := Input

	if Truth == false {
		Output = Input
		OutputCounter = 0
	} else {
		for Truth == true {
			OutputCounter++
			//Make every Element in B true, and enter its set multiplier.
			for i := 0; i < len(Set); i++ {
				for j := 0; j < len(Input); j++ {
					if IzEqualTTV(Set[i], Input[j].TTV) == true && InputToBeOperatedUpon[j].SetBoolean == false {
						InputToBeOperatedUpon[j].SetBoolean = true
						InputToBeOperatedUpon[j].SM = SetMultiplier
						break
					}
				}
			}

			//Compute Truth again
			Truth = IzIncluded(Set, InputToBeOperatedUpon)
		}
		Output = InputToBeOperatedUpon
	}

	return Output, OutputCounter
}

// 3.3)Function returns two if 2 []TTV are equal
func IzEqualTTV(InputA, InputB []TTV) bool {
	var Output bool
	if InputA[0] == InputB[0] && InputA[1] == InputB[1] && InputA[2] == InputB[2] && InputA[3] == InputB[3] && InputA[4] == InputB[4] && InputA[5] == InputB[5] && InputA[6] == InputB[6] && InputA[7] == InputB[7] && InputA[8] == InputB[8] {
		Output = true
	} else {
		Output = false
	}
	return Output
}

//
//======================================================================================================================
//
//	4)	Adding the Week Multiplier by using a string as the weekly trait Input
//			Function(AddWeekMultiplier)
//			1 More Auxiliary Function is needed:
//				4.1(BloodshedWeeklyBoosting)
//
//

func AddWeekMultiplier(WeekType string, Input []EstarIndividualNFTScoreChain) []EstarIndividualNFTScoreChain {
	var (
		Unit          EstarIndividualNFTScoreChain
		Output        []EstarIndividualNFTScoreChain
		WeeklyBoosted []TTVScore
	)

	for i := 0; i < len(Input); i++ {
		WeeklyBoosted = BloodshedWeeklyBoosting(WeekType, Input[i].ID)
		Unit.Address = Input[i].Address
		Unit.ID = WeeklyBoosted
		Output = append(Output, Unit)
	}
	return Output
}

func BloodshedWeeklyBoosting(WeekType string, Input []TTVScore) (WeeklyBoosted []TTVScore) {
	WeekMultiplier := GetWeekMultiplier(WeekType)
	WeeklyBoosted = Input

	for i := 0; i < len(Input); i++ {
		NFTDacian := GetBloodshedDacian(Input[i].TTV)
		NFTPotency := GetBloodshedPotency(Input[i].TTV)
		NFTBloodshedTier := GetBloodshedTier(Input[i].TTV)
		NFTBackground := GetBloodshedBackground(Input[i].TTV)
		NFTMainProtection := GetBloodshedMainProtection(Input[i].TTV)
		NFTSecondaryProtection := GetBloodshedSecondaryProtection(Input[i].TTV)
		NFTMainHand := GetBloodshedMainHand(Input[i].TTV)
		NFTOffHand := GetBloodshedMOffHand(Input[i].TTV)

		if WeekType == NFTDacian || WeekType == NFTPotency || WeekType == NFTBloodshedTier || WeekType == NFTBackground || WeekType == NFTMainProtection || WeekType == NFTSecondaryProtection || WeekType == NFTMainHand || WeekType == NFTOffHand {
			WeeklyBoosted[i].WM = WeekMultiplier
		} else if WeekType == "None" || WeekType == "none" || WeekType == "NONE" {
			WeeklyBoosted[i].WM = WeekMultiplier
		}
	}
	return WeeklyBoosted
}

//======================================================================================================================
//
//	5)	Creates a Slice with individual NFT Value, in order to have individual NFT Scores for export.
//			Function(MakeIndividualScoreChain)
//			2 More Auxiliary Functions are needed:
//				5.1(ConvertTTVScoreChainToDecimalChain)
//				5.2(ConvertTTVScoreToValue)
//
//	Converts an []EstarIndividualNFTScoreChain to an []EstarIndividualNFTValueChain

func MakeIndividualScoreChain(Input []EstarIndividualNFTScoreChain) []EstarIndividualNFTValueChain {
	var (
		Unit   EstarIndividualNFTValueChain
		Output []EstarIndividualNFTValueChain
	)

	for i := 0; i < len(Input); i++ {
		Unit.Address = Input[i].Address
		Unit.ID = ConvertTTVScoreChainToDecimalChain(Input[i].ID)

		Output = append(Output, Unit)
	}
	return Output
}

// 5.1)Converts a []TTVScore Slice to a []*p.Decimal Slice, that is
//
//	Converts a Score Structure Slice to its Score Slice.
func ConvertTTVScoreChainToDecimalChain(Input []TTVScore) []*p.Decimal {
	var (
		Unit   *p.Decimal
		Output []*p.Decimal
	)
	for i := 0; i < len(Input); i++ {
		Unit = ConvertTTVScoreToValue(Input[i])
		Output = append(Output, Unit)
	}
	return Output
}

// 5.2)Converts TTVScore to *p.Decimal, that is
//
//	Converts an individual Score Structure to its respective Score
func ConvertTTVScoreToValue(Input TTVScore) *p.Decimal {
	Score := sm.PRDxc(p.NFI(int64(Input.Base)), Input.OM, Input.SM, Input.WM)
	ScoreTr := sm.TruncateCustom(Score, 18)
	return ScoreTr
}

//
//======================================================================================================================
//
//	6)	Creates a Slice with the added individual NFT Value comprising a total score per ERD.
//			Function(ConvertFinal)
//			2 More Auxiliary Function is needed:
//				6.1)SummingDecimalChain
//				6.2)SortBalanceDecimalChain (part of MvxApiScanner Package)
//			Output-Type(mvx.BalanceESDT)
//

func ConvertFinal(Input []EstarIndividualNFTValueChain) []mvx.BalanceESDT {
	var (
		Unit           mvx.BalanceESDT
		Output         []mvx.BalanceESDT
		BalanceDecimal *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Unit.Address = Input[i].Address
		BalanceDecimal = SummingDecimalChain(Input[i].ID)
		Unit.Balance = sm.DTS(BalanceDecimal)
		Output = append(Output, Unit)
	}
	Output2 := mvx.SortBalanceDecimalChain(Output)
	return Output2
}

// 6.1)Converts a *p.Decimal Chain into a single *p.Decimal by adding all elements of the chain together
func SummingDecimalChain(Input []*p.Decimal) *p.Decimal {
	Sum := p.NFS("0")
	for i := 0; i < len(Input); i++ {
		Sum = sm.SUMxc(Sum, Input[i])
	}
	return Sum
}

//
//======================================================================================================================
//
//	7)	Trims the exceptions from the Slice created above. An exception Chain is used.
//			Function(mvx.MakeExChainFromBalanceESDT) (part of MvxApiScanner Package)
//
//======================================================================================================================
//
//	8)	Using this final Slice of mvx.BalanceESDT which contains the ERD Scores, a reward Slice is computed.
//			Function(mvx.ExactPercentualDecimalRewardSplitter) (part of MvxApiScanner Package)
//
//======================================================================================================================
