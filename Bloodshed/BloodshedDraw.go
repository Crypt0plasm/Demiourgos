package Bloodshed

import (
	mvx "MvxApiScanner"
	"strconv"
	"strings"
)

func ReadDrawFile(Path string) []string {
	var (
		Output, Unit []string
	)

	ProcessLine := func(Line string) []string {
		StringSlice := strings.Split(Line, " ")
		return StringSlice
	}

	WholeStringSlice := mvx.ReadFile(Path)
	for i := 0; i < len(WholeStringSlice); i++ {
		Unit = ProcessLine(WholeStringSlice[i])
		Output = append(Output, Unit...)
	}

	return Output
}

func MakeBloodshedSendFile(Input []string, Addy mvx.MvxAddress, Collection string) []string {
	var (
		Output []string
		Unit   string
	)
	NFTStrings := mvx.MakeTotalNFTString(8861)

	for i := 0; i < len(Input); i++ {
		Number, _ := strconv.Atoi(Input[i])
		Unit = string(Addy) + ";" + Collection + "-" + NFTStrings[Number]
		Output = append(Output, Unit)
	}

	return Output
}
