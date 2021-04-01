package main

import (
	"Laboratory/gSheets/services"
	"log"
	"os"
)

func main() {
	sheetURL := "https://docs.google.com/spreadsheets/d/e/2PACX-1vS2hjxKTSevp2Jq6BmX8eR7ykoEnKO0GUNy11IWTLfIBdAsq98ic7D1XEtaxkUl5ialflfSsq3DiRIb/pub?gid=0&single=true&output=csv"

	fileLocation := "TopSpender.csv"

	err := services.Download(sheetURL, fileLocation, 5000)
	if err != nil {
		log.Println("MAIN ERROR:", err)
	}

	err = services.SendToFirestore(fileLocation)
	err = os.Remove(fileLocation)
	if err != nil {
		log.Fatalln(err)
	}
}
