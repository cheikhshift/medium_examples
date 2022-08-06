package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

type ServerConn struct {
	Connection net.Conn
	ID string
	Open bool
}

func ShowConnection(p * atomic.Pointer[ServerConn]){

	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, p.Load())
	}
	
}


func main() {

	c := make(chan bool)
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ ID : "first_conn"}
	p.Store( &s )

	go ShowConnection(&p)


	go func(){
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ ID : "new_conn"}
			p.Swap(&newConn)
		}
	}()
	

	fmt.Println(p.Load().ID)

	<- c
}
