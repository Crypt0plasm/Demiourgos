package Blooming

import (
	p "Firefly-APD"
	mt "SuperMath"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func MakeBaseString() []string {
	var (
		SingleString string
		StringSlice  []string
	)
	for i := int64(0); i < 16; i++ {
		SingleString = strconv.FormatInt(i, 16)
		StringSlice = append(StringSlice, SingleString)
	}
	return StringSlice
}

func Make2DigitsString(Number int64) []string {
	var (
		String       string
		OutputString []string
	)
	BS := MakeBaseString()

	if Number >= 0 && Number < 16 {
		SS := strconv.FormatInt(Number, 16)
		for i := 0; i < len(BS); i++ {
			String = SS + BS[i]
			OutputString = append(OutputString, String)
		}
	} else {
		return BS
	}
	return OutputString
}

func Make256String() []string {
	var (
		SingleStringChain []string
		String256         []string
	)
	for i := int64(0); i < 16; i++ {
		SingleStringChain = Make2DigitsString(i)
		String256 = append(String256, SingleStringChain...)
	}
	return String256
}

func MakeTotalString(Size int64) []string {
	var (
		Output   []string
		Addition string
		ToAppend []string
	)

	AppendedStringChain := func(Prefix string, StringChain []string) []string {
		var (
			Unit   string
			Result []string
		)
		for i := 0; i < len(StringChain); i++ {
			Unit = Prefix + StringChain[i]
			Result = append(Result, Unit)
		}
		return Result
	}
	GetPrefix := func(Number int) string {
		var Prefix string
		if Number < 10 {
			Prefix = "0" + strconv.Itoa(Number)
		} else {
			Prefix = strconv.Itoa(Number)
		}
		return Prefix
	}

	Whole := mt.DivInt(p.NFI(Size), p.NFS("256"))
	Rest := mt.DivMod(p.NFI(Size), p.NFS("256"))
	WholeInt, _ := strconv.Atoi(mt.DTS(Whole))
	RestInt, _ := strconv.Atoi(mt.DTS(Rest))

	String256 := Make256String()

	if WholeInt == 0 {
		Output = String256[:RestInt+1]
	} else {
		for i := 0; i < WholeInt+1; i++ {
			if i == 0 {
				Output = append(Output, String256...)
			} else if i >= 1 && i < WholeInt {
				Addition = GetPrefix(i)
				ToAppend = AppendedStringChain(Addition, String256)
				Output = append(Output, ToAppend...)
			} else if i == WholeInt {
				Addition = GetPrefix(i)
				ToAppend = AppendedStringChain(Addition, String256[:RestInt+1])
				Output = append(Output, ToAppend...)
			}
		}
	}
	return Output
}

func MakeNFTChain(Col string, Unit int64) (Output []BalanceNFT) {
	var (
		NFTBasePrefix = "https://api.elrond.com/nfts/"
		NFTBaseSufix  = "/owners?size=1"
		NftUnit       BalanceNFT
		//Output        []BalanceNFT
	)

	ApiScanner := func(Api string) (Output []BalanceSFT) {
		SS := OnPage(Api)
		_ = json.Unmarshal([]byte(SS), &Output)
		return
	}

	NFTIdentifierChain := MakeTotalString(Unit)

	for i := int64(1); i < Unit+1; i++ {
		NFTIdentifier := NFTIdentifierChain[i]
		NftUnit.NFT = NFT{Col, NFTIdentifier}

		ApiPingString := NFTBasePrefix + Col + "-" + NFTIdentifier + NFTBaseSufix
		time.Sleep(time.Duration(1000) * time.Millisecond)
		fmt.Println("API:", ApiPingString)
		NftUnit.Address = ApiScanner(ApiPingString)[0].Address

		//NftUnit.AB = IterationSFT
		//NftUnit.NFT = IterationNFT
		Output = append(Output, NftUnit)
		fmt.Println(i, NftUnit)
		time.Sleep(time.Duration(1000) * time.Millisecond)
		//Output = append(Output, NftUnit)
	}
	return
}
