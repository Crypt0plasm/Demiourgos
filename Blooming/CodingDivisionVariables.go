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
	ExA13 = mvx.DHV4

	ExA14 = mvx.MvxAddress("erd1j43ssgjr8x9m0s7zz0xhn7qdgssdetm86jc6a2dzsegs74fmyl5ssv44c4") //Alex Snake Hold, holds the remaining 200 Snakes at 5000 Dolars
	ExA15 = mvx.MvxAddress("erd1qqqqqqqqqqqqqpgqrujrjjnaeqc3srdqn8vtzz8wh5dmnq9xyl5s8797wk") //Last 40 Snakes at 2500 Dollars
	ExA16 = mvx.MvxAddress("erd1qqqqqqqqqqqqqpgq4u78a0fk85lhye9rj4w3dxgn90ke9frc03aszeuwcu")

	ExA17 = mvx.MvxAddress("erd1qqqqqqqqqqqqqpgquqwc8v09e5pmcz9h4569gynle8qwjdenyl5sayfsl0") //BS Original Lottery SC
	ExA18 = mvx.MvxAddress("erd1qqqqqqqqqqqqqpgqcc2dakhdz23hk8gvlnn054uhzzeewn5xwmfsyqdssd") //BS Second Mint sc

	//This is the Amount Exception
	//Paul holds 50 Company SFTs that aren't Include in the computation, that are excepted from computation
	ExAm1 = AddressINTExceptions{mvx.Hefe, 0}

	//Smart Contract Exceptions
	VestaExceptions = []mvx.MvxAddress{ExA1, ExA2, ExA3, ExA4, ExA6, ExA7, ExA8, ExA9, ExA10, ExA11, ExA12, ExA13}
	CDExceptions    = []mvx.MvxAddress{ExA1, ExA2, ExA3, ExA4, ExA5, ExA6, ExA7, ExA8, ExA9, ExA10, ExA11, ExA12, ExA13}
	SnakeExceptions = []mvx.MvxAddress{ExA3, ExA4, ExA5, ExA6, ExA7, ExA8, ExA9, ExA10, ExA11, ExA12, ExA13, ExA14, ExA15, ExA16}

	//Amount Exceptions
	AmountExceptions = []AddressINTExceptions{ExAm1}

	//Set Exceptions
	SetExceptions = []mvx.MvxAddress{mvx.KosonicTreasury}
)

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
