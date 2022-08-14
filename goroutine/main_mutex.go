package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type LockType struct {
	LockCount       *atomic.Int32
	CompleteChannel chan bool
}

func (l LockType) Lock() {
	fmt.Println("Lock")
	l.LockCount.Add(-1)

	fmt.Println(l.LockCount.Load(), "Count")

	if l.LockCount.Load() == 0 {
		l.CompleteChannel <- true
	}
}

func (l LockType) Unlock() {
	fmt.Println("Unlock")
	l.LockCount.Add(1)
}

func main() {

	waitchan := make(chan bool)

	var m sync.Mutex

	fmt.Println("First call")

	go func() {

		m.Lock()
		defer m.Unlock()
		time.Sleep(10 * time.Second)
		fmt.Println("First")

	}()

	go func() {

		m.Lock()
	    defer m.Unlock()
		fmt.Println("Second")
		time.Sleep(10 * time.Second)

	}()

	go func() {

		m.Lock()
		defer m.Unlock()
		fmt.Println("third")
		time.Sleep(10 * time.Second)

	}()

	go func() {

		m.Lock()
		defer m.Unlock()
		fmt.Println("fourth")
		time.Sleep(10 * time.Second)

	}()

	<-waitchan
	fmt.Println("Foo")

}
