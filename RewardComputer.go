package main

import (
	rw "Demiourgos/Rewards"
	p "Firefly-APD"
	"flag"
)

func main() {
	var (
		SnakeMultiplierRewards = `--snm  <Token Amount> as string;
Computes Snake Rewards by multiplying Snake Amounts with the inputted Token amount
`
		CodingDivisionMultiplierRewards = `--cdm  <Token Amount> as string;
Computes CodingDivision Rewards by multiplying CD SFT Amounts with the inputted Token amount
`
		SnakeTotalRewards = `--cdm  <Token Amount> as string;
Computes Snake Rewards by splitting a total amount to all Snake NFTs
`
		CodingDivisionTotalRewards = `--cdm  <Token Amount> as string;
Computes CodingDivision Rewards by splitting a total amount to all CD SFTs
`
		CDRewards = `--cd  <Token Amount> as string;
Computes rewards for Coding Division Distribution, using a total amount for distribution
50% go to all CD SFTs on user addresses
15% go to all Snakes on user addresses (60% of 50% of 50%)
35% remains at Demiourgos.Holdings™
`
	)

	const (
		//Multiplier Rewards
		SNRM = "snm" //string
		CDRM = "cdm" //string

		//Single Amount Percentual Rewards
		CDRT = "cdt" //string
		SNRT = "snt" //string

		//Computation Percentual Rewards
		CDR = "cd" //string
	)

	FlagSnakeMultiplierRewards := flag.String(SNRM, "0", SnakeMultiplierRewards)
	FlagCodingDivisionMultiplierRewards := flag.String(CDRM, "0", CodingDivisionMultiplierRewards)
	FlagCodingDivisionTotalRewards := flag.String(CDRT, "0", CodingDivisionTotalRewards)
	FlagSnakeTotalRewards := flag.String(SNRT, "0", SnakeTotalRewards)
	FlagCDRewards := flag.String(CDR, "0", CDRewards)

	flag.Parse()

	//First Option
	if *FlagSnakeMultiplierRewards != "0" {
		rw.ComputeSnakeRewardsByMultiplication(*FlagSnakeMultiplierRewards)
	}

	//Second Option
	if *FlagCodingDivisionMultiplierRewards != "0" {
		rw.ComputeCodingDivisionRewardsByMultiplication(*FlagCodingDivisionMultiplierRewards)
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
	if *FlagCDRewards != "0" {
		rw.ComputeCodingDivisionRewards(p.NFS(*FlagCDRewards))
	}
}
