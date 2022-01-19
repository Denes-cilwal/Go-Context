package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// create a root context
	ctx := context.Background()

	// return new-context and cancel func
	ctx, cancel := context.WithCancel(ctx)

	// start go-routine
	go func() {
		// sleep for 1 sec
		time.Sleep(time.Second)
		// after 1 sec, call cancel func context
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("hello")
		// catch the cancel context
	case <-ctx.Done():
		log.Fatalf(ctx.Err().Error())
	}
}
