package Vesta

import (
	"Demiourgos/Blooming"
	mt "Demiourgos/SuperMath"
	"encoding/json"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
)

type VestaSplit struct {
	Pool  VestaPool
	Vesta *p.Decimal
}

type VestaPool struct {
	VEGLD *p.Decimal
	Token *p.Decimal
}

type ESDT string

var (
	VestaGold = "https://api.multiversx.com/nfts/VESTAXDAO-e6c48c-01/accounts?size=10000"

	//TokenIdentifiers
	WEGLDIdentifier = ESDT("WEGLD-bd4d79")
	VEGLDIdentifier = ESDT("VEGLD-2b9319")

	SuperIdentifier = ESDT("SUPER-507aa6")
	XLHIdentifier   = ESDT("XLH-8daa50")
	CrustIdentifier = ESDT("CRU-a5f4aa")
	AeroIdentifier  = ESDT("AERO-458bbf")

	//LP Tokens
	SUPEREGLD = ESDT("SUPEREGLD-a793b9")
	CRUSTEGLD = ESDT("CRUWEGLD-76c269")
	AEROEGLD  = ESDT("AEROWEGLD-81cc37")

	//Liquidity Pools
	EGLD_Super_LP = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqdx6z3sauy49c5k6c6lwhjqclrfwlxlud2jpsvwj5dp")
	EGLD_Crust_LP = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqj6daemefdk5kjgy9rs4zsng03kezgxdm2jps3h5n07")
	EGLD_Aero_LP  = Blooming.ElrondAddress("erd1qqqqqqqqqqqqqpgqzjctu8xrgn8jmfp503tajjvzz2zq60v92jpsslkh5a")
)

func MakeESDTSnapshotLink(Token ESDT) string {
	String1 := "https://api.multiversx.com/tokens/"
	String2 := "/accounts?size=10000"
	return String1 + string(Token) + String2
}

func VestaGoldSnapshot() []Blooming.BalanceSFT {
	var OutputChain []Blooming.BalanceSFT
	SS := Blooming.OnPage(VestaGold)
	_ = json.Unmarshal([]byte(SS), &OutputChain)
	return OutputChain
}

func CreateVestaGoldChain() []Blooming.BalanceSFT {
	fmt.Println("Snapshotting VestaGold Addresses and Amounts")
	VestaGoldChain := VestaGoldSnapshot()
	S01 := Blooming.AddBalanceSFTChain(VestaGoldChain)
	fmt.Println(len(VestaGoldChain), "addresses snapshotted with Vesta GOLD SFTs", S01)
	fmt.Println("")
	return VestaGoldChain
}

// Remove Exceptions
func CreateVestaGoldAmounts(Input []Blooming.BalanceSFT) []Blooming.BalanceSFT {
	var (
		Result []Blooming.BalanceSFT
		Unit   Blooming.BalanceSFT
	)

	for i := 0; i < len(Input); i++ {
		if Blooming.ComputeExceptionAddress(Input[i].Address) == false {
			Unit.Address = Input[i].Address
			Unit.Balance = Input[i].Balance
			Result = append(Result, Unit)
		}
	}
	return Result
}

// Functions that distribute Vesta
func ReadESDTAmount(Addy Blooming.ElrondAddress, Token ESDT) string {
	var (
		String1 = "https://api.multiversx.com/accounts/"
		String2 = "/tokens/"

		ScannedJSON ESDTSuperStructure
		Balance     string
	)
	ScanURL := String1 + string(Addy) + String2 + string(Token)

	Snapshot := Blooming.OnPage(ScanURL)
	_ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)
	if Snapshot == "[]" {
		Balance = "0"
	} else {
		Balance = ScannedJSON.Balance
	}

	Result := Blooming.AtomicUnitsStringToDecimalString(Balance)

	return Result
}

func ScanLiquidity(PoolSC Blooming.ElrondAddress, Token ESDT) VestaPool {
	var Result VestaPool
	Result.VEGLD = p.NFS(ReadESDTAmount(PoolSC, WEGLDIdentifier))
	Result.Token = p.NFS(ReadESDTAmount(PoolSC, Token))
	return Result
}

func WeightPool(Pool VestaPool, Weight *p.Decimal) (Output VestaPool) {
	Output.VEGLD = mt.MULxc(Pool.VEGLD, Weight)
	Output.Token = Pool.Token
	return
}
