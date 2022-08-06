package main

import (
	"fmt"
	"net"
	"time"
)

type ServerConn struct {
	Connection net.Conn
	ID string
	Open bool
}

func ShowConnection(p * ServerConn){

	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, *p)
	}
	
}


func main() {

	c := make(chan bool)
	p :=  ServerConn{ ID : "first_conn"}

	go ShowConnection(&p)


	go func(){
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ ID : "new_conn"}
			p = newConn
		}
	}()
	

	<- c
}
