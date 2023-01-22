package main

import (
	bloom "Demiourgos/Blooming"
	mvx "MvxApiScanner"
	"flag"
)

func ExtractKYC() []mvx.MvxAddress {
	Path := "d:\\_Crypto\\Demiourgos\\Subsidiaries\\Coding Division\\Blooming\\export_416055238_02-2023.csv"
	KycERDs := bloom.KycScanner(Path)
	return KycERDs
}

func main() {
	var (
		KYCSScan = `--kyc  Scans the Downloaded Synapse File
`
	)

	const (
		KYCScan = "kyc"
	)

	FlagKYCScan := flag.Bool(KYCScan, false, KYCSScan)

	flag.Parse()

	if *FlagKYCScan != false {
		KYC := ExtractKYC()
		mvx.WriteChainMvxAddresses("KYC-Addresses.txt", KYC)
	}
}
