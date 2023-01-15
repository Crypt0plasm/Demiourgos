package Blooming

import (
	mt "Demiourgos/SuperMath"
	"encoding/json"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"strconv"
)

// ======================================================================================================================
//
//	Blooming/CDChainMaker.go
//	Creates Chains with Coding Division SFT Values.
//
// [A]00         CreateBalanceChain			Creates a Chain of values
//
// [B]01         CreateCodingDivisionChain       	Create a Chain of Character Chain Values
// [B]02         CreateCodingDivisionOwners      	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
// [B]03         GetAddressAmount                	Gets the Amount for a given address in a Balance Chain
// [B]04         CreateCodingDivisionSetChain    	Create a Chain of all Elrond Addresses Containing Coding Division SFTs Sets
// [B]05         CreateCodingDivisionAmountChain 	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
// [B]06         SortBalanceSFTChain             	Sorts a BalanceSFTChain Chain from highest Balance to lowest Balance
// [B]07         AddBalanceSFTChain              	Add the Balances of an SFT Chain
// ======================================================================================================================
// ======================================================================================================================
//
// [A]00         CreateBalanceChain
//
//		Creates a balance Chain (Address and Balance)
//	     	Link must be of type that outputs and Address and a Balance
func CreateBalanceChain(Link string) []BalanceSFT {
	var OutputChain []BalanceSFT
	SS := OnPage(Link)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

// ======================================================================================================================
//
// [B]01         CreateCodingDivisionChain
//
//	Create a Chain of all Character Chain Values
func CreateCodingDivisionChain() (DecaChain, int, *p.Decimal) {
	var OutputChain DecaChain

	fmt.Println("Snapshotting SnakeEye Addresses and Amounts")
	SnakeEyeChain := CreateBalanceChain(CD01SnakeEye)
	//SortSnake := SortBalanceSFTChain(SnakeEyeChain)
	WriteListOneByOne("Snake.txt", SnakeEyeChain)
	S01 := AddBalanceSFTChain(SnakeEyeChain)
	fmt.Println(len(SnakeEyeChain), "addresses snapshotted with Snake Eye SFTs", S01)
	fmt.Println("")

	fmt.Println("Snapshotting Rudis Addresses and Amounts")
	RudisChain := CreateBalanceChain(CD02Rudis)
	S02 := AddBalanceSFTChain(RudisChain)
	fmt.Println(len(RudisChain), "addresses snapshotted with Rudis SFTs", S02)
	fmt.Println("")

	fmt.Println("Snapshotting Gwen Addresses and Amounts")
	GwenChain := CreateBalanceChain(CD03Gwen)
	S03 := AddBalanceSFTChain(GwenChain)
	fmt.Println(len(GwenChain), "addresses snapshotted with Gwen SFTs", S03)
	fmt.Println("")

	fmt.Println("Snapshotting Clutter Addresses and Amounts")
	ClutterChain := CreateBalanceChain(CD04Clutter)
	S04 := AddBalanceSFTChain(ClutterChain)
	fmt.Println(len(ClutterChain), "addresses snapshotted with Clutter SFTs", S04)
	fmt.Println("")

	fmt.Println("Snapshotting Bangai Addresses and Amounts")
	BangaiChain := CreateBalanceChain(CD05Bangai)
	S05 := AddBalanceSFTChain(BangaiChain)
	fmt.Println(len(BangaiChain), "addresses snapshotted with Bangai SFTs", S05)
	fmt.Println("")

	fmt.Println("Snapshotting Binos Addresses and Amounts")
	BinosChain := CreateBalanceChain(CD06Binos)
	S06 := AddBalanceSFTChain(BinosChain)
	fmt.Println(len(BinosChain), "addresses snapshotted with Binos SFTs", S06)
	fmt.Println("")

	fmt.Println("Snapshotting Rubia Addresses and Amounts")
	RubiaChain := CreateBalanceChain(CD07Rubia)
	S07 := AddBalanceSFTChain(RubiaChain)
	fmt.Println(len(RubiaChain), "addresses snapshotted with Rubia SFTs", S07)
	fmt.Println("")

	fmt.Println("Snapshotting Ocultus Addresses and Amounts")
	OcultusChain := CreateBalanceChain(CD08Ocultus)
	S08 := AddBalanceSFTChain(OcultusChain)
	fmt.Println(len(OcultusChain), "addresses snapshotted with Ocultus SFTs", S08)
	fmt.Println("")

	fmt.Println("Snapshotting Oreta Addresses and Amounts")
	OretaChain := CreateBalanceChain(CD09Oreta)
	S09 := AddBalanceSFTChain(OretaChain)
	fmt.Println(len(OretaChain), "addresses snapshotted with Oreta SFTs", S09)
	fmt.Println("")

	fmt.Println("Snapshotting Binar Addresses and Amounts")
	BinarChain := CreateBalanceChain(CD10Binar)
	S10 := AddBalanceSFTChain(BinarChain)
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

func RemoveDuplicateElrondAddresses(Input []ElrondAddress) []ElrondAddress {
	//3.    Remove Duplicate Elrond Addresses.
	//3.1   Make a hash map from ElrondAddress to int
	Check := make(map[ElrondAddress]int)

	//3.2   Make the empty Output Slice that will contain unique Elrond Addresses
	Unique := make([]ElrondAddress, 0)

	//3.3   Iterate through the Slice containing duplicates and map each element to 0. (or any other thing)
	for _, val := range Input {
		Check[val] = 1
	}

	//3.4   Now finally iterate through the map and append each key of the map to a
	//      new slice of strings. Since any duplicate value too will be mapped to the
	//      same number as the previous one, hence all the keys will be unique.
	for letter, _ := range Check {
		Unique = append(Unique, letter)
	}
	return Unique
}

// ======================================================================================================================
//
// [B]02         CreateCodingDivisionOwners
//
//	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
func CreateCodingDivisionOwners(InputChain DecaChain) []ElrondAddress {

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

	//2.    Make a slice with all Elrond Address (will contain duplicate Elrond Addresses.
	//      basically removes the balance value.
	ElrondSlice := make([]ElrondAddress, len(AllSlice))

	for i := 0; i < len(AllSlice); i++ {
		ElrondSlice[i] = AllSlice[i].Address
	}

	return RemoveDuplicateElrondAddresses(ElrondSlice)
}

// ======================================================================================================================
//
// [B]03         GetAddressAmount
//
//	Gets the Amount for a given address in a Balance Chain
func GetAddressAmount(Addy ElrondAddress, Chain []BalanceSFT) string {
	var Result string
	for i := 0; i < len(Chain); i++ {
		if Chain[i].Address == Addy {
			Result = Chain[i].Balance
			break
		} else {
			Result = "0"
		}
	}
	return Result
}

// ======================================================================================================================
//
// [B]04         CreateCodingDivisionSetChain
//
//	Create a Chain of all Elrond Addresses Containing Coding Division SFTs Sets
func CreateCodingDivisionSetChain(Owners []ElrondAddress, Snapshot DecaChain) []BalanceSFT {
	var (
		SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar *p.Decimal
		Minimum                                                                     *p.Decimal
		SetChain                                                                    []BalanceSFT
		Unit                                                                        BalanceSFT
	)

	for i := 0; i < len(Owners); i++ {
		SnakeEye = p.NFS(GetAddressAmount(Owners[i], Snapshot.SnakeEye))
		Rudis = p.NFS(GetAddressAmount(Owners[i], Snapshot.Rudis))
		Gwen = p.NFS(GetAddressAmount(Owners[i], Snapshot.Gwen))
		Clutter = p.NFS(GetAddressAmount(Owners[i], Snapshot.Clutter))
		Bangai = p.NFS(GetAddressAmount(Owners[i], Snapshot.Bangai))
		Binos = p.NFS(GetAddressAmount(Owners[i], Snapshot.Binos))
		Rubia = p.NFS(GetAddressAmount(Owners[i], Snapshot.Rubia))
		Ocultus = p.NFS(GetAddressAmount(Owners[i], Snapshot.Ocultus))
		Oreta = p.NFS(GetAddressAmount(Owners[i], Snapshot.Oreta))
		Binar = p.NFS(GetAddressAmount(Owners[i], Snapshot.Binar))

		Minimum = MinDecimal(SnakeEye, Rudis)
		Minimum = MinDecimal(Minimum, Gwen)
		Minimum = MinDecimal(Minimum, Clutter)
		Minimum = MinDecimal(Minimum, Bangai)
		Minimum = MinDecimal(Minimum, Binos)
		Minimum = MinDecimal(Minimum, Rubia)
		Minimum = MinDecimal(Minimum, Ocultus)
		Minimum = MinDecimal(Minimum, Oreta)
		Minimum = MinDecimal(Minimum, Binar)

		//If Minimum is greater than zero AND Address is non Exception
		if mt.DecimalGreaterThan(Minimum, p.NFS("0")) == true && ComputeExceptionAddress(Owners[i]) == false {
			Unit.Address = Owners[i]
			Unit.Balance = strconv.Itoa(int(p.INT64(Minimum)))

			SetChain = append(SetChain, Unit)
		}
	}
	return SetChain
}

func CreateCodingDivisionSetExceptionChain(Input []BalanceSFT, UseException bool) []BalanceSFT {
	var (
		Output    []BalanceSFT
		Unit      BalanceSFT
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

// ======================================================================================================================
//
// [B]05         CreateCodingDivisionAmountChain
//
//	Create a Chain of all Elrond Addresses Containing Coding Division SFTs
func CreateCodingDivisionAmountChain(Owners []ElrondAddress, Snapshot DecaChain) []BalanceSFT {
	var (
		SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar *p.Decimal
		Sum                                                                         *p.Decimal
		SetChain                                                                    []BalanceSFT
		Unit                                                                        BalanceSFT
	)

	for i := 0; i < len(Owners); i++ {
		SnakeEye = p.NFS(GetAddressAmount(Owners[i], Snapshot.SnakeEye))
		Rudis = p.NFS(GetAddressAmount(Owners[i], Snapshot.Rudis))
		Gwen = p.NFS(GetAddressAmount(Owners[i], Snapshot.Gwen))
		Clutter = p.NFS(GetAddressAmount(Owners[i], Snapshot.Clutter))
		Bangai = p.NFS(GetAddressAmount(Owners[i], Snapshot.Bangai))
		Binos = p.NFS(GetAddressAmount(Owners[i], Snapshot.Binos))
		Rubia = p.NFS(GetAddressAmount(Owners[i], Snapshot.Rubia))
		Ocultus = p.NFS(GetAddressAmount(Owners[i], Snapshot.Ocultus))
		Oreta = p.NFS(GetAddressAmount(Owners[i], Snapshot.Oreta))
		Binar = p.NFS(GetAddressAmount(Owners[i], Snapshot.Binar))

		Sum = mt.SUMxc(SnakeEye, Rudis, Gwen, Clutter, Bangai, Binos, Rubia, Ocultus, Oreta, Binar)

		//If Sum is greater than zero AND Address is non Exception
		if mt.DecimalGreaterThan(Sum, p.NFS("0")) == true && ComputeExceptionAddress(Owners[i]) == false {
			Unit.Address = Owners[i]
			Unit.Balance = strconv.Itoa(int(p.INT64(Sum)))

			SetChain = append(SetChain, Unit)
		}
	}
	return SetChain
}

func CreateCodingDivisionAmountExceptionChain(Input []BalanceSFT, UseException bool) []BalanceSFT {
	var (
		Unit           BalanceSFT
		Output         []BalanceSFT
		Exception      bool
		ExceptionValue string
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
				Unit.Balance = mt.DTS(mt.SUBxc(p.NFS(Input[i].Balance), p.NFS(ExceptionValue)))
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

//
//

// ======================================================================================================================
//
// [B]05         SortBalanceSFTChain
//
//	Sorts a BalanceSFTChain Chain from highest Balance to lowest Balance
func SortBalanceSFTChain(Chain []BalanceSFT) []BalanceSFT {
	var (
		SortedChain []BalanceSFT
	)
	GetMaxElement := func(Chain []BalanceSFT) int {
		Max := 0
		for i := 0; i < len(Chain); i++ {
			if mt.DecimalGreaterThanOrEqual(p.NFS(Chain[i].Balance), p.NFS(Chain[Max].Balance)) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := BalanceSFT{Address: Chain2Sort[Biggest].Address, Balance: Chain2Sort[Biggest].Balance}
		SortedChain = append(SortedChain, Unit)

		//Removing biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)
	}
	return SortedChain
}
func SortBalanceESDTChain(Chain []BalanceESDT) []BalanceESDT {
	var (
		SortedChain []BalanceESDT
	)
	GetMaxElement := func(Chain []BalanceESDT) int {
		Max := 0
		for i := 0; i < len(Chain); i++ {
			if mt.DecimalGreaterThanOrEqual(p.NFS(Chain[i].Balance), p.NFS(Chain[Max].Balance)) == true {
				Max = i
			}
		}
		return Max
	}
	Chain2Sort := Chain

	for i := 0; i < len(Chain); i++ {
		Biggest := GetMaxElement(Chain2Sort)
		Unit := BalanceESDT{Address: Chain2Sort[Biggest].Address, Balance: Chain2Sort[Biggest].Balance}
		SortedChain = append(SortedChain, Unit)

		//Removing biggest element
		//This syntax removes from a slice the element on position Biggest
		Chain2Sort = append(Chain2Sort[:Biggest], Chain2Sort[Biggest+1:]...)
	}
	return SortedChain
}

// ======================================================================================================================
//
// [B]06         AddBalanceSFTChain
//
//	Add the Balances of an SFT Chain
func AddBalanceSFTChain(Chain []BalanceSFT) *p.Decimal {
	Sum := p.NFS("0")
	for i := 0; i < len(Chain); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Chain[i].Balance))
	}
	return Sum
}

func ComputeRewards(Input []BalanceSFT, Reward string) []BalanceSFT {
	var (
		Unit   BalanceSFT
		Output []BalanceSFT
	)
	for i := 0; i < len(Input); i++ {
		Unit.Address = Input[i].Address
		Unit.Balance = mt.DTS(mt.MULxc(p.NFS(Input[i].Balance), p.NFS(Reward)))
		Output = append(Output, Unit)
	}
	return Output
}
