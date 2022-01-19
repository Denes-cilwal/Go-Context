package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// create root context
	ctx := context.Background()

	// takes parent context, timeout --- return new ctx and timeout
	// set deadline of 2 secs from now, then timeout
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))

	defer cancel()

	select {
	// func is taking 4 sec to response
	case <-time.After(4 * time.Second):
		fmt.Println("Hello")
		// but timing out after 2 sec, result context deadline exceed
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	}

}
