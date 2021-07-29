package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	// func context.WithDeadliine(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
	// takes context(parent and timeDuratio) and return new ctx and cancel func
	// from 2 sec now add the deadline
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("hello")

		// when cance; function gets called the channel which calls the cancel function
		// is sent to that done channel
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	}
}

// case here is you are timeing out after 2 sec but function here is actually taking 3 sec to respond
