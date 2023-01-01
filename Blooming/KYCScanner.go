package Blooming

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(Path string) []string {
	readFile, err := os.Open(Path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	_ = readFile.Close()

	return fileLines
}

func ProcessLine(Line string) ElrondAddress {
	var Addy ElrondAddress
	StringSlice := strings.Split(Line, ",")
	if StringSlice[7] == "APPROVED" && len([]rune(StringSlice[5])) == 62 {
		Addy = ElrondAddress(StringSlice[5])
	}
	return Addy
}

func KycScanner(Path string) []ElrondAddress {
	var (
		Output []ElrondAddress
		Addy   ElrondAddress
	)
	StringSlice := ReadFile(Path)
	for i := 1; i < len(StringSlice); i++ {
		Addy = ProcessLine(StringSlice[i])
		if len([]rune(Addy)) == 62 {
			Output = append(Output, Addy)
		}
	}

	return Output
}

func MakeSetKYC(Input []BalanceSFT, KYC []ElrondAddress) []TrueBalanceSFT {
	var (
		Unit     TrueBalanceSFT
		Output   []TrueBalanceSFT
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

func CleanKycSet(Input []TrueBalanceSFT) []BalanceSFT {
	var (
		Unit   BalanceSFT
		Output []BalanceSFT
	)
	for i := 0; i < len(Input); i++ {
		if Input[i].KYC == true {
			Unit = Input[i].AB
			Output = append(Output, Unit)
		}
	}
	return Output
}
