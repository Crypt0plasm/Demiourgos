package Vesta

import (
	p "Firefly-APD"
	mvx "MvxApiScanner"
	sm "SuperMath"
	"fmt"
)

type VestaHoldings struct {
	Address mvx.MvxAddress
	Gold    int64
	Silver  int64
	Bronze  int64
}

type VestaLPHoldings struct {
	Address  mvx.MvxAddress
	VestaLiq VestaLPs
}

type VestaLPs struct {
	Gold    *p.Decimal
	Silver  *p.Decimal
	Bronze  *p.Decimal
	UGold   *p.Decimal
	USilver *p.Decimal
	UBronze *p.Decimal
}

type VestaLPDesignation struct {
	Gold    string
	Silver  string
	Bronze  string
	UGold   string
	USilver string
	UBronze string
}

type FarmMx struct {
	DEB *p.Decimal
	VLM *p.Decimal
	TM  *p.Decimal
	D   *p.Decimal
	UM  *p.Decimal
	IM  *p.Decimal
	DM  *p.Decimal
	VM  *p.Decimal
}

var (
	AncientDEB = p.NFS("2.5")
	Neutral    = p.NFS("1")
	VestaTM    = p.NFS("5")
	VestaUM    = p.NFS("4.3475")
	Zero       = p.NFS("0")
	Empty      = "empty"

	AncientHodler = mvx.AH
	TrDaniel      = mvx.MvxAddress("erd1hg9q84tyzxretw2a8nce6q3lwgfzku587ndwr5k7202xt5pw4vyqp76vxe")
	DRX           = mvx.MvxAddress("erd1gqh79mxfr4al0gfpplvm79lxrx93tgclm3kvlfklmfgr3sp3yx8s7d3qjp")
	Patryx        = mvx.MvxAddress("erd1k64gwm43pqtydv978a5h27zfs48av7hq79j6e6uxz3dlg704vdyq6kuzym")
	Lavinia       = mvx.MvxAddress("erd1xjutp9aj4amd8r5mtedl4aad6jrxv2ajzw5d4x0c3rkwj7cxsgmssdjr49")
	Sandu         = mvx.MvxAddress("erd1pkxygrm9dswrludcxjk9hqmep62gutz7vfknlfpwzufxyara27sqjyhf4h")
	Cuciorva      = mvx.MvxAddress("erd1knx4hu2u6zxyt2mqk9zfyf8m9pz980puw98hqsmv26d3eggxvfmsv4xlpq")
	Codarcea      = mvx.MvxAddress("erd1mpcp5jvkm2y5cxyjj99esfhgk99776r740yrxr6xm0taamqjw4hqxm2q8d")
	Pulecs        = mvx.MvxAddress("erd1wa7h9ea30j0enjm8k9x8ymf0s334qde8x4c3fpzkgj28xfkyt7nqdttpst")
	Laurentiu     = mvx.MvxAddress("erd1mf2wtc4jh2zujhu4nynvrpaua9e98c4lwdyjnpp57qcx2gyqfy6q8l8ccp")
	Frostedk9     = mvx.MvxAddress("erd1ktu3qy5ehe42sk6z7ygfvwna6wull2suq49qj8urx9nd6dw79s2qn5qqea")
	IonutDRD      = mvx.MvxAddress("erd1ez8ww48xj9gr6yyzem7g7gdvknjdh3te7lcyanz9tkrzyazyzxeqzssrhx")
	Buhaici       = mvx.MvxAddress("erd18n5y3884pdkeq6gl0wng22z2yqexhemwkyewcksesn2vgmgcqxes0e93cc")
	TheKid        = mvx.MvxAddress("erd1zl890854dweghll9faf67ft26965v8u20d6z63cntr9a5ykhcekqmyzcsf")
	RaulTM        = mvx.MvxAddress("erd1640c9n2cck2326jy0tg87nawhgxdxvzqg9psezg2yjzkxvg6gkcqzfdp6j")
	MakeAStep     = mvx.MvxAddress("erd1heus28d80kkengfz4ltn2m9xhvd4th8pajrsg9hkeeu202cjpfwq03m38w")

	VestaVault = FarmMx{AncientDEB, Neutral, VestaTM, Neutral, VestaUM, Neutral, Neutral, Neutral}

	//Users
	UserChain = []VestaHoldings{User000,
		User001, User002, User003, User004, User005, User006, User007, User008, User009, User010,
		User011, User012, User013, User014, User015}

	User000 = VestaHoldings{AncientHodler, 35, 151, 590}
	User001 = VestaHoldings{TrDaniel, 16, 0, 32}
	User002 = VestaHoldings{DRX, 9, 3, 10}
	User003 = VestaHoldings{Patryx, 5, 6, 5}
	User004 = VestaHoldings{Lavinia, 1, 6, 2}
	User005 = VestaHoldings{Sandu, 4, 21, 0}
	User006 = VestaHoldings{Cuciorva, 0, 3, 45}
	User007 = VestaHoldings{Codarcea, 3, 3, 0}
	User008 = VestaHoldings{Pulecs, 7, 10, 0}
	User009 = VestaHoldings{Laurentiu, 11, 1, 0}
	User010 = VestaHoldings{Frostedk9, 18, 2, 0}
	User011 = VestaHoldings{IonutDRD, 0, 0, 8}
	User012 = VestaHoldings{Buhaici, 0, 0, 2}
	User013 = VestaHoldings{TheKid, 0, 0, 45}
	User014 = VestaHoldings{RaulTM, 0, 0, 9}
	User015 = VestaHoldings{MakeAStep, 1, 5, 20}

	//Liquidity From Users
	LiquidityUserChain = []VestaLPHoldings{VLQUser000,
		VLQUser001, VLQUser002, VLQUser003, VLQUser004, VLQUser005, VLQUser006, VLQUser007, VLQUser008, VLQUser009, VLQUser010,
		VLQUser011, VLQUser012, VLQUser013, VLQUser014, VLQUser015}
	LQDEmpty = VestaLPs{Zero, Zero, Zero, Zero, Zero, Zero}
	DSGEmpty = VestaLPDesignation{Empty, Empty, Empty, Empty, Empty, Empty}

	VLQUser000 = VestaLPHoldings{AncientHodler, LQDUser000}
	LQDUser000 = VestaLPs{p.NFS("24625.015212261015287848"), Zero, Zero, p.NFS("0.3242"), Zero, Zero}
	//
	VLQUser001 = VestaLPHoldings{TrDaniel, LQDUser001}
	LQDUser001 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser002 = VestaLPHoldings{DRX, LQDUser002}
	LQDUser002 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser003 = VestaLPHoldings{Patryx, LQDUser003}
	LQDUser003 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser004 = VestaLPHoldings{Lavinia, LQDUser004}
	LQDUser004 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser005 = VestaLPHoldings{Sandu, LQDUser005}
	LQDUser005 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser006 = VestaLPHoldings{Cuciorva, LQDUser006}
	LQDUser006 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser007 = VestaLPHoldings{Codarcea, LQDUser007}
	LQDUser007 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser008 = VestaLPHoldings{Pulecs, LQDUser008}
	LQDUser008 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser009 = VestaLPHoldings{Laurentiu, LQDUser009}
	LQDUser009 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser010 = VestaLPHoldings{Frostedk9, LQDUser010}
	LQDUser010 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser011 = VestaLPHoldings{IonutDRD, LQDUser011}
	LQDUser011 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser012 = VestaLPHoldings{Buhaici, LQDUser012}
	LQDUser012 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser013 = VestaLPHoldings{TheKid, LQDUser013}
	LQDUser013 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser014 = VestaLPHoldings{RaulTM, LQDUser014}
	LQDUser014 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
	//
	VLQUser015 = VestaLPHoldings{MakeAStep, LQDUser015}
	LQDUser015 = VestaLPs{p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0"), p.NFS("0")}
)

// Individual Multiplier Computation
func ComputeTotalIM(Input []VestaHoldings) *p.Decimal {
	var (
		TotalIM = p.NFS("0")
	)
	for i := 0; i < len(Input); i++ {
		CycleIM := GetSingleIM(Input[i])
		TotalIM = sm.ADDxc(TotalIM, CycleIM)
	}
	return TotalIM
}

func ComputeFinalTotalIM(Input []VestaHoldings) *p.Decimal {
	return sm.ADDxc(sm.DIVxc(ComputeTotalIM(Input), p.NFS("100")), p.NFS("1"))
}

func GetSingleIM(Input VestaHoldings) *p.Decimal {
	V1 := sm.MULxc(p.NFI(Input.Gold), p.NFS("5"))
	V2 := sm.MULxc(p.NFI(Input.Silver), p.NFS("2.5"))
	V3 := sm.MULxc(p.NFI(Input.Bronze), p.NFS("1.25"))
	return sm.SUMxc(V1, V2, V3)
}

// Computes Guest IM, and Percentual Split Chain given a Guest Position
func ComputeVestaSplit(Position int64, Input []VestaHoldings) (GuestIM *p.Decimal, PercentualSplitChain []*p.Decimal) {

	var (
		UnitPercent = new(p.Decimal)
	)
	PercentualSplitChain = make([]*p.Decimal, len(Input))

	GIM := GetSingleIM(Input[Position])
	GuestIM = sm.ADDxc(sm.DIVxc(GIM, p.NFS("100")), p.NFS("1"))
	TIM := ComputeTotalIM(Input)
	TIM2 := sm.SUBxc(TIM, GIM)

	for i := 0; i < len(Input); i++ {
		if int64(i) == Position {
			PercentualSplitChain[i] = p.NFS("1")
		} else {
			CurrentIM := GetSingleIM(Input[i])
			UnitPercent = sm.TruncateCustom(sm.DIVxc(CurrentIM, TIM2), 18)
			PercentualSplitChain[i] = UnitPercent
		}
	}
	return GuestIM, PercentualSplitChain
}

// Outputs Farm Multiplier Structure
func OutputVVMx(InputIM, InputUM *p.Decimal) (Output FarmMx) {
	Output = FarmMx{AncientDEB, Neutral, VestaTM, Neutral, InputUM, InputIM, Neutral, Neutral}
	return Output
}

// Computes Total Multiplier of a farm Structure
func GetAllMx(Input FarmMx) *p.Decimal {
	return sm.PRDxc(Input.DEB, Input.VLM, Input.TM, Input.D, Input.UM, Input.IM, Input.DM, Input.VM)
}

// Creates a split of Vesta Values, given an amount, an chain of percents.
func MakeVestaSplit(Amount *p.Decimal, VS []*p.Decimal) []*p.Decimal {
	var (
		OutputVestaChain = make([]*p.Decimal, len(VS))
		UnitVesta        = new(p.Decimal)
	)
	Half := sm.TruncateCustom(sm.DIVxc(Amount, p.NFS("2")), 18)
	for i := 0; i < len(VS); i++ {
		UnitVesta = sm.TruncateCustom(sm.MULxc(Half, VS[i]), 18)
		OutputVestaChain[i] = UnitVesta
	}
	return OutputVestaChain
}

// Given an InputUM, a Guest position, and a RawVesta Amount to be split, computes a Chain of Vesta Token Values
// According to the List of SFT Holders.
func AbsolutSplitWithVesta(RWAmount, InputUM *p.Decimal, Position int64, Input []VestaHoldings) (Remaining *p.Decimal, VestaAmountChain []*p.Decimal) {
	TotalIM := ComputeFinalTotalIM(Input)

	GIM, VS := ComputeVestaSplit(Position, Input)

	ALLVestaFarmMultipliers := OutputVVMx(TotalIM, InputUM)
	//fmt.Println("M1: ", ALLVestaFarmMultipliers)
	MyVestaFarmMultipliers := OutputVVMx(GIM, InputUM)
	//fmt.Println("M2: ", MyVestaFarmMultipliers)

	PersonalAmount := sm.TruncateCustom(sm.MULxc(RWAmount, GetAllMx(MyVestaFarmMultipliers)), 18)
	//fmt.Println("PersonalAmount is, ", PersonalAmount)
	BoostedAmount := sm.TruncateCustom(sm.MULxc(RWAmount, GetAllMx(ALLVestaFarmMultipliers)), 18)
	GainedAmount := sm.SUBxc(BoostedAmount, PersonalAmount)
	//fmt.Println("PA: ", PersonalAmount)
	//fmt.Println("BA: ", BoostedAmount)

	return PersonalAmount, MakeVestaSplit(GainedAmount, VS)
}

//Compute Raw Splits based on Individual Liquidity.

// Computes the VLP for an individual, used later for the VLP Split given individual user Liquidity
func ComputeUserVLP(Input VestaLPHoldings) *p.Decimal {
	var (
		Output = new(p.Decimal)
	)
	if Input.VestaLiq == LQDEmpty {
		Output = p.NFS("0")
	} else {
		V1 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.Gold, p.NFS("2.5")), 18)
		V2 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.Silver, p.NFS("1.6")), 18)
		V3 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.Bronze, p.NFS("1")), 18)
		V4 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.UGold, p.NFS("1.75")), 18)
		V5 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.USilver, p.NFS("0.96")), 18)
		V6 := sm.TruncateCustom(sm.MULxc(Input.VestaLiq.UBronze, p.NFS("0.5")), 18)
		Output = sm.SUMxc(V1, V2, V3, V4, V5, V6)
	}

	return Output
}

// Computes total VLP considering the Liquidity Holdings of all Participants
func ComputeTotalVLP(Input []VestaLPHoldings) *p.Decimal {
	var (
		VLPSum = p.NFS("0")
		Unit   = new(p.Decimal)
	)
	for i := 0; i < len(Input); i++ {
		Unit = ComputeUserVLP(Input[i])
		VLPSum = sm.ADDxc(VLPSum, Unit)
	}
	return VLPSum
}

// Computes the VLP Split given the Liquidity Pooled by all participants.
// Used to compute the RawVesta Amount each individual user would earn
// This is further used when computing individual user yield based on individual liquidity
func ComputeVLPSplit(Input []VestaLPHoldings) []*p.Decimal {
	var (
		VLPSplitChain = make([]*p.Decimal, len(Input))
		UnitVLPSplit  = new(p.Decimal)
	)
	GroupVLP := ComputeTotalVLP(Input)

	for i := 0; i < len(Input); i++ {
		VLP := ComputeUserVLP(Input[i])
		if sm.DecimalEqual(VLP, p.NFS("0")) == true {
			UnitVLPSplit = p.NFS("0")
		} else {
			UnitVLPSplit = sm.TruncateCustom(sm.DIVxc(VLP, GroupVLP), 18)
		}
		VLPSplitChain[i] = UnitVLPSplit
	}
	return VLPSplitChain

}

// Variadic Vesta Chain Adder
// Seems i didnt need it
func TotalVestaChainAdder(First []*p.Decimal, Rest ...[]*p.Decimal) []*p.Decimal {
	MakeZeroSlice := func(length int, Item *p.Decimal) []*p.Decimal {
		OutputSlice := make([]*p.Decimal, length)
		for i := range OutputSlice {
			OutputSlice[i] = Item
		}
		return OutputSlice
	}
	sum := MakeZeroSlice(len(First), p.NFS("0"))
	restsum := MakeZeroSlice(len(First), p.NFS("0"))

	for _, item := range Rest {
		restsum = VestaChainAdder(restsum, item)
	}
	sum = VestaChainAdder(First, restsum)
	return sum
}

// Function for adding two slices of decimals
func VestaChainAdder(C1, C2 []*p.Decimal) []*p.Decimal {
	var (
		OutputChain = make([]*p.Decimal, len(C1))
		Unit        = new(p.Decimal)
	)
	for i := 0; i < len(C1); i++ {
		Unit = sm.TruncateCustom(sm.ADDxc(C1[i], C2[i]), 18)
		OutputChain[i] = Unit
	}
	return OutputChain
}

// Computing Individual Minting Amounts
// Multiplies a decimal with each decimal in a chain of decimals
func CreateRawVestaSplit(RawAmount *p.Decimal, VLPSplit []*p.Decimal) []*p.Decimal {
	var (
		OutputChain = make([]*p.Decimal, len(VLPSplit))
		Unit        = new(p.Decimal)
	)
	for i := 0; i < len(VLPSplit); i++ {
		Unit = sm.TruncateCustom(sm.MULxc(RawAmount, VLPSplit[i]), 18)
		OutputChain[i] = Unit
	}
	return OutputChain
}

// The Final Function that computes individual Vesta yields considering all participants guests.
// Then adds the individual computed Vesta Yields Together
func MultipleAbsoluteSplitWithVesta(RawVestaAmount, InputUM *p.Decimal, VestaSFTsChain []VestaHoldings, LPChain []VestaLPHoldings) (*p.Decimal, []*p.Decimal) {
	VLPSplit := ComputeVLPSplit(LPChain)
	RawVestaSplit := CreateRawVestaSplit(RawVestaAmount, VLPSplit)

	MakeZeroSlice := func(length int, Item *p.Decimal) []*p.Decimal {
		OutputSlice := make([]*p.Decimal, length)
		for i := range OutputSlice {
			OutputSlice[i] = Item
		}
		return OutputSlice
	}
	SummedChain := MakeZeroSlice(len(VestaSFTsChain), p.NFS("0"))

	var (
		VestaRewardForPosition = make([]*p.Decimal, len(VestaSFTsChain))
		PersonalAmountsChain   = make([]*p.Decimal, len(VestaSFTsChain))
		PA                     = new(p.Decimal)
		AH                     = new(p.Decimal)
	)

	for i := 0; i < len(VestaSFTsChain); i++ {
		//fmt.Println("**************")
		PA, VestaRewardForPosition = AbsolutSplitWithVesta(RawVestaSplit[i], InputUM, int64(i), VestaSFTsChain)
		//fmt.Println("Chain on position ", i, " is ", VestaRewardForPosition)
		//fmt.Println("**************")
		SummedChain = VestaChainAdder(SummedChain, VestaRewardForPosition)
		if i == 0 {
			AH = PA
			PersonalAmountsChain[i] = p.NFS("0")
		} else {
			PersonalAmountsChain[i] = PA
		}
	}
	//Personal Amount Chain is the chain with personal Vesta Amounts owner of sent liquidity is earning
	//with his liquidity and with his SFTs. This amount is his in entirety, and he only splits the extra
	//Vesta Token amounts generated with his boosters.
	OutputChain := VestaChainAdder(SummedChain, PersonalAmountsChain)
	return AH, OutputChain
}

func SumChain(InputChain []*p.Decimal) *p.Decimal {
	var (
		SUM = p.NFS("0")
	)
	for i := 0; i < len(InputChain); i++ {
		SUM = sm.ADDxc(SUM, InputChain[i])
	}
	return SUM
}

func ComputeMintPercent(PersonalAmount *p.Decimal, InputChain []*p.Decimal) *p.Decimal {
	ChainSum := SumChain(InputChain)
	TotalVST := sm.ADDxc(ChainSum, PersonalAmount)
	fmt.Println("Total VST is: ", TotalVST)
	TotalPersonalAmount := sm.ADDxc(PersonalAmount, InputChain[0])
	fmt.Println("TotalPersonal Amount is ", TotalPersonalAmount)
	OutgoingAmount := sm.SUBxc(ChainSum, InputChain[0])
	fmt.Println("Outgoing Amount is: ", OutgoingAmount)

	OutgoingAmountPercent := sm.TruncateCustom(sm.DIVxc(OutgoingAmount, TotalVST), 18)
	fmt.Println("OutgoingPercent := ", OutgoingAmountPercent)
	OutgoingAmountRoundUP := sm.TruncateCustom(sm.MULxc(OutgoingAmountPercent, p.NFS("100")), 0)
	FinalOutgoingAmount := sm.ADDxc(OutgoingAmountRoundUP, p.NFS("1"))
	fmt.Println("OutgoingPercent final = ", FinalOutgoingAmount)
	fmt.Println("================================================")
	return OutgoingAmountRoundUP
}

func ExportOutgoingVestas(MainChain []VestaHoldings, Rewards []*p.Decimal) []mvx.BalanceESDT {
	var (
		OutputChain = make([]mvx.BalanceESDT, len(MainChain))
	)
	for i := 0; i < len(MainChain); i++ {
		OutputChain[i].Address = MainChain[i].Address
		OutputChain[i].Balance = sm.DTS(Rewards[i])
	}
	mvx.ConvertToBulkCSV("ExportVesta.csv", OutputChain[1:])
	return OutputChain[1:]
}
