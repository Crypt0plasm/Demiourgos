package Blooming

import (
	mvx "MvxApiScanner"
)

var (
	ExA1 = mvx.CodingDivisionMintSC //Mint SC
	ExA2 = mvx.MarketXoxno          //Xoxno SC
	ExA3 = mvx.SupercietyVaultSnake //Snakes Vault
	ExA4 = mvx.SupercietyVaultCD    //CD Vault
	ExA5 = mvx.MarketFrameIt1       //market FrameItLot??
	ExA6 = mvx.MarketFrameIt2       //market Frameit
	ExA7 = mvx.MarketNFTr           //market NFTr
	ExA8 = mvx.Krogan               //market Krogan
	ExA9 = mvx.VestaMinter          //Vesta Miner

	//This is the Amount Exception
	//Paul holds 50 Company SFTs that aren't Include in the computation, that are excepted from computation
	ExAm1 = AddressINTExceptions{mvx.Hefe, 50}

	//Smart Contract Exceptions
	SCExceptions = []mvx.MvxAddress{ExA1, ExA2, ExA3, ExA4, ExA5, ExA6, ExA7, ExA8, ExA9}

	//Amount Exceptions
	AmountExceptions = []AddressINTExceptions{ExAm1}

	//Set Exceptions
	SetExceptions = []mvx.MvxAddress{mvx.KosonicTreasury}
)

func ComputeExceptionAddress(Addy mvx.MvxAddress) bool {
	var Result = false
	for i := 0; i < len(SCExceptions); i++ {
		if Addy == SCExceptions[i] {
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
