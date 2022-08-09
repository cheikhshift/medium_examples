package main

import (
	"fmt"
)



func main() {

	fmt.Println("Hello world, too")
}


func throw2(e *byte, l int)

func throw(e *string)

// Must add the following comment for it
// to work.
//export Add
func Add(x, y int) int {

	if x < 2 {

		err := "X must be greater than 1!" 
		throwError(err, len(err))
	}

	return x + y
}

func throwError(e string, l int){
	firstByte := &(([]byte)(e)[0])
	throw2(firstByte,l)
}
