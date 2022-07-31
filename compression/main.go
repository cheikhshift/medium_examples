package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {

	var path string
	flag.StringVar(&path, "path", ".", "Path of directory to compress")

	flag.Parse()

	start1 := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go GoCompressDirectory(path, 0, &wg)
	wg.Wait()

	elapsed1 := time.Since(start1)
	log.Printf("Goroutine compression took %s", elapsed1)



	start2 := time.Now()
	CompressDirectory(path, 0)

	elapsed2 := time.Since(start2)
	log.Printf("Normal compression took %s", elapsed2)
}

func CompressDirectory(path string, level int) {

	parts := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf("archive.%d.zip", level)
	zipArchive, _ := os.Create(fileName)
	defer zipArchive.Close()

	writer := zip.NewWriter(zipArchive)

	for _, f := range files {

		fP := filepath.Join(path, f.Name())

		if f.IsDir() {
			if level == 0 {
				parts++
				CompressDirectory(fP, parts)
			}
			continue
		}

		f1, _ := os.Open(fP)
		w1, _ := writer.Create(fP)
		io.Copy(w1, f1)
		f1.Close()

		
	}

	writer.Close()

}

func GoCompressDirectory(path string, level int, wg *sync.WaitGroup) {

	parts := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf("archive.goroutine.%d.zip", level)
	zipArchive, _ := os.Create(fileName)
	defer zipArchive.Close()

	writer := zip.NewWriter(zipArchive)

	for _, f := range files {

		fP := filepath.Join(path, f.Name())

		if f.IsDir() {
			if level == 0 {
				parts++
				wg.Add(1)
				go GoCompressDirectory(fP, parts, wg)
			}
			continue
		}

		f1, _ := os.Open(fP)
		w1, _ := writer.Create(fP)
		io.Copy(w1, f1)
		f1.Close()

	}

	
	writer.Close()
	wg.Done()

}
