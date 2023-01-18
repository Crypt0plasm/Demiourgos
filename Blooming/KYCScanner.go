package Blooming

import (
	mvx "MvxApiScanner"
	"strings"
)

//Reads Kyc CSV File

func KycScanner(Path string) []mvx.MvxAddress {
	var (
		Output []mvx.MvxAddress
		Addy   mvx.MvxAddress
	)

	ProcessLine := func(Line string) mvx.MvxAddress {
		var Adddy mvx.MvxAddress
		StringSlice := strings.Split(Line, ",")
		if StringSlice[7] == "APPROVED" && len([]rune(StringSlice[5])) == 62 {
			Adddy = mvx.MvxAddress(StringSlice[5])
		}
		return Adddy
	}

	StringSlice := mvx.ReadFile(Path)
	for i := 1; i < len(StringSlice); i++ {
		Addy = ProcessLine(StringSlice[i])
		if len([]rune(Addy)) == 62 {
			Output = append(Output, Addy)
		}
	}

	return Output
}

func MakeSetKYC(Input []mvx.BalanceSFT, KYC []mvx.MvxAddress) []mvx.TrueBalanceSFT {
	var (
		Unit     mvx.TrueBalanceSFT
		Output   []mvx.TrueBalanceSFT
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

func CleanKycSet(Input []mvx.TrueBalanceSFT) []mvx.BalanceSFT {
	var (
		Unit   mvx.BalanceSFT
		Output []mvx.BalanceSFT
	)
	for i := 0; i < len(Input); i++ {
		if Input[i].KYC == true {
			Unit = Input[i].AB
			Output = append(Output, Unit)
		}
	}
	return Output
}
