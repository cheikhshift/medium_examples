package main

import (
	"fmt"
	"math"
)

func main(){
	makeChange(1.35)
	makeChange(1.39)
}

func doGoToB(a int){

	b := 20
	
	if a == 100 {
		a *= b
	}
	
	if a != 2000 {
		a -= 1
	}

	fmt.Println(a)
}

func makeChange(a float64){

	// Define fix int array,
	// so the skipped values are accounted
	// for, by defaults
	result := [5]int{}
	number := a * 100

	switch {
		case number >= 100:
			coins := math.Floor(number/100.0)
			number = math.Mod(number,100)
			result[0] = int(coins)
			fallthrough
		case number >= 25:
			coins := math.Floor(number/25.0)
			number = math.Mod(number,25)
			result[1] = int(coins)
			fallthrough
		case number >= 10:
			coins := math.Floor(number/10.0)
			number = math.Mod(number,10)
			result[2] = int(coins)
			fallthrough
		case number >= 5:
			coins := math.Floor(number/5.0)
			number = math.Mod(number,5)
			result[3] = int(coins)
			fallthrough
		case number >= 1:
			coins := math.Floor(number/1.0)
			number = math.Mod(number,1)
			result[4] = int(coins)
			
	}

	fmt.Println(result)
}

func doGoTo(a int){

	b := 20
	
	if a == 100 {
		a *= b
		goto OUTPUT
	}
	
	if a == 2000 {
		a -= 1
	}

OUTPUT:
	fmt.Println(a)
}