package main

import "fmt"


type MyObject struct {
	FieldOne string
	FieldTwo int
}

func infWork(counter int){
	// Break
	if counter > 20 {
		return
	}

	counter++
	infWork(counter)
}

func infWorkBad(counter int){
	// Run again
	counter++
	infWorkBad(counter)
}


func main(){

	str := "Hello"
	str1 := "World" + str
	str1 += "huh"
	fmt.Println(str1 + str +  str)
	fmt.Println("Hello" + "World" + "Four" + str)
	fmt.Println("Hello" +  str)

	obj := new(MyObject)
	obj2 := MyObject{ FieldOne : "Hello" }

	fmt.Println(obj, obj2)
}