package Bloodshed

import (
	p "Firefly-APD"
	sm "SuperMath"
)

type NFTBloodShed struct {
	Rarity              string
	Dacian              string
	Potency             string
	BloodShedTier       string
	Background          string
	MainProtection      string
	SecondaryProtection string
	MainHand            string
	OffHand             string
}

// NFT Variants
var (
	BSC = []string{"C144", "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9",
		"C10", "C11", "C12", "C13", "C14", "C15", "C16", "C17", "C18", "C19",
		"C20", "C21", "C22", "C23", "C24", "C25", "C26", "C27", "C28", "C29",
		"C30", "C31", "C32", "C33", "C34", "C35", "C36", "C37", "C38", "C39",
		"C40", "C41", "C42", "C43", "C44", "C45", "C46", "C47", "C48", "C49",
		"C50", "C51", "C52", "C53", "C54", "C55", "C56", "C57", "C58", "C59",
		"C60", "C61", "C62", "C63", "C64", "C65", "C66", "C67", "C68", "C69",
		"C70", "C71", "C72", "C73", "C74", "C75", "C76", "C77", "C78", "C79",
		"C80", "C81", "C82", "C83", "C84", "C85", "C86", "C87", "C88", "C89",
		"C90", "C91", "C92", "C93", "C94", "C95", "C96", "C97", "C98", "C99",
		"C100", "C101", "C102", "C103", "C104", "C105", "C106", "C107", "C108", "C109",
		"C110", "C111", "C112", "C113", "C114", "C115", "C116", "C117", "C118", "C119",
		"C120", "C121", "C122", "C123", "C124", "C125", "C126", "C127", "C128", "C129",
		"C130", "C131", "C132", "C133", "C134", "C135", "C136", "C137", "C138", "C139",
		"C140", "C141", "C142", "C143"}

	BSR = []string{"R72", "R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9",
		"R10", "R11", "R12", "R13", "R14", "R15", "R16", "R17", "R18", "R19",
		"R20", "R21", "R22", "R23", "R24", "R25", "R26", "R27", "R28", "R29",
		"R30", "R31", "R32", "R33", "R34", "R35", "R36", "R37", "R38", "R39",
		"R40", "R41", "R42", "R43", "R44", "R45", "R46", "R47", "R48", "R49",
		"R50", "R51", "R52", "R53", "R54", "R55", "R56", "R57", "R58", "R59",
		"R60", "R61", "R62", "R63", "R64", "R65", "R66", "R67", "R68", "R69",
		"R70", "R71"}

	BSE = []string{"E48", "E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8", "E9",
		"E10", "E11", "E12", "E13", "E14", "E15", "E16", "E17", "E18", "E19",
		"E20", "E21", "E22", "E23", "E24", "E25", "E26", "E27", "E28", "E29",
		"E30", "E31", "E32", "E33", "E34", "E35", "E36", "E37", "E38", "E39",
		"E40", "E41", "E42", "E43", "E44", "E45", "E46", "E47"}

	BSL = []string{"L8", "L1", "L2", "L3", "L4", "L5", "L6", "L7"}
)

func GetNFTVariant(InputNonce int64) (Output string) {
	var DivisionRest int64
	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Dacian Algorithm
		DivisionRest = InputNonce % 8
		Output = BSL[DivisionRest]
	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Dacian Algorithm
		DivisionRest = (InputNonce - 152) % 48
		Output = BSE[DivisionRest]
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Dacian Algorithm
		DivisionRest = (InputNonce - 554) % 72
		Output = BSR[DivisionRest]
	} else {
		//Common Dacian Algorithm
		DivisionRest = (InputNonce - 1157) % 144
		Output = BSC[DivisionRest]
	}
	return Output
}

func CheckPosition(InputVariable int64, InputChain []int64) (Output bool) {
	for i := 0; i < len(InputChain); i++ {
		if InputVariable == InputChain[i] {
			Output = true
		}
	}
	return
}

func GetDacian2(InputNonce int64) (Output string) {
	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Dacian Algorithm

	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Dacian Algorithm

	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Dacian Algorithm

	} else {
		//Common Dacian Algorithm

	}
	return
}

func GetNFTPosition(InputNonce int64) int64 {
	var Position int64
	if InputNonce >= 1 && InputNonce <= 152 {
		Position = InputNonce
	} else if InputNonce >= 153 && InputNonce <= 554 {
		Position = InputNonce - 152
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		Position = InputNonce - 554
	} else {
		Position = InputNonce - 1157
	}
	return Position
}

func GetNFTRarity(InputNonce int64) string {
	var Rarity string
	if InputNonce >= 1 && InputNonce <= 152 {
		Rarity = "Legendary"
	} else if InputNonce >= 153 && InputNonce <= 554 {
		Rarity = "Epic"
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		Rarity = "Rare"
	} else {
		Rarity = "Common"
	}
	return Rarity
}

func GetNFTDacian(InputNonce int64) (DacianOutput string) {
	var DivisionRest int64

	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Dacian Algorithm
		DivisionRest = InputNonce % 8
		if DivisionRest == 1 {
			DacianOutput = Dacian[0]
		} else if DivisionRest == 2 {
			DacianOutput = Dacian[1]
		} else if DivisionRest == 3 {
			DacianOutput = Dacian[2]
		} else if DivisionRest == 4 {
			DacianOutput = Dacian[3]
		} else if DivisionRest == 5 {
			DacianOutput = Dacian[4]
		} else if DivisionRest == 6 {
			DacianOutput = Dacian[5]
		} else if DivisionRest == 7 {
			DacianOutput = Dacian[6]
		} else if DivisionRest == 0 {
			DacianOutput = Dacian[7]
		}
	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Dacian Algorithm
		DivisionRest = (InputNonce - 152) % 48
		if DivisionRest >= 1 && DivisionRest <= 6 {
			DacianOutput = Dacian[0]
		} else if DivisionRest >= 7 && DivisionRest <= 12 {
			DacianOutput = Dacian[1]
		} else if DivisionRest >= 13 && DivisionRest <= 18 {
			DacianOutput = Dacian[2]
		} else if DivisionRest >= 19 && DivisionRest <= 24 {
			DacianOutput = Dacian[3]
		} else if DivisionRest >= 25 && DivisionRest <= 30 {
			DacianOutput = Dacian[4]
		} else if DivisionRest >= 31 && DivisionRest <= 36 {
			DacianOutput = Dacian[5]
		} else if DivisionRest >= 37 && DivisionRest <= 42 {
			DacianOutput = Dacian[6]
		} else {
			DacianOutput = Dacian[7]
		}
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Dacian Algorithm
		DivisionRest = (InputNonce - 554) % 72
		if DivisionRest >= 1 && DivisionRest <= 9 {
			DacianOutput = Dacian[0]
		} else if DivisionRest >= 10 && DivisionRest <= 18 {
			DacianOutput = Dacian[1]
		} else if DivisionRest >= 19 && DivisionRest <= 27 {
			DacianOutput = Dacian[2]
		} else if DivisionRest >= 28 && DivisionRest <= 36 {
			DacianOutput = Dacian[3]
		} else if DivisionRest >= 37 && DivisionRest <= 45 {
			DacianOutput = Dacian[4]
		} else if DivisionRest >= 46 && DivisionRest <= 54 {
			DacianOutput = Dacian[5]
		} else if DivisionRest >= 55 && DivisionRest <= 63 {
			DacianOutput = Dacian[6]
		} else {
			DacianOutput = Dacian[7]
		}
	} else {
		//Common Dacian Algorithm
		DivisionRest = (InputNonce - 1157) % 144
		if DivisionRest >= 1 && DivisionRest <= 9 {
			DacianOutput = Dacian[0]
		} else if DivisionRest >= 10 && DivisionRest <= 18 {
			DacianOutput = Dacian[1]
		} else if DivisionRest >= 19 && DivisionRest <= 27 {
			DacianOutput = Dacian[2]
		} else if DivisionRest >= 28 && DivisionRest <= 36 {
			DacianOutput = Dacian[3]
		} else if DivisionRest >= 37 && DivisionRest <= 45 {
			DacianOutput = Dacian[4]
		} else if DivisionRest >= 46 && DivisionRest <= 54 {
			DacianOutput = Dacian[5]
		} else if DivisionRest >= 55 && DivisionRest <= 63 {
			DacianOutput = Dacian[6]
		} else {
			DacianOutput = Dacian[7]
		}
	}
	return
}

func GetNFTPotency(InputNonce int64) (PotencyOutput string) {
	var DivisionRest int64
	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Potency Algorithm
		PotencyOutput = Potency[2]
	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 7, 8, 13, 14, 19, 20, 25, 26, 31, 32, 37, 38, 43, 44}
		PremiumEpicPositions := []int64{3, 4, 9, 10, 15, 16, 21, 22, 27, 28, 33, 34, 39, 40, 45, 46}
		EliteEpicPositions := []int64{5, 6, 11, 12, 17, 18, 23, 24, 29, 30, 35, 36, 41, 42, 47, 0}
		DivisionRest = (InputNonce - 152) % 48
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			PotencyOutput = Potency[0]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			PotencyOutput = Potency[1]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			PotencyOutput = Potency[2]
		}
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 3, 10, 11, 12, 19, 20, 21, 28, 29, 30, 37, 38, 39, 46, 47, 48, 55, 56, 57, 64, 65, 66}
		PremiumEpicPositions := []int64{4, 5, 6, 13, 14, 15, 22, 23, 24, 31, 32, 33, 40, 41, 42, 49, 50, 51, 58, 59, 60, 67, 68, 69}
		EliteEpicPositions := []int64{7, 8, 9, 16, 17, 18, 25, 26, 27, 34, 35, 36, 43, 44, 45, 52, 53, 54, 61, 62, 63, 70, 71, 0}
		DivisionRest = (InputNonce - 554) % 72
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			PotencyOutput = Potency[0]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			PotencyOutput = Potency[1]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			PotencyOutput = Potency[2]
		}
	} else {
		//Common Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 3, 4, 5, 6, 19, 20, 21, 22, 23, 24, 37, 38, 39, 40, 41, 42, 55, 56, 57, 58, 59, 60, 73, 74, 75, 76, 77, 78, 91, 92, 93, 94, 95, 96, 109, 110, 111, 112, 113, 114, 127, 128, 129, 130, 131, 132}
		PremiumEpicPositions := []int64{7, 8, 9, 10, 11, 12, 25, 26, 27, 28, 29, 30, 43, 44, 45, 46, 47, 48, 61, 62, 63, 64, 65, 66, 79, 80, 81, 82, 83, 84, 97, 98, 99, 100, 101, 102, 115, 116, 117, 118, 119, 120, 133, 134, 135, 136, 137, 138}
		EliteEpicPositions := []int64{13, 14, 15, 16, 17, 18, 31, 32, 33, 34, 35, 36, 49, 50, 51, 52, 53, 54, 67, 68, 69, 70, 71, 72, 85, 86, 87, 88, 89, 90, 103, 104, 105, 106, 107, 108, 121, 122, 123, 124, 125, 126, 139, 140, 141, 142, 143, 0}
		DivisionRest = (InputNonce - 1157) % 144
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			PotencyOutput = Potency[0]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			PotencyOutput = Potency[1]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			PotencyOutput = Potency[2]
		}
	}
	return
}

func GetNFTBloodshedTier(InputNonce int64) (BloodshedTierOutput string) {
	var DivisionRest int64
	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Potency Algorithm
		BloodshedTierOutput = Bloodshed[5]
	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 7, 8, 13, 14, 19, 20, 25, 26, 31, 32, 37, 38, 43, 44}
		PremiumEpicPositions := []int64{3, 4, 9, 10, 15, 16, 21, 22, 27, 28, 33, 34, 39, 40, 45, 46}
		EliteEpicPositions := []int64{5, 6, 11, 12, 17, 18, 23, 24, 29, 30, 35, 36, 41, 42, 47, 0}
		DivisionRest = (InputNonce - 152) % 48
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[2]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[3]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[4]
		}
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 3, 10, 11, 12, 19, 20, 21, 28, 29, 30, 37, 38, 39, 46, 47, 48, 55, 56, 57, 64, 65, 66}
		PremiumEpicPositions := []int64{4, 5, 6, 13, 14, 15, 22, 23, 24, 31, 32, 33, 40, 41, 42, 49, 50, 51, 58, 59, 60, 67, 68, 69}
		EliteEpicPositions := []int64{7, 8, 9, 16, 17, 18, 25, 26, 27, 34, 35, 36, 43, 44, 45, 52, 53, 54, 61, 62, 63, 70, 71, 0}
		DivisionRest = (InputNonce - 554) % 72
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[1]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[2]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[3]
		}
	} else {
		//Common Potency Algorithm
		StandardEpicPositions := []int64{1, 2, 3, 4, 5, 6, 19, 20, 21, 22, 23, 24, 37, 38, 39, 40, 41, 42, 55, 56, 57, 58, 59, 60, 73, 74, 75, 76, 77, 78, 91, 92, 93, 94, 95, 96, 109, 110, 111, 112, 113, 114, 127, 128, 129, 130, 131, 132}
		PremiumEpicPositions := []int64{7, 8, 9, 10, 11, 12, 25, 26, 27, 28, 29, 30, 43, 44, 45, 46, 47, 48, 61, 62, 63, 64, 65, 66, 79, 80, 81, 82, 83, 84, 97, 98, 99, 100, 101, 102, 115, 116, 117, 118, 119, 120, 133, 134, 135, 136, 137, 138}
		EliteEpicPositions := []int64{13, 14, 15, 16, 17, 18, 31, 32, 33, 34, 35, 36, 49, 50, 51, 52, 53, 54, 67, 68, 69, 70, 71, 72, 85, 86, 87, 88, 89, 90, 103, 104, 105, 106, 107, 108, 121, 122, 123, 124, 125, 126, 139, 140, 141, 142, 143, 0}
		DivisionRest = (InputNonce - 1157) % 144
		if CheckPosition(DivisionRest, StandardEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[0]
		} else if CheckPosition(DivisionRest, PremiumEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[1]
		} else if CheckPosition(DivisionRest, EliteEpicPositions) == true {
			BloodshedTierOutput = Bloodshed[2]
		}
	}
	return
}

func GetNFTBackground(InputNonce int64) (BGOutput string) {
	var DivisionRest int64
	if InputNonce >= 1 && InputNonce <= 152 {
		//Legendary Dacian Algorithm
		DivisionRest = InputNonce % 8
		if DivisionRest == 1 {
			BGOutput = LegendaryBG[0]
		} else if DivisionRest == 2 {
			BGOutput = LegendaryBG[1]
		} else if DivisionRest == 3 {
			BGOutput = LegendaryBG[2]
		} else if DivisionRest == 4 {
			BGOutput = LegendaryBG[3]
		} else if DivisionRest == 5 {
			BGOutput = LegendaryBG[4]
		} else if DivisionRest == 6 {
			BGOutput = LegendaryBG[5]
		} else if DivisionRest == 7 {
			BGOutput = LegendaryBG[6]
		} else if DivisionRest == 0 {
			BGOutput = LegendaryBG[7]
		}
	} else if InputNonce >= 153 && InputNonce <= 554 {
		//Epic Dacian Algorithm
		EP1Positions := []int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47}
		EP2Positions := []int64{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46}
		DivisionRest = (InputNonce - 152) % 48
		if CheckPosition(DivisionRest, EP1Positions) == true {
			BGOutput = EpicBG[0]
		} else if CheckPosition(DivisionRest, EP2Positions) == true {
			BGOutput = EpicBG[1]
		}
	} else if InputNonce >= 555 && InputNonce <= 1157 {
		//Rare Dacian Algorithm
		RR1Positions := []int64{1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34, 37, 40, 43, 46, 49, 52, 55, 58, 61, 64, 67, 70}
		RR2Positions := []int64{2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35, 38, 41, 44, 47, 50, 53, 56, 59, 62, 65, 68, 71}
		RR3Positions := []int64{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 0}
		DivisionRest = (InputNonce - 554) % 72
		if CheckPosition(DivisionRest, RR1Positions) == true {
			BGOutput = RareBG[0]
		} else if CheckPosition(DivisionRest, RR2Positions) == true {
			BGOutput = RareBG[1]
		} else if CheckPosition(DivisionRest, RR3Positions) == true {
			BGOutput = RareBG[2]
		}
	} else {
		//Common Dacian Algorithm
		CC1Positions := []int64{1, 7, 13, 19, 25, 31, 37, 43, 49, 55, 61, 67, 73, 79, 85, 91, 97, 103, 109, 115, 121, 127, 133, 139}
		CC2Positions := []int64{2, 8, 14, 20, 26, 32, 38, 44, 50, 56, 62, 68, 74, 80, 86, 92, 98, 104, 110, 116, 122, 128, 134, 140}
		CC3Positions := []int64{3, 9, 15, 21, 27, 33, 39, 45, 51, 57, 63, 69, 75, 81, 87, 93, 99, 105, 111, 117, 123, 129, 135, 141}
		CC4Positions := []int64{4, 10, 16, 22, 28, 34, 40, 46, 52, 58, 64, 70, 76, 82, 88, 94, 100, 106, 112, 118, 124, 130, 136, 142}
		CC5Positions := []int64{5, 11, 17, 23, 29, 35, 41, 47, 53, 59, 65, 71, 77, 83, 89, 95, 101, 107, 113, 119, 125, 131, 137, 143}
		CC6Positions := []int64{6, 12, 18, 24, 30, 36, 42, 48, 54, 60, 66, 72, 78, 84, 90, 96, 102, 108, 114, 120, 126, 132, 138, 0}
		DivisionRest = (InputNonce - 1157) % 144
		if CheckPosition(DivisionRest, CC1Positions) == true {
			BGOutput = CommonBG[0]
		} else if CheckPosition(DivisionRest, CC2Positions) == true {
			BGOutput = CommonBG[1]
		} else if CheckPosition(DivisionRest, CC3Positions) == true {
			BGOutput = CommonBG[2]
		} else if CheckPosition(DivisionRest, CC4Positions) == true {
			BGOutput = CommonBG[3]
		} else if CheckPosition(DivisionRest, CC5Positions) == true {
			BGOutput = CommonBG[4]
		} else if CheckPosition(DivisionRest, CC6Positions) == true {
			BGOutput = CommonBG[5]
		}
	}
	return
}

func GetNFTGearFromDacian(InputDacian string) (S1, S2, S3, S4 string) {
	GMP1 := Protection[1]
	GMP2 := Protection[3]
	GSP1 := Protection[0]
	GSP2 := Protection[2]

	GMH1 := Weapons[0]
	GMH2 := Weapons[1]
	GMH3 := Weapons[4]
	GOH1 := Weapons[0]
	GOH2 := Weapons[1]
	GOH3 := Weapons[2]
	GOH4 := Weapons[3]
	GOH5 := Weapons[4]
	if InputDacian == Dacian[0] {
		S1 = GMP1
		S2 = GSP1
		S3 = GMH3
		S4 = GOH2
	} else if InputDacian == Dacian[1] {
		S1 = GMP1
		S2 = GSP1
		S3 = GMH1
		S4 = GOH1
	} else if InputDacian == Dacian[2] {
		S1 = GMP1
		S2 = GSP2
		S3 = GMH2
		S4 = GOH4
	} else if InputDacian == Dacian[3] {
		S1 = GMP1
		S2 = GSP1
		S3 = GMH2
		S4 = GOH3
	} else if InputDacian == Dacian[4] {
		S1 = GMP2
		S2 = GSP1
		S3 = GMH3
		S4 = GOH5
	} else if InputDacian == Dacian[5] {
		S1 = GMP2
		S2 = GSP2
		S3 = GMH2
		S4 = GOH4
	} else if InputDacian == Dacian[6] {
		S1 = GMP2
		S2 = GSP1
		S3 = GMH2
		S4 = GOH5
	} else if InputDacian == Dacian[7] {
		S1 = GMP2
		S2 = GSP1
		S3 = GMH1
		S4 = GOH1
	}
	return
}

func GetNFTGear(InputNonce int64) (S1, S2, S3, S4 string) {
	BSDacian := GetNFTDacian(InputNonce)
	S1, S2, S3, S4 = GetNFTGearFromDacian(BSDacian)
	return
}

func ProcessNonce(InputNonce int64) (NFT NFTBloodShed, Position int64) {
	NFT.Rarity = GetNFTRarity(InputNonce)
	NFT.Dacian = GetNFTDacian(InputNonce)
	NFT.Potency = GetNFTPotency(InputNonce)
	NFT.BloodShedTier = GetNFTBloodshedTier(InputNonce)
	NFT.Background = GetNFTBackground(InputNonce)

	V1, V2, V3, V4 := GetNFTGearFromDacian(NFT.Dacian)
	NFT.MainProtection = V1
	NFT.SecondaryProtection = V2
	NFT.MainHand = V3
	NFT.OffHand = V4

	return
}

func GetBloodShedOrderMultiplier(InputNonce int64) *p.Decimal {
	var OrderMultiplier *p.Decimal
	OrderM := func(RarityRange, RarityNumber, Position int64) *p.Decimal {
		Var1 := Position - 1
		Var2 := RarityNumber - 1
		Var3 := sm.DIVxc(p.NFI(RarityRange), p.NFI(Var2))
		Var4 := sm.MULxc(Var3, p.NFI(Var1))
		OrderPercent := sm.SUBxc(p.NFI(RarityRange), Var4)
		Output1 := sm.ADDxc(p.NFS("1"), sm.DIVxc(OrderPercent, p.NFS("100")))
		Output2 := sm.TruncateCustom(Output1, 18)
		return Output2
	}
	Rarity := GetNFTRarity(InputNonce)
	Position := GetNFTPosition(InputNonce)
	if Rarity == "Legendary" {
		OrderMultiplier = OrderM(100, 152, Position)
	} else if Rarity == "Epic" {
		OrderMultiplier = OrderM(200, 603, Position)
	} else if Rarity == "Rare" {
		OrderMultiplier = OrderM(300, 402, Position)
	} else if Rarity == "Common" {
		OrderMultiplier = OrderM(400, 7704, Position)
	}
	return OrderMultiplier
}
