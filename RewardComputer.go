package main

import (
	rw "Demiourgos/Rewards"
	p "Firefly-APD"
	"flag"
	"strings"
)

func main() {
	var (
		//Multiplication calls
		SnakeMultiplierRewards = `--snm  <Token Amount> as string;
Computes Snake Rewards by multiplying Snake Amounts with the inputted Token amount.
`
		CodingDivisionMultiplierRewards = `--cdm  <Token Amount> as string;
Computes CodingDivision Rewards by multiplying CD SFT Amounts with the inputted Token amount.
`

		CustomSFTMultiplierRewards = `--sftm  <SFT-Designation_Token-Amount> as string (separated by _);
Computes Custom SFT Rewards by multiplying Custom SFT Amounts with the inputted Token amount.
`

		//Totalisation percentual Splits calls
		SnakeTotalRewards = `--snt  <Token Amount> as string;
Computes Snake Rewards by splitting a total amount to all Snake NFTs.
Used to send Snake DAO Amounts to users.
`
		CodingDivisionTotalRewards = `--cdt  <Token Amount> as string;
Computes CodingDivision Rewards by splitting a total amount to all CD SFTs
Used to send CodingDivision DAO Amounts to users.
`
		VestaTotalRewards = `--vst  <Token Amount> as string;
Computes Vesta Rewards by splitting a total amount to all CD SFTs
Used to send Vesta DAO Amounts to users.
`

		//RAW Distribution calls
		CodingDivisionRaw = `--cdr  <Token Amount> as string;
Computes CodingDivision Raw Reward Distribution.
`
		VestaRaw = `--vsr  <Token Amount> as string;
Computes Vesta Raw Reward Distribution.
`
		LiquidStakingRaw = `--lsr  <Token Amount> as string;
Computes LiquidStaking Raw Reward Distribution.
`
		//Custom ALL Send Calls
		CDRewards = `--cd  <Token Amount> as string;
Computes rewards for Coding Division Distribution, using a total amount for distribution
50% go to all CD SFTs on user addresses
15% go to all Snakes on user addresses (60% of 50% of 50%)
35% remains at Demiourgos.Holdings™
`
	)

	const (
		//Multiplier Rewards
		SNM  = "snm"  //string
		CDM  = "cdm"  //string
		SFTM = "sftm" //string

		//Single Amount Percentual Rewards
		SNT = "snt" //string	*	sending Snake Rewards from Snake DAO
		CDT = "cdt" //string	*	sending CD Rewards from CD DAO
		VST = "vst" //string	*	sending Vesta Rewards from the Vesta.Finance DAO

		//Raw Distribution
		CDR = "cdr" //string	*	Distributing Raw Coding Division Profits
		VSR = "vsr" //string	*	Distributing Raw Vesta Profits
		LSR = "lsr" //string	*	Distributing Liquid Staking Profits

		//Computation Percentual Rewards
		CD = "cd" //string
	)

	FlagSnakeMultiplierRewards := flag.String(SNM, "0", SnakeMultiplierRewards)
	FlagCodingDivisionMultiplierRewards := flag.String(CDM, "0", CodingDivisionMultiplierRewards)
	FlagCustomSFTMultiplierRewards := flag.String(SFTM, "0", CustomSFTMultiplierRewards)

	FlagSnakeTotalRewards := flag.String(SNT, "0", SnakeTotalRewards)
	FlagCodingDivisionTotalRewards := flag.String(CDT, "0", CodingDivisionTotalRewards)
	FlagCodingVestaTotalRewards := flag.String(VST, "0", VestaTotalRewards)

	FlagCodingDivisionRaw := flag.String(CDR, "0", CodingDivisionRaw)
	FlagVestaRaw := flag.String(VSR, "0", VestaRaw)
	FlagLiquidStakingRaw := flag.String(LSR, "0", LiquidStakingRaw)

	FlagCDRewards := flag.String(CD, "0", CDRewards) //Deprecated

	flag.Parse()

	//First Option
	if *FlagSnakeMultiplierRewards != "0" {
		rw.ComputeSnakeRewardsByMultiplication(*FlagSnakeMultiplierRewards)
	}

	//Second Option
	if *FlagCodingDivisionMultiplierRewards != "0" {
		rw.ComputeCodingDivisionRewardsByMultiplication(*FlagCodingDivisionMultiplierRewards)
	}

	//Second Second Option
	if *FlagCustomSFTMultiplierRewards != "0" {
		BlockHeights := strings.Split(*FlagCustomSFTMultiplierRewards, "_")
		//fmt.Println("Primut text este", BlockHeights)
		rw.ComputeCustomSFTRewardsByMultiplication(BlockHeights[0], BlockHeights[1])
	}

	//Third Option
	if *FlagSnakeTotalRewards != "0" {
		rw.ComputeSnakeRewardsByTotalisation(*FlagSnakeTotalRewards)
	}

	//Fourth Option
	if *FlagCodingDivisionTotalRewards != "0" {
		rw.ComputeCodingDivisionRewardsByTotalisation(*FlagCodingDivisionTotalRewards)
	}

	//Fifth Option
	if *FlagCodingVestaTotalRewards != "0" {
		rw.ComputeVestaRewardsByTotalisation(*FlagCodingVestaTotalRewards, false)
	}

	//Sixth Option
	if *FlagCodingDivisionRaw != "0" {
		rw.DistributeCodingDivisionRewards(p.NFS(*FlagCodingDivisionRaw))
	}

	//Seventh Option
	if *FlagVestaRaw != "0" {
		rw.DistributeVestaRewards(p.NFS(*FlagVestaRaw))
	}

	//Eighth Option
	if *FlagLiquidStakingRaw != "0" {
		rw.DistributeLiquidStakeRewards(p.NFS(*FlagLiquidStakingRaw))
	}

	//Ninth Option
	if *FlagCDRewards != "0" {
		rw.ComputeCodingDivisionRewards(p.NFS(*FlagCDRewards))
	}
}
