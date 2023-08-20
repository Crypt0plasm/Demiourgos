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
	DistributionType3 = "Totalisation"

	// Mode
	DistributionMode1 = Payee3 + " Rewards"
	DistributionMode2 = "Snakes Only"
	DistributionMode3 = Payee3 + " Only"
	DistributionMode4 = Payee3 + " Raw Split"
	DistributionMode5 = Payee6 + " Raw Split"
	DistributionMode6 = Payee7 + " Raw Split"

	// Payees
	Payee1 = "Demiourgos.Holdings™"
	Payee2 = "Snakes"
	Payee3 = "Coding.Division™"
	Payee4 = "CD.DAO + Sn.DAO + DH™"
	Payee5 = "VS.DAO + VS.Vault + Sn.DAO + DH™"
	Payee6 = "VestaX.Finance™"
	Payee7 = "vEGLD Liquid.Staking™"
)

func MakeSinglePayeesString(PayeeeX string, Amount *p.Decimal) string {
	L1 := PayeeeX + " (" + sm.DTS(Amount) + ")"
	return L1
}

func MakeDoublePayeesString(PayeeX, PayeeY string, Amount1, Amount2 *p.Decimal) string {
	L1 := PayeeX + " (" + sm.DTS(Amount1) + ")"
	L2 := "                    : " + PayeeY + " (" + sm.DTS(Amount2) + ")"
	LT := L1 + "\n" + L2
	return LT
}

func MakeSingleSplit(Unit string, Amount *p.Decimal) string {
	L1 := Unit + " * " + sm.DTS(Amount)
	return L1
}

func MakeSingleTotalisationSplit(Unit string, Amount *p.Decimal) string {
	L1 := sm.DTS(Amount) + " / " + Unit
	return L1
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

func MakeQuadCDSplit(Unit string, Amount *p.Decimal, Split []*p.Decimal) string {
	L1 := Unit + " = " + sm.DTS(Split[0])
	L2 := "                    : " + Payee2 + " = " + sm.DTS(Split[1])
	L3 := "                    : " + "CEO" + " = " + sm.DTS(Split[2])
	L4 := "                    : " + "HoV" + " = " + sm.DTS(Split[2])
	L5 := "                    : " + "CTO" + " = " + sm.DTS(Split[3])
	LT := L1 + "\n" + L2 + "\n" + L3 + "\n" + L4 + "\n" + L5
	return LT
}

func MakeMultiplicationEvidence(Type, Mode, PayeeeX string, Payeee1Q, DistributedAmount *p.Decimal, Token mvx.ESDT) DistributionEvidence {
	var Output DistributionEvidence

	Output.Type = Type
	Output.Mode = Mode
	Output.Payees = MakeSinglePayeesString(PayeeeX, Payeee1Q)
	Output.Token = Token
	Output.Amount = DistributedAmount
	Output.Split = MakeSingleSplit(PayeeeX, DistributedAmount)
	return Output
}

func MakeTotalisationEvidence(Type, Mode, PayeeeX string, Payeee1Q, DistributedAmount *p.Decimal, Token mvx.ESDT) DistributionEvidence {
	var Output DistributionEvidence

	Output.Type = Type
	Output.Mode = Mode
	Output.Payees = MakeSinglePayeesString(PayeeeX, Payeee1Q)
	Output.Token = Token
	Output.Amount = DistributedAmount
	Output.Split = MakeSingleTotalisationSplit(PayeeeX, DistributedAmount)
	return Output
}

func MakeTotalCDEvidence(Type, Mode, PayeeeX, PayeeeY string, Payeee1Q, Payeee2Q, DistributedAmount *p.Decimal, Token mvx.ESDT, CDSplit []*p.Decimal) DistributionEvidence {
	var Output DistributionEvidence

	Output.Type = Type
	Output.Mode = Mode
	Output.Payees = MakeDoublePayeesString(PayeeeX, PayeeeY, Payeee1Q, Payeee2Q)
	Output.Token = Token
	Output.Amount = DistributedAmount
	Output.Split = MakeQuadCDSplit(Payee3, DistributedAmount, CDSplit)
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

func ExportEvidenceDoubleVesta(ExportName string, InputChain1, InputChain2 []mvx.BalanceESDT, Evidence DistributionEvidence) {
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

	//Print VestaGold Chain
	S1 := "===========GoldenVESTA-Snapshots====================================="
	_, _ = fmt.Fprintln(f, S1)
	for _, v := range InputChain1 {
		_, _ = fmt.Fprintln(f, v)
	}

	//Print VestaSilver Chain
	S2 := "===========SilverVESTA-Snapshots====================================="
	_, _ = fmt.Fprintln(f, S2)
	for _, v := range InputChain2 {
		_, _ = fmt.Fprintln(f, v)
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Evidence file written successfully")

}

func ExportEvidenceTripleVesta(ExportName string, InputChain1, InputChain2, InputChain3 []mvx.BalanceESDT, Evidence DistributionEvidence) {
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

	//Print VestaGold Chain
	S1 := "===========GoldenVESTA-Snapshots====================================="
	_, _ = fmt.Fprintln(f, S1)
	for _, v := range InputChain1 {
		_, _ = fmt.Fprintln(f, v)
	}

	//Print VestaSilver Chain
	S2 := "===========SilverVESTA-Snapshots====================================="
	_, _ = fmt.Fprintln(f, S2)
	for _, v := range InputChain2 {
		_, _ = fmt.Fprintln(f, v)
	}

	//Print BronzeSilver Chain
	S3 := "===========SilverVESTA-Snapshots====================================="
	_, _ = fmt.Fprintln(f, S3)
	for _, v := range InputChain3 {
		_, _ = fmt.Fprintln(f, v)
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Evidence file written successfully")

}

func ExportEvidenceMultiplication(ExportName string, InputChain []mvx.BalanceESDT, Evidence DistributionEvidence) {
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
	S1 := "===========Chain-Snapshots=========================================="
	_, _ = fmt.Fprintln(f, S1)
	for _, v := range InputChain {
		_, _ = fmt.Fprintln(f, v)
	}

	err = f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Evidence file written successfully")

}

func ExportEvidenceMultiplicationSFT(ExportName string, InputChain []mvx.BalanceESDT, Evidence DistributionEvidence) {
	ExportEvidenceMultiplication(ExportName, InputChain, Evidence)
	return
}

func ExportEvidenceCD(ExportName string, SnakeChain []mvx.BalanceESDT, CDChain []mvx.BalanceESDT, Evidence DistributionEvidence) {
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
