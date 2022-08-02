package main


import (
	"syscall/js"
	"fmt"
)


func save(this js.Value, args []js.Value) interface{} {
	fmt.Println(args[0], args[1].String())

	value := js.Global().Get("localStorage").Call("setItem", args[0].String(), args[1])

 	println("Save called!")

 	return value
}

func get(this js.Value, args []js.Value) interface{} {
 	println("Get called!")
 	if len(args) == 0 {
 		return nil
 	}

 	value := js.Global().Get("localStorage").Call("getItem", args[0].String())

 	return value
}

func main(){

	c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    
    js.Global().Set("saveData", js.FuncOf(save))
 	js.Global().Set("getData", js.FuncOf(get))

    <-c
}