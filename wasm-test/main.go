package main

import (
	"fmt"
	"syscall/js"
)

func Sample(this js.Value, args []js.Value) interface{} {

	fmt.Println(this, args)
	js.Global().Get("document").Call("getElementById", "output").Set("innerHTML", "Hello World")

	return nil
}

func main() {

	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	js.Global().Set("sample", js.FuncOf(Sample))

	<-c
}
