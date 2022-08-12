package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/godbus/dbus/v5"
)

type BusLock struct {
	conn *dbus.Conn
	mu  *sync.Mutex
	id   string
	//timeout field
}

func NewBusLock() (b BusLock, err error) {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(os.Stderr, "Failed to connect to session bus:", err)
	}

	err = conn.Export(true, "/medium/examples/lock", "medium.examples.lock")

	if err != nil {
		return
	}

	if err = conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/medium/examples/lock"),
	); err != nil {
		return
	}

	ids := conn.Names()
	var m sync.Mutex

	b = BusLock{
		conn: conn,
		id:   ids[0],
		mu : &m,
	}

	return
}

func (b BusLock) Listen() {

	c := make(chan *dbus.Signal, 10)

	b.conn.Signal(c)

	fmt.Println("listenning")

	for v := range c {

		if v.Sender == b.id {
			continue
		}

		action := v.Body[0].(bool)
		fmt.Println("From : ", v.Sender, v)

		if action {
			b.mu.Lock()
			fmt.Println("Locking")
			continue
		}

		fmt.Println("UnLocking")
		b.mu.Unlock()
		
	}
}

func (b *BusLock) Lock (){

	b.mu.Lock()
	b.conn.Emit("/medium/examples/lock", "medium.examples.lock", true)
}

func (b *BusLock) UnLock () {

	b.mu.Unlock()
	b.conn.Emit("/medium/examples/lock", "medium.examples.lock", false)
}

func main() {

	bl, err := NewBusLock()

	if err != nil {
		log.Fatal(err)
	}

	go bl.Listen()

	go func(){

		  now := time.Now()
		  // wait for other instance to set lock
	      time.Sleep(5 * time.Second)

	      bl.Lock()
	      time.Sleep(10 * time.Second)
	      bl.UnLock()

	      fmt.Println( time.Since(now) )

	}()


	c := make(chan int)

	<-c

}
