package main

import (
	"Demiourgos/Blooming"
	mt "Demiourgos/SuperMath"
	vt "Demiourgos/Vesta"
	"fmt"
	p "github.com/Crypt0plasm/Firefly-APD"
	"os"
)

func WriteListOneByOneC(Name string, List []vt.VestaSplit) {
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

func ScanPools() (Output []vt.VestaPool) {
	//Scan Pool Number 1, Super Pool
	SuperPool := vt.ScanLiquidity(vt.EGLD_Super_LP, vt.SuperIdentifier)
	WeighedSuperPool := vt.WeightPool(SuperPool, p.NFS("10"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedSuperPool.VEGLD, p.NFS("0")) == true {
		Output = append(Output, WeighedSuperPool)
	}

	//Scan Pool Number 2, Crust
	CrustPool := vt.ScanLiquidity(vt.EGLD_Crust_LP, vt.CrustIdentifier)
	WeighedCrustPool := vt.WeightPool(CrustPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedCrustPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedCrustPool)
	}

	//Scan Pool Number 3 Aero
	AeroPool := vt.ScanLiquidity(vt.EGLD_Aero_LP, vt.AeroIdentifier)
	WeighedAeroPool := vt.WeightPool(AeroPool, p.NFS("1"))
	//Super Pool always generates Vesta, regardless of VEGLD amount it holds.
	if mt.DecimalGreaterThan(WeighedAeroPool.VEGLD, p.NFS("500")) == true {
		Output = append(Output, WeighedAeroPool)
	}

	return
}

func ComputePoolVestaSplit(TotalVesta *p.Decimal, Input []vt.VestaPool) []vt.VestaSplit {
	Output := make([]vt.VestaSplit, len(Input))
	Sum := p.NFS("0")
	var (
		Remainder    *p.Decimal
		Last         *p.Decimal
		VestaForPool *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Sum = mt.ADDxc(Sum, Input[i].VEGLD)
	}

	for i := 0; i < len(Input); i++ {

		Output[i].Pool = Input[i]
		VestaForPool = mt.TruncateCustom(mt.MULxc(mt.DIVxc(Input[i].VEGLD, Sum), TotalVesta), 18)
		if i == 0 {
			Remainder = mt.SUBxc(TotalVesta, VestaForPool)
		} else {
			Remainder = mt.SUBxc(Remainder, VestaForPool)
		}
		if i == len(Input)-2 {
			Last = Remainder
		}

		if i == len(Input)-1 {
			Output[i].Vesta = Last
		} else {
			Output[i].Vesta = VestaForPool
		}
	}

	return Output
}

func ComputeIndividualVestaSplit(TotalVesta *p.Decimal, Input []Blooming.BalanceESDT) []Blooming.BalanceESDT {
	Output := make([]Blooming.BalanceESDT, len(Input))
	Sum := p.NFS("0")
	var (
		Remainder          *p.Decimal
		Last               *p.Decimal
		VestaForIndividual *p.Decimal
	)

	for i := 0; i < len(Input); i++ {
		Sum = mt.ADDxc(Sum, p.NFS(Input[i].Balance))
	}

	for i := 0; i < len(Input); i++ {

		VestaForIndividual = mt.TruncateCustom(mt.MULxc(mt.DIVxc(p.NFS(Input[i].Balance), Sum), TotalVesta), 18)
		if i == 0 {
			Remainder = mt.SUBxc(TotalVesta, VestaForIndividual)
		} else {
			Remainder = mt.SUBxc(Remainder, VestaForIndividual)
		}
		if i == len(Input)-2 {
			Last = Remainder
		}

		if i == len(Input)-1 {
			Output[i].Balance = mt.DTS(Last)
		} else {
			Output[i].Balance = mt.DTS(VestaForIndividual)
		}

		Output[i].Address = Input[i].Address

	}

	return Output
}

func ScanLPOwners(LPToken vt.ESDT, ExportName string) (Output []Blooming.BalanceESDT) {
	LPScan := Blooming.CreateBalanceChain(vt.MakeESDTSnapshotLink(LPToken))
	SortedLPScan := Blooming.SortBalanceSFTChain(LPScan)
	SortedLPScanESDT := Blooming.ConvertSFTtoESDTChain(SortedLPScan)
	Sum := Blooming.AtomicUnitsDecimalToDecimalString(Blooming.AddBalanceSFTChain(SortedLPScan))

	fmt.Println("LP-Token ", string(LPToken), " sum is ", Sum, " on ", len(SortedLPScan), " addresses.")
	Blooming.WriteListOneByOneB(ExportName, SortedLPScanESDT)
	return SortedLPScanESDT
}

func main() {
	//Scan Liquidity Pools
	LiquidityPools := ScanPools()
	PoolVestaSplit := ComputePoolVestaSplit(p.NFS("800"), LiquidityPools)
	for i := 0; i < len(PoolVestaSplit); i++ {
		fmt.Println("Pool ", i, "has ", LiquidityPools[i].VEGLD, "vEGLD equivalent")
		fmt.Println("Pool ", i, "gets ", PoolVestaSplit[i].Vesta, "Vesta")
	}
	WriteListOneByOneC("V_VestaSplit.txt", PoolVestaSplit)

	//Scan LP Owners
	LP00 := ScanLPOwners(vt.SUPEREGLD, "V_Pool_00_LPs.txt")
	LP01 := ScanLPOwners(vt.CRUSTEGLD, "V_Pool_01_LPs.txt")
	LP02 := ScanLPOwners(vt.AEROEGLD, "V_Pool_02_LPs.txt")

	//Compute Individual Vesta Split
	LP00VS := ComputeIndividualVestaSplit(PoolVestaSplit[0].Vesta, LP00)
	Blooming.WriteListOneByOneB("V_LP00_VS.txt", LP00VS)

	LP01VS := ComputeIndividualVestaSplit(PoolVestaSplit[1].Vesta, LP01)
	Blooming.WriteListOneByOneB("V_LP01_VS.txt", LP01VS)

	LP02VS := ComputeIndividualVestaSplit(PoolVestaSplit[2].Vesta, LP02)
	Blooming.WriteListOneByOneB("V_LP02_VS.txt", LP02VS)

}
