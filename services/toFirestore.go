package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func SendToFirestore(csvPath string) error {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Error open csv::", err)
		return err
	}

	csvContent, err := csv.NewReader(csvFile).ReadAll()
	dataStore := []map[string]string{}

	for _, row := range csvContent[1:] {
		if row[1] == "" {
			continue
		}
		dataMap := map[string]string{}
		for j, col := range row[:] {
			key := csvContent[0][j]
			dataMap[key] = col
		}

		fmt.Println(dataMap)
		dataStore = append(dataStore, dataMap)
	}

	return err
}
