package Rewards

import (
	p "Firefly-APD"
	mvx "MvxApiScanner"
	sm "SuperMath"
	"fmt"
	"os"
)

type DistributionEvidence struct {
	Type   string
	Mode   string
	Payees string
	Token  mvx.ESDT
	Amount *p.Decimal
	Split  string
}

var (
	// Type
	DistributionType1 = "Total"
	DistributionType2 = "Multiplication"

	// Mode
	DistributionMode1 = Payee3 + "Rewards"
	DistributionMode2 = "Snakes Only"
	DistributionMode3 = Payee3 + "Only"

	// Payees
	Payee1 = "Demiourgos.Holdings™"
	Payee2 = "Snakes"
	Payee3 = "Coding.Division™"
)

func MakeSingleDistributionSplit(Unit string, Amount *p.Decimal) string {
	Output := Unit + " * " + sm.DTS(Amount)
	return Output
}

func MakeDoublePayeesString(Payee1, Payee2 string, Amount1, Amount2 *p.Decimal) string {
	L1 := Payee1 + " (" + sm.DTS(Amount1) + ")"
	L2 := "                    : " + Payee2 + " (" + sm.DTS(Amount2) + ")"
	LT := L1 + "\n" + L2
	return LT
}

func MakeTripleCDSplit(Unit string, Amount *p.Decimal) string {
	UnitDistribution := sm.MULxc(Amount, p.NFS("0.5"))
	SnakesDistribution := sm.MULxc(Amount, p.NFS("0.15"))
	DemiourgosDistribution := sm.MULxc(Amount, p.NFS("0.35"))

	L1 := Unit + " = " + sm.DTS(UnitDistribution)
	L2 := "                    : " + Payee2 + " = " + sm.DTS(SnakesDistribution)
	L3 := "                    : " + Payee1 + " = " + sm.DTS(DemiourgosDistribution)
	LT := L1 + "\n" + L2 + "\n" + L3
	return LT
}
func MakeTotalCDEvidence(Type, Mode, Payeee1, Payeee2 string, Payeee1Q, Payeee2Q, DistributedAmount *p.Decimal, Token mvx.ESDT) DistributionEvidence {
	var Output DistributionEvidence

	Output.Type = Type
	Output.Mode = Mode
	Output.Payees = MakeDoublePayeesString(Payeee1, Payeee2, Payeee1Q, Payeee2Q)
	Output.Token = Token
	Output.Amount = DistributedAmount
	Output.Split = MakeTripleCDSplit(Payee3, DistributedAmount)
	return Output
}

func DistributionEvidenceMLS(Input DistributionEvidence) string {
	L1 := "Distribution Type   : " + Input.Type
	L2 := "Distribution Mode   : " + Input.Mode
	L3 := "Payees              : " + Input.Payees
	L4 := "Distributed Tokens  : " + string(Input.Token)
	L5 := "Distributed Amount  : " + sm.DTS(Input.Amount)
	L6 := "Distribution Split  : " + Input.Split
	LT := L1 + "\n" + L2 + "\n" + L3 + "\n" + L4 + "\n" + L5 + "\n" + L6
	return LT
}

func ExportEvidenceCD(ExportName string, SnakeChain []mvx.BalanceSFT, CDChain []mvx.BalanceESDT, Evidence DistributionEvidence) {
	f, err := os.Create(ExportName)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}
	//PrintDistribution Info
	S00 := "===========Summary=================================================="
	S0 := DistributionEvidenceMLS(Evidence)
	_, _ = fmt.Fprintln(f, S00)
	_, _ = fmt.Fprintln(f, S0)

	//Print Snake Chain
	S1 := "===========SNAKEs-Snapshots=========================================="
	_, _ = fmt.Fprintln(f, S1)
	for _, v := range SnakeChain {
		_, _ = fmt.Fprintln(f, v)
	}

	//Print CD Chain
	S2 := "===========Coding.Division™-Snapshots================================"
	_, _ = fmt.Fprintln(f, S2)
	for _, w := range CDChain {
		_, _ = fmt.Fprintln(f, w)
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Evidence file written successfully")
}
