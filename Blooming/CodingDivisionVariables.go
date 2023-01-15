package Blooming

import (
	"fmt"
	"log"
	"os"
)

var (
	CD01SnakeEye = "https://api.multiversx.com/nfts/DHCD-bc9963-01/accounts?size=10000"
	CD02Rudis    = "https://api.multiversx.com/nfts/DHCD-bc9963-02/accounts?size=10000"
	CD03Gwen     = "https://api.multiversx.com/nfts/DHCD-bc9963-03/accounts?size=10000"
	CD04Clutter  = "https://api.multiversx.com/nfts/DHCD-bc9963-04/accounts?size=10000"
	CD05Bangai   = "https://api.multiversx.com/nfts/DHCD-bc9963-05/accounts?size=10000"
	CD06Binos    = "https://api.multiversx.com/nfts/DHCD-bc9963-06/accounts?size=10000"
	CD07Rubia    = "https://api.multiversx.com/nfts/DHCD-bc9963-07/accounts?size=10000"
	CD08Ocultus  = "https://api.multiversx.com/nfts/DHCD-bc9963-08/accounts?size=10000"
	CD09Oreta    = "https://api.multiversx.com/nfts/DHCD-bc9963-09/accounts?size=10000"
	CD10Binar    = "https://api.multiversx.com/nfts/DHCD-bc9963-0a/accounts?size=10000"

	SCMint          = ElrondAddress("erd1qqqqqqqqqqqqqpgqk7t2cc8awcgwnftnn4p9w7m8fjxplkfcj9qselntcv")
	MarketXoxno     = ElrondAddress("erd1qqqqqqqqqqqqqpgq6wegs2xkypfpync8mn2sa5cmpqjlvrhwz5nqgepyg8")
	VaultSnake      = ElrondAddress("erd1qqqqqqqqqqqqqpgql9zuu0r5pj9xcx94y08wujmwkn2hcuns27rsdcmzf0")
	VaultCD         = ElrondAddress("erd1qqqqqqqqqqqqqpgqqkyp6auk2h7y6sdj2w55qp8zad5ddurn27rs47vs2n")
	MarketFrameIt1  = ElrondAddress("erd1qqqqqqqqqqqqqpgqx00c5udg9uql5sgk5vswfr8cp7kalqgcyawq9xpw26")
	MarketFrameIt2  = ElrondAddress("erd1qqqqqqqqqqqqqpgq705fxpfrjne0tl3ece0rrspykq88mynn4kxs2cg43s")
	MarketNFTr      = ElrondAddress("erd1qqqqqqqqqqqqqpgqz2tgn80j5p5hqh4hx69uc27gz0drcjmmg20skf0wru")
	Krogan          = ElrondAddress("erd1qqqqqqqqqqqqqpgq8xwzu82v8ex3h4ayl5lsvxqxnhecpwyvwe0sf2qj4e")
	Hefe            = ElrondAddress("erd1vj40fxw0yah34mmdxly7l28w097ju6hf8pczpcdxs05n2vyx8hcspyxm2c")
	KosonicTreasury = ElrondAddress("erd1h0ymqdgl6vf0pud0klz5nstwra3sxj06afaj86x0pg7p52dvve9qqtg7x4")
	VestaMinter     = ElrondAddress("erd1qqqqqqqqqqqqqpgqtwe67hk3rwpjx2rky74csj069gftw32j2d2ssd2mvf")

	ExA1 = SCMint         //Mint SC
	ExA2 = MarketXoxno    //Xoxno SC
	ExA3 = VaultSnake     //Snakes Vault
	ExA4 = VaultCD        //CD Vault
	ExA5 = MarketFrameIt1 //market FrameItLot??
	ExA6 = MarketFrameIt2 //market Frameit
	ExA7 = MarketNFTr     //market NFTr
	ExA8 = Krogan         //market Krogan
	ExA9 = VestaMinter    //Vesta Miner

	ExAm1 = AddressExceptions{Hefe, "50"}

	SCExceptions     = []ElrondAddress{ExA1, ExA2, ExA3, ExA4, ExA5, ExA6, ExA7, ExA8, ExA9}
	AmountExceptions = []AddressExceptions{ExAm1}
	SetExceptions    = []ElrondAddress{KosonicTreasury}
)

func ComputeExceptionAddress(Addy ElrondAddress) bool {
	var Result = false
	for i := 0; i < len(SCExceptions); i++ {
		if Addy == SCExceptions[i] {
			Result = true
		}
	}
	return Result
}

type AddressExceptions struct {
	Address ElrondAddress
	Amount  string
}

type NFT struct {
	Collection string
	Identifier string
}

type ElrondAddress string

type BalanceNFT struct {
	Address ElrondAddress
	NFT     NFT
}

type BalanceSFT struct {
	Address ElrondAddress `json:"address"`
	Balance string        `json:"balance"`
}

type BalanceESDT struct {
	Address ElrondAddress `json:"address"`
	Balance string        `json:"balance"`
}

type TrueBalanceSFT struct {
	AB  BalanceSFT
	KYC bool
}

type DecaChain struct {
	SnakeEye []BalanceSFT
	Rudis    []BalanceSFT
	Gwen     []BalanceSFT
	Clutter  []BalanceSFT
	Bangai   []BalanceSFT
	Binos    []BalanceSFT
	Rubia    []BalanceSFT
	Ocultus  []BalanceSFT
	Oreta    []BalanceSFT
	Binar    []BalanceSFT
}

// ================================================
//
// # Function 07.06 - WriteList
//
// WriteList writes the strings from the slice to an external file
// as Name can be used "File.txt" as the output file.
func WriteListOneByOne(Name string, List []BalanceSFT) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}

func WriteListOneByOneB(Name string, List []BalanceESDT) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}

func WriteListOneByOne2(Name string, List []TrueBalanceSFT) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}

func WriteListOneAfterTheOther(Name string, List []BalanceSFT) {
	//Printing Snapshot
	//fmt.Println("Outputting SFT Balance to", Name)
	OutputFile, err := os.Create(Name)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()
	_, _ = fmt.Fprintln(OutputFile, List)
	//fmt.Println("DONE Outputting SFT Balance Amounts to", Name)
	//fmt.Println("")
}
