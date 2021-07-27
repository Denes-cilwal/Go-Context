package main

import (
	"fmt"
	"time"
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

	// gorun time to launch goroutine to run producer function with paramter 1
	go producer(1)

	// gorun time to launch goroutine to run producer function with paramter 2
	go producer(2)

	//  pitfall :  but before any of this function runs, main function returns
	// immediately(ends main function and all goroutines get killed)

	// solve:
	time.Sleep(1 * time.Second)

	// creating goroutines from anonymous function
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Foo() - message # %v\n", i)
		}
		go producer(3)
	}()

}

//
func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer # %v - message # %v\n", id, i)
	}
}
