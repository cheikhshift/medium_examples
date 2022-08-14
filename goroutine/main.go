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
	counter := atomic.Int32{}
	counter.Store(0)
	m := make(map[string]int)

	locker := LockType{
		LockCount:       &counter,
		CompleteChannel: waitchan,
	}

	c := sync.NewCond(locker)

	fmt.Println("First call")

	go func() {

		fmt.Println("First")
		m["foo"] = 100
		time.Sleep(10 * time.Second)
		c.Broadcast()
	}()

	go func() {

		c.Wait()
		fmt.Println("Second", m["foo"])
		
	}()

	go func() {

		c.Wait()
		fmt.Println("third", m["foo"])
		

	}()

	go func() {

		c.Wait()
		fmt.Println("fourth", m["foo"])

	}()

	<-waitchan
	fmt.Println("Foo")

}
