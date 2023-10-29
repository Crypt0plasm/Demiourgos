package Bloodshed

import (
	p "Firefly-APD"
	mvx "MvxApiScanner"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//EstarAPI Based NFT Structure

type EstarNFT []struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type EstarIndividualNFT struct {
	Address mvx.MvxAddress
	ID      int64
}

// The Snapshot format is a chain of this structure.
type EstarIndividualNFTChain struct {
	Address mvx.MvxAddress
	ID      []int64
}

// The Snapshot Score format is a chain of this structure.
type EstarIndividualNFTScoreChain struct {
	Address mvx.MvxAddress
	ID      []TTVScore
}

// The Chain with NFT individual Score Values.
type EstarIndividualNFTValueChain struct {
	Address mvx.MvxAddress
	ID      []*p.Decimal
}

//Creates a Snapshot Chain for a given NFT Collection using Estar API.
//Input1: NFT Collection, for example: BLOODSHED-a62781
//Input2: NFT Collection Units, for example 8861 for bloodshed.

func MakeNFTSnapshotChain(CollectionID string, MaxNftNonce int64) []EstarIndividualNFTChain {
	var (
		Output []EstarIndividualNFTChain
		Owner  EstarIndividualNFT
	)
	FirstOwner := GetNFTOwner(CollectionID, 1)
	SecondOwner := GetNFTOwner(CollectionID, 2)
	AddFirstTwo := AddNFTOwner(FirstOwner, SecondOwner)
	Output = append(Output, AddFirstTwo...)

	//8862
	for i := int64(3); i < MaxNftNonce+1; i++ {
		Owner = GetNFTOwner(CollectionID, int64(i))
		Output = AddNFTOwnerToChain(Owner, Output)
		fmt.Println("Snapshoting Nonce ", i)
	}
	return Output

}

// Scans for a single NFT using the Estar API.
// Saves Output in a EstarIndividualNFT structure

func GetNFTOwner(CollectionID string, MaxNftNonce int64) EstarIndividualNFT {
	var (
		OutputChain EstarNFT
		Output      EstarIndividualNFT
	)
	V1 := "https://mvx-api.estar.games/nfts/"
	V2 := "/accounts"
	NftID := CollectionID + "-" + mvx.MvxNftId(MaxNftNonce)
	Link := V1 + NftID + V2

	SS := mvx.OnPage(Link)
	_ = json.Unmarshal([]byte(SS), &OutputChain)

	Owner := mvx.MvxAddress(OutputChain[0].Address)

	Output.Address = Owner
	Output.ID = MaxNftNonce

	return Output
}

// Adds two EstarIndividualNFT structures in a []EstarIndividualNFTChain

func AddNFTOwner(Owner1, Owner2 EstarIndividualNFT) []EstarIndividualNFTChain {
	var (
		V1, V2       EstarIndividualNFTChain
		SingleOutput EstarIndividualNFTChain
		Output       []EstarIndividualNFTChain
		NonceChain   []int64
	)

	if Owner1.Address == Owner2.Address {
		NonceChain = append(NonceChain, Owner1.ID)
		NonceChain = append(NonceChain, Owner2.ID)
		SingleOutput.Address = Owner1.Address
		SingleOutput.ID = NonceChain
		Output = append(Output, SingleOutput)
	} else if Owner1.Address != Owner2.Address {
		V1 = ConvertToChain(Owner1)
		V2 = ConvertToChain(Owner2)
		Output = append(Output, V1)
		Output = append(Output, V2)
	}
	return Output
}

// Creates a Chain from a single Unit. From a EstarIndividualNFT structure to a EstarIndividualNFTChain structure
func ConvertToChain(Input EstarIndividualNFT) (Output EstarIndividualNFTChain) {
	var Chain []int64
	Chain = append(Chain, Input.ID)
	Output.Address = Input.Address
	Output.ID = Chain
	return Output
}

//Adds an EstarIndividualNFT structure to a []EstarIndividualNFTChain structure

func AddNFTOwnerToChain(Owner EstarIndividualNFT, InputChain []EstarIndividualNFTChain) (OutputChain []EstarIndividualNFTChain) {
	var (
		NonceChain []int64
	)
	//Function to check if owner is in chain
	IzInChain := func(Owner EstarIndividualNFT, InputChain []EstarIndividualNFTChain) (Result bool, Position int) {
		for i := 0; i < len(InputChain); i++ {
			if Owner.Address == InputChain[i].Address {
				Result = true
				Position = i
				break
			} else {
				Result = false
				Position = len(InputChain)
			}
		}
		return Result, Position
	}

	Truth, Position := IzInChain(Owner, InputChain)

	//Owner is in chain
	if Truth == true {
		NonceChain = append(NonceChain, InputChain[Position].ID...)
		NonceChain = append(NonceChain, Owner.ID)
		OutputChain = InputChain
		OutputChain[Position].ID = NonceChain
	} else if Truth == false {
		OwnerToAdd := ConvertToChain(Owner)
		OutputChain = append(InputChain, OwnerToAdd)
	}
	return OutputChain

	//Owner is not in chain
}

// Write NFT Snapshot Chain to a file, exporting its Data.
func ExportNFTSnapshotChain(Name string, List []EstarIndividualNFTChain) {
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

// Read NFT Snapshot Chain, reads the information from an already exported file at the target location

func ImportNFTSnapshotChain(Path string) []EstarIndividualNFTChain {
	var (
		OutputLine EstarIndividualNFTChain
		Output     []EstarIndividualNFTChain
	)

	//Processes a Line of the form of:
	//{erd1p4n2ujjtpth4y9rr8er95pakmamyczk906nazmajhpzqz48e2mjszvkkgj [74 3809 4204]}
	//To a bs.EstarIndividualNFTChain Format

	ProcessLine := func(Line string) EstarIndividualNFTChain {
		var Result EstarIndividualNFTChain

		convertStringToIntSlice := func(input string) []int64 {
			parts := regexp.MustCompile(`\s+`).Split(input, -1)
			var result []int64
			for _, part := range parts {
				val, err := strconv.ParseInt(part, 10, 64)
				if err == nil {
					result = append(result, val)
				}
			}
			return result
		}

		// Regular expressions to match the desired string and integers
		stringPattern := "{([^\\[]+)"
		intPattern := "\\[([0-9 ]+)\\]"

		// Extract the string
		stringRegex := regexp.MustCompile(stringPattern)
		stringMatches := stringRegex.FindStringSubmatch(Line)

		Addy := mvx.MvxAddress(strings.TrimSpace(stringMatches[1]))

		// Extract the integers
		intRegex := regexp.MustCompile(intPattern)
		intMatches := intRegex.FindStringSubmatch(Line)

		intStr := intMatches[1]
		intStr = regexp.MustCompile(`\s+`).ReplaceAllString(intStr, " ")
		IntSlice := convertStringToIntSlice(intStr)

		Result.Address = Addy
		Result.ID = IntSlice

		return Result
	}

	readFile, err := os.Open(Path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	_ = readFile.Close()

	//Creating the []bs.EstarIndividualNFTChain Chain
	for i := 0; i < len(fileLines); i++ {
		Line := fileLines[i]
		OutputLine = ProcessLine(Line)
		Output = append(Output, OutputLine)
	}

	return Output
}

// Scans a Snapshot for an address and returns its position
func GetAddyPosition(Addy mvx.MvxAddress, Input []EstarIndividualNFTChain) int {
	var Output int

	for i := 0; i < len(Input); i++ {
		//fmt.Println("Input de i este", i, Input[i])
		if Addy == Input[i].Address {
			Output = i
			break
		}
	}

	return Output
}
