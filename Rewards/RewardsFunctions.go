package Rewards

import (
	bloom "Demiourgos/Blooming"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	sm "SuperMath"
	"fmt"
	"time"
)

var (
	RewardPath = "_REWRD-Snapshots\\"
	//OMNsctPath = "_SNAKE-Snapshots\\DEMIOU-6d1b5c-top-holders-omniscient-tools.csv"
	SnakePath = "_SNAKE-Snapshots\\"
	SnakeName = "DEMIOU-6d1b5c-top-holders-omniscient-tools.csv"
)

func MakeSnakeChain() ([]mvx.BalanceESDT, *p.Decimal) {
	//Make Snakes Chain
	Snakes, _ := bloom.CreateSnakeChain(mvx.SnakeSFT)
	SnakeTrim := bloom.CreateSnakeAmountChain(Snakes)
	SnakeTrimSorted := mvx.SortBalanceIntegerChain(SnakeTrim)
	SnakeTrimSortedSum := mvx.AddBalanceDecimalChain(SnakeTrimSorted)
	return SnakeTrimSorted, SnakeTrimSortedSum
}

func MakeCodingDivisionChain() ([]mvx.BalanceESDT, *p.Decimal) {
	//Make Coding Division Chain
	//Uses Amount Exception (Paul holds 50 Company SFTs that aren't Include in the computation)
	MultiChain, _, _ := bloom.CreateCodingDivisionChain()
	Owners := bloom.CreateCodingDivisionOwners(MultiChain)

	All := bloom.CreateCodingDivisionAmountChain(Owners, MultiChain)
	AllException := bloom.CreateCodingDivisionAmountExceptionChain(All, true)
	AllExceptionSorted := mvx.SortBalanceIntegerChain(AllException)
	AllExceptionSortedSum := mvx.AddBalanceDecimalChain(AllExceptionSorted)
	return AllExceptionSorted, AllExceptionSortedSum
}

func MakeGoldenVestaChain() ([]mvx.BalanceESDT, *p.Decimal) {
	GoldenVesta, _ := bloom.CreateVestaChain(mvx.VestaGold)
	GoldenVestaTrim := bloom.CreateVestaAmountChains(GoldenVesta)
	GoldenVestaTrimSorted := mvx.SortBalanceIntegerChain(GoldenVestaTrim)
	GoldenVestaTrimSortedSum := mvx.AddBalanceDecimalChain(GoldenVestaTrimSorted)
	return GoldenVestaTrimSorted, GoldenVestaTrimSortedSum
}

func MakeSilverVestaChain() ([]mvx.BalanceESDT, *p.Decimal) {
	SilverVesta, _ := bloom.CreateVestaChain(mvx.VestaSilver)
	SilverVestaTrim := bloom.CreateVestaAmountChains(SilverVesta)
	SilverVestaTrimSorted := mvx.SortBalanceIntegerChain(SilverVestaTrim)
	SilverVestaTrimSortedSum := mvx.AddBalanceDecimalChain(SilverVestaTrimSorted)
	return SilverVestaTrimSorted, SilverVestaTrimSortedSum
}

func MakeBronzeVestaChain() ([]mvx.BalanceESDT, *p.Decimal) {
	BronzeVesta, _ := bloom.CreateVestaChain(mvx.VestaBronze)
	BronzeVestaTrim := bloom.CreateVestaAmountChains(BronzeVesta)
	BronzeVestaTrimSorted := mvx.SortBalanceIntegerChain(BronzeVestaTrim)
	BronzeVestaTrimSortedSum := mvx.AddBalanceDecimalChain(BronzeVestaTrimSorted)
	return BronzeVestaTrimSorted, BronzeVestaTrimSortedSum
}

//Raw Functions

func RawCodingDivisionRewardsChain(Amount *p.Decimal) []mvx.BalanceESDT {
	A1 := sm.MULxc(Amount, p.NFS("0.50"))
	A2 := sm.MULxc(Amount, p.NFS("0.15"))
	A3 := sm.SUBxc(Amount, sm.ADDxc(A1, A2))
	R1 := mvx.BalanceESDT{Address: mvx.CodingDivisionDAO, Balance: sm.DTS(A1)}
	R2 := mvx.BalanceESDT{Address: mvx.SnakeDAO, Balance: sm.DTS(A2)}
	R3 := mvx.BalanceESDT{Address: mvx.DHV1, Balance: sm.DTS(A3)}
	Result := []mvx.BalanceESDT{R1, R2, R3}
	return Result
}

func RawVestaRewardsChain(Amount *p.Decimal) []mvx.BalanceESDT {
	A1 := sm.MULxc(Amount, p.NFS("0.45"))
	A2 := sm.MULxc(Amount, p.NFS("0.05"))
	A3 := sm.MULxc(Amount, p.NFS("0.15"))
	A4 := sm.SUBxc(Amount, sm.SUMxc(A1, A2, A3))
	R1 := mvx.BalanceESDT{Address: mvx.VestaDAO, Balance: sm.DTS(A1)}
	R2 := mvx.BalanceESDT{Address: mvx.DHV3, Balance: sm.DTS(A2)}
	R3 := mvx.BalanceESDT{Address: mvx.SnakeDAO, Balance: sm.DTS(A3)}
	R4 := mvx.BalanceESDT{Address: mvx.DHV1, Balance: sm.DTS(A4)}
	Result := []mvx.BalanceESDT{R1, R2, R3, R4}
	return Result
}

func RawLiquidStakingRewardsChain(Amount *p.Decimal) []mvx.BalanceESDT {
	A1 := sm.MULxc(Amount, p.NFS("0.20"))
	A2 := sm.MULxc(Amount, p.NFS("0.28"))
	A3 := sm.MULxc(Amount, p.NFS("0.12"))
	A4 := sm.MULxc(Amount, p.NFS("0.36"))
	A5 := sm.SUBxc(Amount, sm.SUMxc(A1, A2, A3, A4))
	R1 := mvx.BalanceESDT{Address: mvx.Synchronic, Balance: sm.DTS(A1)}
	R2 := mvx.BalanceESDT{Address: mvx.DHV1, Balance: sm.DTS(A2)}
	R3 := mvx.BalanceESDT{Address: mvx.SnakeDAO, Balance: sm.DTS(A3)}
	R4 := mvx.BalanceESDT{Address: mvx.VestaDAO, Balance: sm.DTS(A4)}
	R5 := mvx.BalanceESDT{Address: mvx.DHV3, Balance: sm.DTS(A5)}
	Result := []mvx.BalanceESDT{R1, R2, R3, R4, R5}
	return Result
}

//Raw Distribution Functions

func DistributeCodingDivisionRewards(Amount *p.Decimal) {
	RewardsChain := RawCodingDivisionRewardsChain(Amount)

	//Export Files
	EVDName := RewardExport(RewardsChain, "rCDs", sm.DTS(Amount))

	//Evidence
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode4, Payee4, p.NFS("0"), Amount, mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, RewardsChain, Evidence)

	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")

}

func DistributeVestaRewards(Amount *p.Decimal) {
	RewardsChain := RawVestaRewardsChain(Amount)

	//Export Files
	EVDName := RewardExport(RewardsChain, "rVSs", sm.DTS(Amount))

	//Evidence
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode5, Payee5, p.NFS("0"), Amount, mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, RewardsChain, Evidence)

	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")

}

func DistributeLiquidStakeRewards(Amount *p.Decimal) {
	RewardsChain := RawLiquidStakingRewardsChain(Amount)

	//Export Files
	EVDName := RewardExport(RewardsChain, "rLSs", sm.DTS(Amount))

	//Evidence
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode6, Payee7, p.NFS("0"), Amount, mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, RewardsChain, Evidence)

	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")

}

// CD Mixed Function

func ComputeCodingDivisionRewards(Amount *p.Decimal) []mvx.BalanceESDT {

	//Make Snakes Chain
	Snakes, SnakesSum := MakeSnakeChain()

	//Make Coding Division Chain
	CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()

	SnakeNumber := p.NFS("760")
	ASSE := p.NFS("75") //Ancient Snake Share Equivalent

	TotalAllShares := p.NFS("1000000")
	TotalExternalShares := p.NFS("0")
	CompanyShares := sm.SUBxc(TotalAllShares, TotalExternalShares)
	SnakeShare := sm.DIVxc(CompanyShares, p.NFS("2000"))

	ExistingSnakeShares := sm.MULxc(SnakeShare, SnakeNumber)

	AncientShare := sm.MULxc(ASSE, SnakeShare)

	FoundersShare := sm.SUBxc(CompanyShares, sm.SUMxc(AncientShare, ExistingSnakeShares))
	FoundersShareEach := sm.DIVxc(FoundersShare, p.NFS("2"))

	fmt.Println("Total Existing Shares are ", TotalAllShares, " shares.")
	fmt.Println("External Entities have", TotalExternalShares, " shares.")
	fmt.Println("Demiourgos holds ", CompanyShares, " shares")
	fmt.Println("One Snake is ", SnakeShare, "shares.")
	fmt.Println(SnakeNumber, " Snakes have ", ExistingSnakeShares, " shares.")
	fmt.Println("CEO and HoV have ", FoundersShareEach, "shares each.")
	fmt.Println("CTO has ", AncientShare, " shares.")

	CodingDivisionDistribution := sm.MULxc(Amount, p.NFS("0.5"))
	SnakesDistribution := sm.MULxc(sm.DIVxc(SnakeNumber, p.NFS("1000")), sm.DIVxc(CodingDivisionDistribution, p.NFS("2")))
	PDistribution := sm.MULxc(sm.DIVxc(FoundersShareEach, CompanyShares), CodingDivisionDistribution)
	ADistribution := sm.MULxc(sm.DIVxc(AncientShare, CompanyShares), CodingDivisionDistribution)

	CDSplit := []*p.Decimal{CodingDivisionDistribution, SnakesDistribution, PDistribution, ADistribution}
	fmt.Println("=====*****======")
	fmt.Println("Coding.Division gets ", CodingDivisionDistribution, " Tokens.")
	fmt.Println("Snakes get ", SnakesDistribution, " Tokens.")
	fmt.Println("CEO and HoV get ", PDistribution, " Tokens each as Shareholders.")
	fmt.Println("CTO gets ", ADistribution, " Tokens as Shareholder.")
	fmt.Println("=====*****======")
	fmt.Println("")

	SnakeRewardsChain := mvx.ExactPercentualDecimalRewardSplitter(SnakesDistribution, Snakes)
	CodingDivisionRewardsChain := mvx.ExactPercentualDecimalRewardSplitter(CodingDivisionDistribution, CodingDivision)

	SharesChain := []mvx.BalanceESDT{
		{mvx.Hefe, sm.DTS(PDistribution)},
		{mvx.Florian, sm.DTS(PDistribution)},
		{mvx.AH, sm.DTS(ADistribution)},
	}

	TotalRewards1 := mvx.DecimalChainAdder(SnakeRewardsChain, CodingDivisionRewardsChain)
	TotalRewards := mvx.DecimalChainAdder(TotalRewards1, SharesChain)
	TotalRewardsSorted := mvx.SortBalanceDecimalChain(TotalRewards)

	for i := 0; i < len(TotalRewardsSorted); i++ {
		fmt.Println(TotalRewardsSorted[i])
	}

	//Export Files
	EVDName := RewardExport(TotalRewardsSorted, "tCDr", sm.DTS(Amount))

	//Make Evidence and Export it
	Evidence := MakeTotalCDEvidence(DistributionType1, DistributionMode1, Payee3, Payee2, CodingDivisionSum, SnakesSum, Amount, mvx.WrappedEGLD, CDSplit)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceCD(EVDName, Snakes, CodingDivision, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")

	return TotalRewards
}

// Multiplication Rewards

func ComputeSnakeRewardsByMultiplication(MultiplicationAmount string) []mvx.BalanceESDT {
	Snakes, SnakesSum := MakeSnakeChain()
	SnakesRewards := mvx.RewardsComputerDecimalChainMultiplication(Snakes, MultiplicationAmount)
	EVDName := RewardExport(SnakesRewards, "sSr", MultiplicationAmount)

	//Make Evidence and Export it
	Evidence := MakeMultiplicationEvidence(DistributionType2, DistributionMode2, Payee2, SnakesSum, p.NFS(MultiplicationAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplicationSFT(EVDName, Snakes, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return SnakesRewards
}

func ComputeCodingDivisionRewardsByMultiplication(MultiplicationAmount string) []mvx.BalanceESDT {
	CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()
	CodingDivisionRewards := mvx.RewardsComputerDecimalChainMultiplication(CodingDivision, MultiplicationAmount)
	EVDName := RewardExport(CodingDivisionRewards, "sCDr", MultiplicationAmount)

	//Make Evidence and Export it
	Evidence := MakeMultiplicationEvidence(DistributionType2, DistributionMode3, Payee3, CodingDivisionSum, p.NFS(MultiplicationAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, CodingDivision, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return CodingDivisionRewards
}

func ComputeCustomSFTRewardsByMultiplication(SFTDesignation string, MultiplicationAmount string) []mvx.BalanceESDT {
	ScanSFT := mvx.SFT(SFTDesignation)
	SFTChain := mvx.SnapshotSFTChain(ScanSFT)
	CustomSFTChain := mvx.SortBalanceIntegerChain(SFTChain)
	CustomSFTSum := mvx.AddBalanceIntegerChain(SFTChain)

	CustomSFTChainRewards := mvx.RewardsComputerDecimalChainMultiplication(CustomSFTChain, MultiplicationAmount)
	EVDName := RewardExport(CustomSFTChainRewards, "cSFTr", MultiplicationAmount)

	//Make Evidence and Export it
	Evidence := MakeMultiplicationEvidence(DistributionType2, DistributionMode3, Payee3, CustomSFTSum, p.NFS(MultiplicationAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, CustomSFTChain, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return CustomSFTChainRewards
}

//Totalisation Rewards

func ComputeSnakeRewardsByTotalisation(TotalAmount string) []mvx.BalanceESDT {
	Snakes, SnakesSum := MakeSnakeChain()
	SnakesRewards := mvx.ExactPercentualDecimalRewardSplitter(p.NFS(TotalAmount), Snakes)
	EVDName := RewardExport(SnakesRewards, "taSr", TotalAmount)

	//Make Evidence and Export it
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode2, Payee2, SnakesSum, p.NFS(TotalAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplicationSFT(EVDName, Snakes, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return SnakesRewards
}

func ComputeCodingDivisionRewardsByTotalisation(TotalAmount string) []mvx.BalanceESDT {
	CodingDivision, CodingDivisionSum := MakeCodingDivisionChain()
	CodingDivisionRewards := mvx.ExactPercentualDecimalRewardSplitter(p.NFS(TotalAmount), CodingDivision)
	EVDName := RewardExport(CodingDivisionRewards, "taCDr", TotalAmount)

	//Make Evidence and Export it
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode3, Payee3, CodingDivisionSum, p.NFS(TotalAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)
	ExportEvidenceMultiplication(EVDName, CodingDivision, Evidence)
	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return CodingDivisionRewards
}

func ComputeVestaRewardsByTotalisation(TotalAmount string, AllChains bool) []mvx.BalanceESDT {
	var (
		GoldRewards, SilverRewards, BronzeRewards, VestaSFTNumber, BVTS *p.Decimal
		BronzeChainRewards, BVT2                                        []mvx.BalanceESDT
		BronzeVestaTrim                                                 []mvx.BalanceSFT
		VestaRewards                                                    []mvx.BalanceESDT
	)

	if AllChains == true {
		GoldRewards = sm.TruncateCustom(sm.MULxc(p.NFS(TotalAmount), p.NFS("0.333333333333333333")), 18)
		SilverRewards = GoldRewards
		BronzeRewards = sm.SUBxc(p.NFS(TotalAmount), sm.ADDxc(GoldRewards, SilverRewards))

		BronzeVesta, _ := bloom.CreateVestaChain(mvx.VestaBronze)
		BronzeVestaTrim = bloom.CreateVestaAmountChains(BronzeVesta)
		BVT2 = mvx.SortBalanceIntegerChain(BronzeVestaTrim)
		BVTS = mvx.AddBalanceDecimalChain(BVT2)

		BronzeChainRewards = mvx.ExactPercentualDecimalRewardSplitter(BronzeRewards, BVT2)
	} else {
		GoldRewards = sm.TruncateCustom(sm.MULxc(p.NFS(TotalAmount), p.NFS("0.5")), 18)
		SilverRewards = GoldRewards
		BronzeRewards = p.NFS("0")
	}

	GoldenVesta, _ := bloom.CreateVestaChain(mvx.VestaGold)
	GoldenVestaTrim := bloom.CreateVestaAmountChains(GoldenVesta)
	GVT2 := mvx.SortBalanceIntegerChain(GoldenVestaTrim)
	GVTS := mvx.AddBalanceDecimalChain(GVT2)

	SilverVesta, _ := bloom.CreateVestaChain(mvx.VestaSilver)
	SilverVestaTrim := bloom.CreateVestaAmountChains(SilverVesta)
	SVT2 := mvx.SortBalanceIntegerChain(SilverVestaTrim)
	SVTS := mvx.AddBalanceDecimalChain(SVT2)

	GoldenChainRewards := mvx.ExactPercentualDecimalRewardSplitter(GoldRewards, GVT2)
	SilverChainRewards := mvx.ExactPercentualDecimalRewardSplitter(SilverRewards, SVT2)

	if AllChains == true {
		Sum := mvx.DecimalChainAdder(VestaRewards, BronzeChainRewards)
		VestaRewards = mvx.DecimalChainAdder(Sum, BronzeChainRewards)
		VestaSFTNumber = sm.SUMxc(GVTS, SVTS, BVTS)
	} else {
		VestaRewards = mvx.DecimalChainAdder(GoldenChainRewards, SilverChainRewards)
		VestaSFTNumber = sm.SUMxc(GVTS, SVTS)
	}

	VestaRewardsSorted := mvx.SortBalanceDecimalChain(VestaRewards)

	EVDName := RewardExport(VestaRewardsSorted, "taVSr", TotalAmount)

	//Make Evidence and Export it
	Evidence := MakeTotalisationEvidence(DistributionType3, DistributionMode3, Payee6, VestaSFTNumber, p.NFS(TotalAmount), mvx.WrappedEGLD)
	EvidenceString := DistributionEvidenceMLS(Evidence)
	fmt.Println(EvidenceString)

	if AllChains == true {
		ExportEvidenceTripleVesta(EVDName, GVT2, SVT2, BVT2, Evidence)
	} else {
		ExportEvidenceDoubleVesta(EVDName, GVT2, SVT2, Evidence)
	}

	//Copy Exported Evidence to RewardFolder
	B, _ := mvx.MyCopy(EVDName, RewardPath+EVDName)
	fmt.Println(B, " bytes copied for the reward file!")
	return VestaRewards
}

// CSV Functions Rewards Export
// If Type is true extension is csv, else is txt

func OutputRewardsName(RewardName, Amount string) (R1, R2 string) {
	currentTime := time.Now()
	T := currentTime.Format("2006;(01-January);(02-Day)_(T-15h;04m;05s)")
	AmountToPrint := "[" + Amount + "]"
	R1 = RewardName + "_" + AmountToPrint + "_" + T + ".csv"
	R2 = RewardName + "_" + AmountToPrint + "_" + T + ".evidence"
	return
}

func OutputDualRewardsName(RewardName, Amount string) (CSV, EVD string) {
	CSV, EVD = OutputRewardsName(RewardName, Amount)
	return
}

func RewardExport(InputChain []mvx.BalanceESDT, RewardName, Amount string) string {
	OutputNameCSV, OutputNameEVD := OutputDualRewardsName(RewardName, Amount)
	mvx.ConvertToBulkCSV(OutputNameCSV, InputChain)

	//Copy Files in their folders
	B2, _ := mvx.MyCopy(OutputNameCSV, RewardPath+OutputNameCSV)
	fmt.Println(B2, " bytes copied for the CSV!")
	return OutputNameEVD
}
