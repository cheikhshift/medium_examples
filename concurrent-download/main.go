package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	url     = "http://212.183.159.230/5MB.zip"
	workers = 5
)

type Part struct {
	Data  []byte
	Index int
}

func main() {

	var size int
	results := make(chan Part, workers)
	parts := [workers][]byte{}

	client := &http.Client{}

	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	//defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)

	log.Println("Headers : ", resp.Header["Content-Length"])

	if header, ok := resp.Header["Content-Length"]; ok {
		fileSize, err := strconv.Atoi(header[0])

		if err != nil {
			log.Fatal("File size could not be determined : ", err)
		}

		size = fileSize / workers

	} else {
		log.Fatal("File size was not provided!")
	}

	for i := 0; i < 5; i++ {
		go download(i, size, results)
	}

	counter := 0

	for part := range results {
		counter++

		parts[part.Index] = part.Data
		if counter == workers {
			break
		}
	}

	file := []byte{}

	for _, part := range parts {
		file = append(file, part...)
	}

	// Set permissions accordingly, 0700 may not
	// be the best choice
	err = ioutil.WriteFile("./data.zip", file, 0700)

	if err != nil {
		log.Fatal(err)
	}

}

func download(index, size int, c chan Part) {

	client := &http.Client{}

	start := index * size
	dataRange := fmt.Sprintf("bytes=%d-%d", start, start+size-1)

	if index == workers-1 {
		dataRange = fmt.Sprintf("bytes=%d-", start)
	}

	log.Println(dataRange)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		// code to restart download
		return
	}

	req.Header.Add("Range", dataRange)

	resp, err := client.Do(req)

	if err != nil {
		// code to restart download
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		// code to restart download
		return
	}

	c <- Part{Index: index, Data: body}
}
