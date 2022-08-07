package main

import (
	"golang.org/x/sync/singleflight"
	"time"
	"fmt"
)


func timeConsumingFunction() string {

	time.Sleep(10 * time.Second)
	return "Hello"
}

func SayHello[T any](fn func() (interface{}, error)) T {
 var g singleflight.Group

  v, _, _ := g.Do("SayHello", fn)
  return v.(T)
}

func main(){

	var g singleflight.Group
	
	fmt.Println(SayHello[int](func() (interface{}, error){
		return 20,nil
	}))	

	go func(){
		v, _, shared := g.Do("key", func() (interface{}, error) {

			s := timeConsumingFunction()
			return s,nil 
		})
		fmt.Println("Goroutine result : ", v, shared)
	}()

	time.Sleep(2 * time.Second)

	v, _, shared := g.Do("key", func() (interface{}, error) {

		s := timeConsumingFunction()
		return s,nil 
	})

	fmt.Printf("%v %v\n", v.(string), shared)

}