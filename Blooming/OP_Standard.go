package Blooming

import (
	mt "Demiourgos/SuperMath"
	p "github.com/Crypt0plasm/Firefly-APD"
	"io/ioutil"
	"log"
	"net/http"
)

// ======================================================================================================================
//
//			Demiourgos/Blooming/OP_Standard.go
//			Basic Functions, Blockchain related Functions and math.
//
//
//	             [A]Basic Functions
//
// [A]01         PercentSwing            Computes the % difference between 2 values.
// [A]02         OnPage                  Basic Snapshot Function
//
//	[B]AtomicUnit String Converter Functions
//
// [B]01         ConvertAU18             Converts a string of numbers as AU to decimals, considering 18 decimals.
// [B]02         ConvertAU06             Converts a string of numbers as AU to decimals, considering  6 decimals.
//
// [C]01         MinDecimal              Gets the minimum between two decimals
// ======================================================================================================================
// ======================================================================================================================
//
// [A]           Basic Functions
//
// [A]01         PercentSwing
//
//	Computes the % difference between Value2 and Value1.
func PercentSwing(Value1, Value2 *p.Decimal) *p.Decimal {
	M1 := mt.MULxc(Value2, p.NFS("100"))
	D1 := mt.DIVxc(M1, Value1)
	S1 := mt.SUBxc(D1, p.NFS("100"))
	PP := mt.TruncateCustom(S1, 6)
	return PP
}

// ======================================================================================================================
//
// [A]02         OnPage
//
//	Basic Snapshot Function
//	Snapshots Link and returns string
func OnPage(Link string) string {
	res, err := http.Get(Link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// ======================================================================================================================
//
// [B]           AtomicUnit String Converter Functions
//
// [B]01         ConvertAU18 Converts a string of raw numbers as atomic units, to its respective Decimal
//
//	Usage for 18 Decimals
func ConvertAU18(Number string) *p.Decimal {
	Value := p.NFS(Number)
	Result := mt.DIVxc(Value, mt.POWxc(p.NFI(10), p.NFI(18)))
	return Result
}

// ======================================================================================================================
//
// [B]02         ConvertAU06 Converts a string of raw numbers as atomic units, to its respective Decimal
//
//	Usage for 6 Decimals
func ConvertAU06(Number string) *p.Decimal {
	Value := p.NFS(Number)
	Result := mt.DIVxc(Value, mt.POWxc(p.NFI(10), p.NFI(6)))
	return Result
}

// ======================================================================================================================
//
// [C]01         MinDecimal              Gets the minimum between two decimals
func MinDecimal(No1, No2 *p.Decimal) *p.Decimal {
	var Result *p.Decimal

	if mt.DecimalEqual(No1, No2) == true {
		Result = No1
	} else if mt.DecimalLessThan(No1, No2) == true {
		Result = No1
	} else {
		Result = No2
	}

	return Result
}

//
//
