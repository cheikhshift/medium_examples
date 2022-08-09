package main

import (
	"fmt"
	"log"
)


func main() {

	fmt.Println("Hello world, too")
}



// Must add the following comment for it
// to work.
//export Add
func Add(x, y int) int {

	if x < 2 {
		log.Fatal("x must be greater than 1.")
	}

	return x + y
}

