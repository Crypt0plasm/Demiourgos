package main

import (
	bs "Demiourgos/Bloodshed"
	mvx "MvxApiScanner"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ExportJson() {
	DataToExport := bs.BSJ{
		Name:        bs.BSName,
		Description: bs.BSDescription,
		//Image:       "",
		Attributes: bs.Legendary7,
	}

	file, _ := json.MarshalIndent(DataToExport, "", " ")
	_ = ioutil.WriteFile("test2.json", file, 0644)
}

func GetPositionString(Position int) string {
	var TotalPositionString string
	if Position < 10 {
		TotalPositionString = "000" + strconv.Itoa(Position)
	} else if Position > 9 && Position < 100 {
		TotalPositionString = "00" + strconv.Itoa(Position)
	} else if Position > 99 && Position < 1000 {
		TotalPositionString = "0" + strconv.Itoa(Position)
	} else if Position > 999 && Position < 10000 {
		TotalPositionString = strconv.Itoa(Position)
	}
	return TotalPositionString
}

func GetPositionStringCustom(Position, Digits int) string {
	var TotalPositionString string

	if Digits == 1 {
		TotalPositionString = strconv.Itoa(Position)
	} else if Digits == 2 {
		if Position < 10 {
			TotalPositionString = "0" + strconv.Itoa(Position)
		} else if Position > 9 && Position < 100 {
			TotalPositionString = strconv.Itoa(Position)
		}
	} else if Digits == 3 {
		if Position < 10 {
			TotalPositionString = "00" + strconv.Itoa(Position)
		} else if Position > 9 && Position < 100 {
			TotalPositionString = "0" + strconv.Itoa(Position)
		} else if Position > 99 && Position < 1000 {
			TotalPositionString = strconv.Itoa(Position)
		}
	}
	return TotalPositionString
}

func ExportBloodshedJSON(Start, Stop int, Type string) {

	var (
		OutputFolderName = "BloodshedJSON"
		DataToExport     bs.BSJ
	)

	GetSize := func(Input string) int {
		var Result int
		if Input == "Common" {
			Result = 144
		} else if Input == "Rare" {
			Result = 72
		} else if Input == "Epic" {
			Result = 48
		} else if Input == "Legendary" {
			Result = 8
		}
		return Result
	}

	GetChainPosition := func(StartPosition, Input int, ChainType string) int {
		return (Input - StartPosition) % GetSize(ChainType)
	}

	for i := Start; i < Stop+1; i++ {
		if Type == "Common" {
			DataToExport = bs.BSJ{
				Name:        bs.BSName + "_" + GetPositionString(i),
				Description: bs.BSDescription,
				Attributes:  bs.CT[GetChainPosition(Start, i, Type)],
			}
		} else if Type == "Rare" {
			DataToExport = bs.BSJ{
				Name:        bs.BSName + "_" + GetPositionString(i),
				Description: bs.BSDescription,
				Attributes:  bs.RT[GetChainPosition(Start, i, Type)],
			}
		} else if Type == "Epic" {
			DataToExport = bs.BSJ{
				Name:        bs.BSName + "_" + GetPositionString(i),
				Description: bs.BSDescription,
				Attributes:  bs.ET[GetChainPosition(Start, i, Type)],
			}
		} else if Type == "Legendary" {
			DataToExport = bs.BSJ{
				Name:        bs.BSName + "_" + GetPositionString(i),
				Description: bs.BSDescription,
				Attributes:  bs.LT[GetChainPosition(Start, i, Type)],
			}
		}

		// Convert the Person struct to a JSON string
		jsonStr, err := json.MarshalIndent(DataToExport, "", " ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		//CreateFolder If it Doesnt Exit
		if _, err := os.Stat(OutputFolderName); os.IsNotExist(err) {
			err = os.Mkdir(OutputFolderName, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				return
			}
		}

		// Write the JSON string to a file
		JsonName := GetPositionString(i) + ".json"
		file, err := os.Create(OutputFolderName + "/" + JsonName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(string(jsonStr))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func ExportNFTAssets(Start, Stop, OutputPosition int) {
	var (
		IPFS             = "https://ipfs.io/ipfs/QmSaZ5WzT4rUBiBj9iLTzpY1YhYyzar7B5J3hexqQgKtd1/"
		String1, String2 string
		StringChain      []string
	)

	GetRarityPosition := func(Input int) (string, string) {
		var (
			Rarity, Position string
			PositionC        int
		)
		//Get Rarity
		if Input >= 1 && Input <= 152 {
			Rarity = "Legendary"
			PositionC = Input % 8
			if PositionC == 0 {
				Position = GetPositionStringCustom(8, 1)
			} else {
				Position = GetPositionStringCustom(PositionC, 1)
			}
		} else if Input >= 153 && Input <= 554 {
			Rarity = "Epic"
			PositionC = (Input - 152) % 48
			if PositionC == 0 {
				Position = GetPositionStringCustom(48, 2)
			} else {
				Position = GetPositionStringCustom(PositionC, 2)
			}
		} else if Input >= 555 && Input <= 1157 {
			Rarity = "Rare"
			PositionC = (Input - 554) % 72
			if PositionC == 0 {
				Position = GetPositionStringCustom(72, 2)
			} else {
				Position = GetPositionStringCustom(PositionC, 2)
			}
		} else if Input >= 1158 && Input <= 8861 {
			Rarity = "Common"
			PositionC = (Input - 1157) % 144
			if PositionC == 0 {
				Position = GetPositionStringCustom(144, 3)
			} else {
				Position = GetPositionStringCustom(PositionC, 3)
			}
		}
		return Rarity, Position
	}

	for i := Start; i < Stop+1; i++ {
		String1 = IPFS + GetPositionString(i) + ".json"
		StringChain = append(StringChain, String1)

		Rarity, Position := GetRarityPosition(i)
		String2 = IPFS + "Bloodshed_" + Rarity + "_" + Position + ".png"
		StringChain = append(StringChain, String2)
	}

	OutputPositionString := GetPositionStringCustom(OutputPosition, 2) + ".txt"
	mvx.WriteStringChain(OutputPositionString, StringChain)
}

func ExportNFTAssetsSplit(Start, Stop, OutputPosition int) {
	var (
		IPFSJSON                                      = "https://ipfs.io/QmRYWhExzecdrvv3T97yHYkqj4LMTLXJ1djhWZdkjNjXA3/"
		IPFS1                                         = "https://ipfs.io/QmS1KaHXjQQekaesq6VnSmmZvUZghgXYGvU2CUwe6tTzAY/"
		IPFS2                                         = "https://ipfs.io/QmSihZSEsRwpatfxG91CQnmc5fk3mJaHHYwYssiKBJ44CC/"
		IPFS3                                         = "https://ipfs.io/Qme1dSRCSSKKV4kV7HQ91zAMgRXuWFrPMAoK86n3cQV1Bv/"
		IPFS4                                         = "https://ipfs.io/QmSPTM4Bn7FAhzHQb2WcoDsmUZxAaASZaej7FNkw7HysoF/"
		String1, String2, PhotosIPFS, Positionstring2 string
		StringChain                                   []string
	)

	GetRarityPosition := func(Input int) (string, int) {
		var (
			Rarity              string
			Position, PositionC int
		)
		//Get Rarity
		if Input >= 1 && Input <= 152 {
			Rarity = "Legendary"
			PositionC = Input % 8
			if PositionC == 0 {
				Position = 8
				//Position = GetPositionStringCustom(8, 1)
			} else {
				Position = PositionC
				//Position = GetPositionStringCustom(PositionC, 1)
			}
		} else if Input >= 153 && Input <= 554 {
			Rarity = "Epic"
			PositionC = (Input - 152) % 48
			if PositionC == 0 {
				Position = 48
				//Position = GetPositionStringCustom(48, 2)
			} else {
				Position = PositionC
				//Position = GetPositionStringCustom(PositionC, 2)
			}
		} else if Input >= 555 && Input <= 1157 {
			Rarity = "Rare"
			PositionC = (Input - 554) % 72
			if PositionC == 0 {
				Position = 72
				//Position = GetPositionStringCustom(72, 2)
			} else {
				Position = PositionC
				//Position = GetPositionStringCustom(PositionC, 2)
			}
		} else if Input >= 1158 && Input <= 8861 {
			Rarity = "Common"
			PositionC = (Input - 1157) % 144
			if PositionC == 0 {
				Position = 144
				//Position = GetPositionStringCustom(144, 3)
			} else {
				Position = PositionC
				//Position = GetPositionStringCustom(PositionC, 3)
			}
		}
		return Rarity, Position
	}

	for i := Start; i < Stop+1; i++ {
		String1 = IPFSJSON + GetPositionString(i) + ".json"
		StringChain = append(StringChain, String1)

		Rarity, Position := GetRarityPosition(i)
		if Rarity == "Common" {
			Positionstring2 = GetPositionStringCustom(Position, 3)
			if Position >= 1 && Position <= 68 {
				PhotosIPFS = IPFS1
			} else if Position >= 69 && Position <= 136 {
				PhotosIPFS = IPFS2
			} else if Position >= 137 && Position <= 144 {
				PhotosIPFS = IPFS3
			}
		} else if Rarity == "Epic" || Rarity == "Legendary" {
			if Rarity == "Epic" {
				Positionstring2 = GetPositionStringCustom(Position, 2)
			} else if Rarity == "Legendary" {
				Positionstring2 = GetPositionStringCustom(Position, 1)
			}
			PhotosIPFS = IPFS3
		} else if Rarity == "Rare" {
			Positionstring2 = GetPositionStringCustom(Position, 2)
			if Position >= 1 && Position <= 4 {
				PhotosIPFS = IPFS3
			} else if Position >= 5 && Position <= 72 {
				PhotosIPFS = IPFS4
			}
		}

		String2 = PhotosIPFS + "Bloodshed_" + Rarity + "_" + Positionstring2 + ".png"

		StringChain = append(StringChain, String2)
	}

	OutputPositionString := GetPositionStringCustom(OutputPosition, 2) + ".txt"
	mvx.WriteStringChain(OutputPositionString, StringChain)
}

func main() {
	ExportBloodshedJSON(1, 152, "Legendary")
	//ExportBloodshedJSON(153, 554, "Epic")
	//ExportBloodshedJSON(555, 1157, "Rare")
	//ExportBloodshedJSON(1158, 8861, "Common")

	//for i := 1; i <= 89; i++ {
	//End := i * 100
	//Start := End - 99
	////For the Split Variant use
	////ExportNFTAssetsSplit(Start, End, i)
	//ExportNFTAssets(Start, End, i)
	//}
}
