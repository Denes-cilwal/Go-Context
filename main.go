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
	wg.Add(2)
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

/* output
ideal wait would different form non-ideal case:

StartTime: 2021-07-27 18:07:55.357309 +0545 +0545 m=+0.000111459
Producer # 1 ran for 82ms
±Producer # 2 ran for 888ms
±None-ideal wait took1.006546375s
StartTime: 2021-07-27 18:07:56.364134 +0545 +0545 m=+1.006976251
Producer # 1 ran for 60ms
±Producer # 2 ran for 848ms
±Ideal wait took850.397708ms
*/
