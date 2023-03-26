package main

import (
	bs "Demiourgos/Bloodshed"
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

func main() {
	ExportBloodshedJSON(1, 152, "Legendary")
	ExportBloodshedJSON(153, 554, "Epic")
	ExportBloodshedJSON(555, 1157, "Rare")
	ExportBloodshedJSON(1158, 8861, "Common")
}
