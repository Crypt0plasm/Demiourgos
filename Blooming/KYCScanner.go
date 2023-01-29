package Blooming

import (
	mvx "MvxApiScanner"
	"strings"
)

var (
	KYCDirectory = "_KYCss-Snapshots\\"
)

//Reads Kyc CSV File

func KycScanner(Path string) []mvx.MvxAddress {
	var (
		Output []mvx.MvxAddress
		PAddy  mvx.MvxAddress
	)

	ProcessLine := func(Line string) mvx.MvxAddress {
		var Addy mvx.MvxAddress
		CheckFirstThree := func(Addy string) bool {
			var Result bool
			Rune := []rune(Addy)
			RuneThree := Rune[:3]
			RuneThreeString := string(RuneThree)
			if RuneThreeString == "erd" {
				Result = true
			}
			return Result
		}

		StripLength := func(Addy string) mvx.MvxAddress {
			var Result mvx.MvxAddress
			Rune := []rune(Addy)
			RuneSixtyTwo := Rune[:62]
			RuneSixtyTwoString := string(RuneSixtyTwo)
			Result = mvx.MvxAddress(RuneSixtyTwoString)
			return Result
		}

		StringSlice := strings.Split(Line, ",")
		if StringSlice[7] == "APPROVED" && CheckFirstThree(StringSlice[5]) == true {
			Addy = StripLength(StringSlice[5])
		}

		return Addy
	}

	StringSlice := mvx.ReadFile(Path)
	for i := 1; i < len(StringSlice); i++ {
		PAddy = ProcessLine(StringSlice[i])
		if string(PAddy) != "" {
			Output = append(Output, PAddy)
		}
	}
	return Output
}

func MakeSetKYC(Input []mvx.BalanceESDT, KYC []mvx.MvxAddress) []mvx.TrueBalanceESDT {
	var (
		Unit     mvx.TrueBalanceESDT
		Output   []mvx.TrueBalanceESDT
		KycValue bool
	)
	for i := 0; i < len(Input); i++ {
		for j := 0; j < len(KYC); j++ {
			if Input[i].Address == KYC[j] {
				KycValue = true
				break
			} else {
				KycValue = false
			}
		}

		Unit.AB = Input[i]
		Unit.KYC = KycValue
		Output = append(Output, Unit)
	}
	return Output
}

func CleanKycSet(Input []mvx.TrueBalanceESDT) []mvx.BalanceESDT {
	var (
		Unit   mvx.BalanceESDT
		Output []mvx.BalanceESDT
	)
	for i := 0; i < len(Input); i++ {
		if Input[i].KYC == true {
			Unit = Input[i].AB
			Output = append(Output, Unit)
		}
	}
	return Output
}
