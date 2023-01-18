package Blooming

import (
	mvx "MvxApiScanner"
	"encoding/json"
	"fmt"
	"time"
)

func MakeNFTChain2(Collection mvx.SFT, CollectionSize int64) ([]mvx.BalanceNFT, []mvx.BalanceESDT) {
	var (
		NFTOutput []mvx.BalanceNFT
		Full      []mvx.BalanceESDT
		Empty     []mvx.BalanceESDT

		FullUnit  mvx.BalanceESDT
		EmptyUnit mvx.BalanceESDT
		NftUnit   mvx.BalanceNFT
	)

	NFTApiScanner := func(Api string) (Output mvx.MvxAddress) {
		var ScanResult []mvx.BalanceSFT
		SS := mvx.OnPage(Api)
		_ = json.Unmarshal([]byte(SS), &ScanResult)
		return ScanResult[0].Address
	}

	NFTIdentifiersChain := mvx.MakeTotalNFTString(CollectionSize)

	for i := int64(1); i < CollectionSize+1; i++ {
		NFTUniqueIdentifier := NFTIdentifiersChain[i]
		NftUnit.NFT = mvx.NFT{Collection: string(Collection), Identifier: NFTUniqueIdentifier}
		NftUnit.Address = NFTApiScanner(mvx.MakeNFTSnapshotLink(Collection, NFTUniqueIdentifier))
		NFTOutput = append(NFTOutput, NftUnit)

		//Make Dummy SFT List
		FullUnit.Address = NftUnit.Address
		FullUnit.Balance = "1"
		Full = append(Full, FullUnit)

		EmptyUnit.Address = NftUnit.Address
		EmptyUnit.Balance = "0"
		Empty = append(Empty, FullUnit)

		fmt.Println(i, NftUnit)
		time.Sleep(time.Duration(2000) * time.Millisecond)
	}

	SummedNFTList := mvx.DecimalChainAdder(Full, Empty)
	return NFTOutput, SummedNFTList
}
