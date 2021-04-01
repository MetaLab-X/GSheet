package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Download(url string, filename string, timeout int64) error {
	log.Println("Downloading::", url)

	client := http.Client{
		Timeout: time.Duration(timeout * int64(time.Second)),
	}

	res, err := client.Get(url)

	if err != nil {
		log.Println("Error::", err)
		return err
	}

	cache, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error ReadingAll::", err)
		return err
	}

	err = ioutil.WriteFile(filename, cache, 0644)
	if err != nil {
		log.Println("Error Writing::", err)
		return err
	}

	log.Println("Downloaded:: locate at", filename)

	return err
}
