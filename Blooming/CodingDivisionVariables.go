package Blooming

import (
	mvx "MvxApiScanner"
)

var (
	ExA1  = mvx.CodingDivisionMintSC //Mint SC
	ExA2  = mvx.VestaMinter          //Vesta Miner
	ExA3  = mvx.MarketXoxno          //Xoxno SC
	ExA4  = mvx.SnakeDAO             //Snakes Vault
	ExA5  = mvx.CodingDivisionDAO    //CD Vault
	ExA6  = mvx.MarketFrameIt1       //market FrameItLot??
	ExA7  = mvx.MarketFrameIt2       //market Frameit
	ExA8  = mvx.MarketNFTr           //market NFTr
	ExA9  = mvx.Krogan               //market Krogan
	ExA10 = mvx.DHV1
	ExA11 = mvx.DHV2
	ExA12 = mvx.DHV3

	//This is the Amount Exception
	//Paul holds 50 Company SFTs that aren't Include in the computation, that are excepted from computation
	ExAm1 = AddressINTExceptions{mvx.Hefe, 16}

	//Smart Contract Exceptions
	VestaExceptions = []mvx.MvxAddress{ExA1, ExA2, ExA3, ExA4, ExA6, ExA7, ExA8, ExA9, ExA11, ExA12}
	CDExceptions    = []mvx.MvxAddress{ExA1, ExA2, ExA3, ExA4, ExA5, ExA6, ExA7, ExA8, ExA9, ExA10, ExA11, ExA12}

	//Amount Exceptions
	AmountExceptions = []AddressINTExceptions{ExAm1}

	//Set Exceptions
	SetExceptions = []mvx.MvxAddress{mvx.KosonicTreasury}
)

func ComputeExceptionAddress(Addy mvx.MvxAddress, ExceptionList []mvx.MvxAddress) bool {
	var Result = false
	for i := 0; i < len(ExceptionList); i++ {
		if Addy == ExceptionList[i] {
			Result = true
		}
	}
	return Result
}

type AddressINTExceptions struct {
	Address mvx.MvxAddress
	Amount  int
}

type DecaChain struct {
	SnakeEye []mvx.BalanceSFT
	Rudis    []mvx.BalanceSFT
	Gwen     []mvx.BalanceSFT
	Clutter  []mvx.BalanceSFT
	Bangai   []mvx.BalanceSFT
	Binos    []mvx.BalanceSFT
	Rubia    []mvx.BalanceSFT
	Ocultus  []mvx.BalanceSFT
	Oreta    []mvx.BalanceSFT
	Binar    []mvx.BalanceSFT
}
