package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func main() {
	/**

	creating sequential function call
	 producer(1)
	 producer(2)

	 **/

	/**

	Concurrent Process
	creating goroutine :(named func)

	**/

	// non-ideal way to wait complete for all goroutines to complete
	start := time.Now()
	fmt.Println("StartTime:", start)
	// gorun time to launch goroutine to run producer function with paramter 1
	go producer(1)

	// gorun time to launch goroutine to run producer function with paramter 2
	go producer(2)

	//  pitfall :  but before any of this function runs, main function returns
	// immediately(ends main function and all goroutines get killed)

	// give goroutines time to complete work
	time.Sleep(1 * time.Second)
	elaspe := time.Now().Sub(start)
	fmt.Printf("None-ideal wait took%v\n", elaspe)

	// ideal way to wait complete for all goroutines to complete
	start = time.Now()

	// initialize to wait on two goroutine
	// wg.Add(2)
	fmt.Println("StartTime:", start)
	// gorun time to launch goroutine to run producer function with paramter 1
	go producer2(1)

	// gorun time to launch goroutine to run producer function with paramter 2
	go producer2(2)

	// give goroutines time to complete work
	// time.Sleep(1 * time.Second)
	// use Wait() to complete instead of waitgroup
	wg.Wait()
	elaspe = time.Now().Sub(start)
	fmt.Printf("Ideal wait took%v\n", elaspe)

	// wait for goroutines to complete,
	// kick off new setof gooroutines to complete
	// synchonizing the group of goroutines
	start = time.Now()

	// incorrect wayof adding waitgroup
	wg.Add(2)
	go producer2(3)
	wg.Wait()
	elaspe = time.Now().Sub(start)
	fmt.Printf("Producer2(3) took %v\n", elaspe)

	// // incorrect intializtion of waitgroup
	// wg.Add(2)
	// go producer2(1)
	// wg.Wait()

	launchworkers(5)
	wg.Wait()
}

func producer2(id int) {

	// n is random int between 1 to 1000 inclusive

	n := (rand.Intn(1000) % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n±", id, d)
	//
	wg.Done()
}

func producer(id int) {

	// n is random int between 1 to 1000 inclusive

	n := (rand.Intn(1000) % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer # %v ran for %v\n±", id, d)
}

/**
StartTime: 2021-07-27 18:16:17.294496 +0545 +0545 m=+0.000632585
Producer # 2 ran for 82ms
±Producer # 1 ran for 888ms
±None-ideal wait took1.006440291s
StartTime: 2021-07-27 18:16:18.30093 +0545 +0545 m=+1.007105960
Producer # 1 ran for 60ms
±Producer # 2 ran for 848ms
±Ideal wait took852.757583ms
Producer # 3 ran for 82ms
±Producer2(3) took 87.168458ms

**/

// Adding 2 waitgroup counter but running only one routines lead
/*

±fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x1193740)
        /usr/local/go/src/runtime/sema.go:56 +0x45

**/

// launchworkers creates 'c' goroutines using anonymous functions

func launchworkers(c int) {
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {

			fmt.Printf("launching goroutines annoymously %v\n", i)
			wg.Done()
		}()
	}
}
