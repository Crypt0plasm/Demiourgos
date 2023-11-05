package main

import (
	bs "Demiourgos/Bloodshed"
	p "Firefly-APD"
	mvx "MvxApiScanner"
	"flag"
	"fmt"
	"strings"
)

func main() {

	//Tag Variables
	var (
		TagNFTScan = `--nftscan  <Collection-Name:Collection-Units> as string;
Scans the given NFT Collection and creates a Snapshot-File. The Snapshot-File
will be named Collection-Name_Collection-Units.txt 
For example use BLOODSHED-a62781:8861 for snapshoting the Bloodshed Collection
`
		TagSetScan = `--setscan  <Address> as string;
Assumes a Snapshot is executed and an exported txt file already exists. Looks up for the
Input address in the Snapshot and displays how many Sets the address has.
`
		TagComputeScore = `--bsscore <WEEK:Amount>;
Assumes a Snapshot is executed and an exported txt file already exists. Uses its information
To compute the Bloodshed Score for each ERD Address, and creates an output.
To compute score without week modifier use "none" or "NONE" or "None". Allowed strings are (with caps, without <>):
For Dacian: <Comati> <Ursoi> <Pileati> <Smardoi> <Carpian> <Tarabostes> <Costoboc> <Buridavens>
For Potency: <Standard> <Premium> <Elite>
For Bloodshed Tier: <Tier 0> <Tier 1> <Tier 2> <Tier 3> <Tier 4> <Tier 5>
For Common Backgrounds: <Storm> <Grainfield> <Alpine> <Swamp> <Panonic> <Continental>
For Rare Backgrounds: <Rain> <Steppe> <Pontic>
For Epic Backgrounds: <Fire> <Lightning>
For Legendary Backgrounds: <Wolven Trinity> <Dacian Gryphon> <Gryphon Phalera> <Unicorn Bird> <Eight-Legged Stag> <Ram Sacrifice> <Swirling Horses> <Twin Rams>
For Protection: <Dacian-Skin> <Bear-Skin> <Shield> <Armor>
For Weapons: <Cosor> <Falx> <Sicae> <Pavaza> <Dacian-Draco> <Howler>
If the string has a space included, "" must be used. Examples:
--bssscore "Tier 1:100" or --bsscore Comati:100 or -bsscore "Unicorn Bird:100"
`
	)

	//Constants
	const (
		NFTScan = "nftscan"
		SetScan = "setscan"
		BSSplit = "bsscore"
	)

	//Flags
	FlagNFTScan := flag.String(NFTScan, "0", TagNFTScan)
	FlagSetScan := flag.String(SetScan, "0", TagSetScan)
	FlagBSSplit := flag.String(BSSplit, "0", TagComputeScore)

	flag.Parse()

	//Options

	//Option No 01
	if *FlagNFTScan != "0" {
		InputStringAsSlice := strings.Split(*FlagNFTScan, ":")
		CollectionID := InputStringAsSlice[0]
		MaxNftNonce := mvx.ConvertStringToInt64(InputStringAsSlice[1])
		OutputTXT := CollectionID + "_" + InputStringAsSlice[1] + ".txt"

		//Makes Snapshot Directly
		SnapshotChain := bs.MakeNFTSnapshotChain(CollectionID, MaxNftNonce)

		bs.ExportNFTSnapshotChain(OutputTXT, SnapshotChain)

	}

	//Option No 02
	if *FlagSetScan != "0" {
		Addy := mvx.MvxAddress(*FlagSetScan)
		SnapshotChain := bs.ImportNFTSnapshotChain("./BLOODSHED-a62781_8861.txt")
		Position := bs.GetAddyPosition(Addy, SnapshotChain)
		_, B := bs.BloodshedSetSorting(bs.ConvertNonceListToScoreList(SnapshotChain[Position].ID))
		fmt.Println("Address ", Addy, " has the following Sets:")
		bs.ReadSetComposition(B)
	}

	//Option No 03
	if *FlagBSSplit != "0" {
		InputStringAsSlice := strings.SplitN(*FlagBSSplit, ":", 2)

		Week := InputStringAsSlice[0]
		Amount := p.NFS(InputStringAsSlice[1])

		SnapshotChain := bs.ImportNFTSnapshotChain("./BLOODSHED-a62781_8861.txt")
		//Conversion from a chain of integers (Nonces), to a chain of TTVScore structures (the reward structure)
		//In the Creation of the TTVScore structure, the Order Multiplier is added.
		SnapshotRewardChain := bs.ConvertChainToScoreChain(SnapshotChain)

		//Adding the Set Multiplier by performing SetSorting on all ERDs in the Snapshot
		SnapshotWithSetMultiplier := bs.AddSetMultiplier(SnapshotRewardChain)

		//Adding The Week Multiplier:
		//Not yet implemented
		SnapshotWithWeekMultiplier := bs.AddWeekMultiplier(Week, SnapshotWithSetMultiplier)

		//Converting the Snapshot with []TTVScore ID to a Snapshot with []*p.Deciaml ID
		//Basically converting each NFT (which exists as TTVScore) to a Bloodshed Score Value (existing as *p.Decimal)
		SnapshotWithIndividualValue := bs.MakeIndividualScoreChain(SnapshotWithWeekMultiplier)
		bs.ExportIndividualScoreValues("BloodshedIndividualScore.txt", SnapshotWithIndividualValue)

		//Computes the final summed value for each address
		FinalOutput := bs.ConvertFinal(SnapshotWithIndividualValue)
		FinalOutputWithExceptions := mvx.MakeExChainFromBalanceESDT(FinalOutput, bs.BloodshedExceptions)

		//Computes Rewards
		BloodshedRewardsChain := mvx.ExactPercentualDecimalRewardSplitter(Amount, FinalOutputWithExceptions)

		//Outputs the computed Score and Rewards file:
		mvx.WriteChainBalanceESDT("BloodshedScore.txt", FinalOutputWithExceptions)
		mvx.WriteChainBalanceESDT("BloodshedRewards.txt", BloodshedRewardsChain)
	}

}
