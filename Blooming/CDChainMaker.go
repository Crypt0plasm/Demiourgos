package Blooming

import (
	p "Firefly-APD"
	mvx "MvxApiScanner"
	mt "SuperMath"
	"fmt"
)

// =====================================================================================================================
//
//	Blooming/CDChainMaker.go
//	Creates Chains with Coding Division SFT Values.
//
//
// [B]01         CreateCodingDivisionChain       	Create a Chain of Character Chain Values
// [B]02         CreateCodingDivisionOwners      	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
// [B]03         GetAddressAmount                	Gets the Amount for a given address in a Balance Chain
// [B]04         CreateCodingDivisionSetChain    	Create a Chain of all Elrond Addresses Containing Coding Division SFTs Sets
// [B]05         CreateCodingDivisionAmountChain 	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
// [B]06         SortBalanceSFTChain             	Sorts a BalanceSFTChain Chain from the highest Balance to the lowest Balance
// [B]07         AddBalanceSFTChain              	Add the Balances of an SFT Chain
// =====================================================================================================================
// =====================================================================================================================

func CreateSnakeChain(Input mvx.SFT) ([]mvx.BalanceSFT, *p.Decimal) {
	fmt.Println("Snapshotting Snake SFT Addresses and Amounts")
	SnakeChain := mvx.SnapshotSFTChain(Input)
	Sum := mvx.AddBalanceIntegerChain(SnakeChain)
	fmt.Println(len(SnakeChain), "addresses snapshotted with Snake SFTs", Sum)
	fmt.Println("")
	return SnakeChain, Sum
}

func CreateSnakeAmountChain(InputChain []mvx.BalanceSFT) []mvx.BalanceSFT {
	var (
		AllChain []mvx.BalanceSFT
		Unit     mvx.BalanceSFT
	)
	for i := 0; i < len(InputChain); i++ {
		if ComputeExceptionAddress(InputChain[i].Address, SnakeExceptions) == false {
			Unit.Address = InputChain[i].Address
			Unit.Balance = InputChain[i].Balance

			AllChain = append(AllChain, Unit)
		}
	}
	return AllChain
}

func CreateVestaChain(Input mvx.SFT) ([]mvx.BalanceSFT, *p.Decimal) {
	var VestaType string
	if Input == mvx.VestaGold {
		VestaType = "Golden"
	} else if Input == mvx.VestaSilver {
		VestaType = "Silver"
	} else {
		VestaType = "Bronze"
	}
	fmt.Println("Snapshotting ", VestaType, " Vesta Addresses and Amounts")
	VestaChain := mvx.SnapshotSFTChain(Input)
	Sum := mvx.AddBalanceIntegerChain(VestaChain)
	fmt.Println(len(VestaChain), "addresses snapshotted with Vesta SFTs", Sum)
	fmt.Println("")
	return VestaChain, Sum
}

func CreateVestaAmountChains(InputChain []mvx.BalanceSFT) []mvx.BalanceSFT {
	var (
		AllChain []mvx.BalanceSFT
		Unit     mvx.BalanceSFT
	)
	for i := 0; i < len(InputChain); i++ {
		if ComputeExceptionAddress(InputChain[i].Address, VestaExceptions) == false {
			Unit.Address = InputChain[i].Address
			Unit.Balance = InputChain[i].Balance

			AllChain = append(AllChain, Unit)
		}
	}
	return AllChain
}

// CreateCodingDivisionChain ===========================================================================================
//
// [B]01         CreateCodingDivisionChain
//
//	Create a Chain of all Character Chain Values
//	DecaChain Is a Chain of 10 Chains - one for each CD Character
//	int represents its length
//	*p.Decimal represents their Total Sum (that is the total raw sum of all scanned SFTs)
func CreateCodingDivisionChain() (DecaChain, int, *p.Decimal) {
	var OutputChain DecaChain

	fmt.Println("Snapshotting SnakeEye Addresses and Amounts")
	SnakeEyeChain := mvx.SnapshotSFTChain(mvx.CD01SnakeEye)
	S01 := mvx.AddBalanceIntegerChain(SnakeEyeChain)
	fmt.Println(len(SnakeEyeChain), "addresses snapshotted with Snake Eye SFTs", S01)
	fmt.Println("")

	fmt.Println("Snapshotting Rudis Addresses and Amounts")
	RudisChain := mvx.SnapshotSFTChain(mvx.CD02Rudis)
	S02 := mvx.AddBalanceIntegerChain(RudisChain)
	fmt.Println(len(RudisChain), "addresses snapshotted with Rudis SFTs", S02)
	fmt.Println("")

	fmt.Println("Snapshotting Gwen Addresses and Amounts")
	GwenChain := mvx.SnapshotSFTChain(mvx.CD03Gwen)
	S03 := mvx.AddBalanceIntegerChain(GwenChain)
	fmt.Println(len(GwenChain), "addresses snapshotted with Gwen SFTs", S03)
	fmt.Println("")

	fmt.Println("Snapshotting Clutter Addresses and Amounts")
	ClutterChain := mvx.SnapshotSFTChain(mvx.CD04Clutter)
	S04 := mvx.AddBalanceIntegerChain(ClutterChain)
	fmt.Println(len(ClutterChain), "addresses snapshotted with Clutter SFTs", S04)
	fmt.Println("")

	fmt.Println("Snapshotting Bangai Addresses and Amounts")
	BangaiChain := mvx.SnapshotSFTChain(mvx.CD05Bangai)
	S05 := mvx.AddBalanceIntegerChain(BangaiChain)
	fmt.Println(len(BangaiChain), "addresses snapshotted with Bangai SFTs", S05)
	fmt.Println("")

	fmt.Println("Snapshotting Binos Addresses and Amounts")
	BinosChain := mvx.SnapshotSFTChain(mvx.CD06Binos)
	S06 := mvx.AddBalanceIntegerChain(BinosChain)
	fmt.Println(len(BinosChain), "addresses snapshotted with Binos SFTs", S06)
	fmt.Println("")

	fmt.Println("Snapshotting Rubia Addresses and Amounts")
	RubiaChain := mvx.SnapshotSFTChain(mvx.CD07Rubia)
	S07 := mvx.AddBalanceIntegerChain(RubiaChain)
	fmt.Println(len(RubiaChain), "addresses snapshotted with Rubia SFTs", S07)
	fmt.Println("")

	fmt.Println("Snapshotting Ocultus Addresses and Amounts")
	OcultusChain := mvx.SnapshotSFTChain(mvx.CD08Ocultus)
	S08 := mvx.AddBalanceIntegerChain(OcultusChain)
	fmt.Println(len(OcultusChain), "addresses snapshotted with Ocultus SFTs", S08)
	fmt.Println("")

	fmt.Println("Snapshotting Oreta Addresses and Amounts")
	OretaChain := mvx.SnapshotSFTChain(mvx.CD09Oreta)
	S09 := mvx.AddBalanceIntegerChain(OretaChain)
	fmt.Println(len(OretaChain), "addresses snapshotted with Oreta SFTs", S09)
	fmt.Println("")

	fmt.Println("Snapshotting Binar Addresses and Amounts")
	BinarChain := mvx.SnapshotSFTChain(mvx.CD10Binar)
	S10 := mvx.AddBalanceIntegerChain(BinarChain)
	fmt.Println(len(BinarChain), "addresses snapshotted with Binar SFTs", S10)
	fmt.Println("")

	Len1 := len(SnakeEyeChain) + len(RudisChain) + len(GwenChain) + len(ClutterChain) + len(BangaiChain)
	Len2 := len(BinosChain) + len(RubiaChain) + len(OcultusChain) + len(OretaChain) + len(BinarChain)
	MaxLen := Len1 + Len2

	Sum := mt.SUMxc(S01, S02, S03, S04, S05, S06, S07, S08, S09, S10)

	OutputChain.SnakeEye = SnakeEyeChain
	OutputChain.Rudis = RudisChain
	OutputChain.Gwen = GwenChain
	OutputChain.Clutter = ClutterChain
	OutputChain.Bangai = BangaiChain
	OutputChain.Binos = BinosChain
	OutputChain.Rubia = RubiaChain
	OutputChain.Ocultus = OcultusChain
	OutputChain.Oreta = OretaChain
	OutputChain.Binar = BinarChain

	return OutputChain, MaxLen, Sum
}

// CreateCodingDivisionOwners ==========================================================================================
//
// [B]02         CreateCodingDivisionOwners
//
//	Create a Chain of all Mvx Addresses Containing Coding Division SFTs
func CreateCodingDivisionOwners(InputChain DecaChain) []mvx.MvxAddress {

	//1.    Make a slice of all the Slices in the DecaChain
	AllSlice1 := append(InputChain.SnakeEye, InputChain.Rudis...)
	AllSlice2 := append(AllSlice1, InputChain.Gwen...)
	AllSlice3 := append(AllSlice2, InputChain.Clutter...)
	AllSlice4 := append(AllSlice3, InputChain.Bangai...)
	AllSlice5 := append(AllSlice4, InputChain.Binos...)
	AllSlice6 := append(AllSlice5, InputChain.Rubia...)
	AllSlice7 := append(AllSlice6, InputChain.Ocultus...)
	AllSlice8 := append(AllSlice7, InputChain.Oreta...)
	AllSlice := append(AllSlice8, InputChain.Binar...)

	//2.    Make a slice with all Mvx Address (will contain duplicate Elrond Addresses.)
	//      basically removes the balance value.
	ElrondSlice := make([]mvx.MvxAddress, len(AllSlice))

	for i := 0; i < len(AllSlice); i++ {
		ElrondSlice[i] = AllSlice[i].Address
	}

	return mvx.RemoveDuplicateMvxAddresses(ElrondSlice)
}

// CreateCodingDivisionSetChain ========================================================================================
//
// [B]03         CreateCodingDivisionSetChain
//
//	Create a Chain of all Mvx Addresses Containing Coding Division SFTs Sets and applies SC Exception
func CreateCodingDivisionSetChain(Owners []mvx.MvxAddress, Snapshot DecaChain) []mvx.BalanceSFT {
	var (
		SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar *p.Decimal
		Minimum                                                                     *p.Decimal
		SetChain                                                                    []mvx.BalanceSFT
		Unit                                                                        mvx.BalanceSFT
	)

	for i := 0; i < len(Owners); i++ {
		SnakeEye = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.SnakeEye))
		Rudis = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Rudis))
		Gwen = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Gwen))
		Clutter = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Clutter))
		Bangai = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Bangai))
		Binos = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Binos))
		Rubia = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Rubia))
		Ocultus = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Ocultus))
		Oreta = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Oreta))
		Binar = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Binar))

		Minimum = mt.MinDecimal(SnakeEye, Rudis)
		Minimum = mt.MinDecimal(Minimum, Gwen)
		Minimum = mt.MinDecimal(Minimum, Clutter)
		Minimum = mt.MinDecimal(Minimum, Bangai)
		Minimum = mt.MinDecimal(Minimum, Binos)
		Minimum = mt.MinDecimal(Minimum, Rubia)
		Minimum = mt.MinDecimal(Minimum, Ocultus)
		Minimum = mt.MinDecimal(Minimum, Oreta)
		Minimum = mt.MinDecimal(Minimum, Binar)

		//If Minimum is greater than zero AND Address is non Exception
		if mt.DecimalGreaterThan(Minimum, p.NFS("0")) == true && ComputeExceptionAddress(Owners[i], CDExceptions) == false {
			Unit.Address = Owners[i]
			Unit.Balance = mt.DTS(Minimum)

			SetChain = append(SetChain, Unit)
		}
	}
	return SetChain
}

// CreateCodingDivisionSetExceptionChain ========================================================================================
//
// [B]04         CreateCodingDivisionSetChain
//
//	Create a Chain of all Mvx Addresses Containing Coding Division SFTs Sets,
//	Uses as Input a Chain of Raw Sets (computed above)
//	While Excluding the Chain of Set Exceptions.
func CreateCodingDivisionSetExceptionChain(Input []mvx.BalanceSFT, UseException bool) []mvx.BalanceSFT {
	var (
		Output    []mvx.BalanceSFT
		Unit      mvx.BalanceSFT
		Exception bool
	)

	if UseException == true {
		for i := 0; i < len(Input); i++ {
			Unit.Address = Input[i].Address

			for j := 0; j < len(SetExceptions); j++ {
				if Unit.Address == SetExceptions[j] {
					Exception = true
				} else {
					Exception = false
				}
			}

			if Exception == false {
				Unit.Balance = Input[i].Balance
			}

			Output = append(Output, Unit)
		}
	} else if UseException == false {
		Output = Input
	}

	return Output
}

// CreateCodingDivisionAmountChain ========================================================================================
//
// [B]05         CreateCodingDivisionAmountChain
//
//	Create a Chain of all Mvx Addresses Containing ALL CD SFTs and their amounts and applies SC Exception
func CreateCodingDivisionAmountChain(Owners []mvx.MvxAddress, Snapshot DecaChain) []mvx.BalanceSFT {
	var (
		SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar *p.Decimal
		Sum                                                                         *p.Decimal
		AllChain                                                                    []mvx.BalanceSFT
		Unit                                                                        mvx.BalanceSFT
	)

	for i := 0; i < len(Owners); i++ {
		SnakeEye = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.SnakeEye))
		Rudis = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Rudis))
		Gwen = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Gwen))
		Clutter = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Clutter))
		Bangai = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Bangai))
		Binos = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Binos))
		Rubia = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Rubia))
		Ocultus = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Ocultus))
		Oreta = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Oreta))
		Binar = p.NFS(mvx.GetAddressIntegerAmount(Owners[i], Snapshot.Binar))

		Sum = mt.SUMxc(SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar)

		//If Sum is greater than zero AND Address is non Exception
		if mt.DecimalGreaterThan(Sum, p.NFS("0")) == true && ComputeExceptionAddress(Owners[i], CDExceptions) == false {
			Unit.Address = Owners[i]
			Unit.Balance = mt.DTS(Sum)

			AllChain = append(AllChain, Unit)
		}
	}
	return AllChain
}

// CreateCodingDivisionAmountExceptionChain ========================================================================================
//
// [B]06         CreateCodingDivisionSetChain
//
//	Create a Chain of all Mvx Addresses Containing ALL Coding Division SFT
//	Uses as Input a Chain of Raw Amounts (computed above)
//	While Excluding the Chain of AmountExceptions.
func CreateCodingDivisionAmountExceptionChain(Input []mvx.BalanceSFT, UseException bool) []mvx.BalanceSFT {
	var (
		Unit           mvx.BalanceSFT
		Output         []mvx.BalanceSFT
		Exception      bool
		ExceptionValue int
	)

	if UseException == true {
		for i := 0; i < len(Input); i++ {
			Unit.Address = Input[i].Address

			for j := 0; j < len(AmountExceptions); j++ {
				if Unit.Address == AmountExceptions[j].Address {
					Exception = true
					ExceptionValue = AmountExceptions[j].Amount
				} else {
					Exception = false
				}
			}

			if Exception == true {
				//Input[i] Balance = int
				//ExceptionValue = int
				//UnitBalance = int
				Unit.Balance = mt.DTS(mt.SUBxc(p.NFS(Input[i].Balance), p.NFI(int64(ExceptionValue))))
			} else {
				Unit.Balance = Input[i].Balance
			}

			Output = append(Output, Unit)
		}
	} else if UseException == false {
		Output = Input
	}
	return Output
}
