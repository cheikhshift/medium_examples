package main

import (
	lib "multi-os"
	"log"
)

func main(){

	e, err := lib.GetExe()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("System : ", e.OS, "Path :", e.Path)
}