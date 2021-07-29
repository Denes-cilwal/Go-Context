package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
// cancel function not captured

func main() {
	// initialize the new context
	ctx := context.Background()
	// takes parent context and returns new context and cancel func
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second)
		// sleep for 1 sec and cancel
		cancel()
	}()

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("hello")
	})
}

*/

// with cancel been captured by done()
func main() {
	// initialize the new context
	ctx := context.Background()
	// takes parent context and returns new context and cancel func
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second)
		// sleep for 1 sec and cancel
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("hello")

		// catch the context has been cancelled
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
