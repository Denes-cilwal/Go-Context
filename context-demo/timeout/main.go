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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()

	select {
	// func is taking 3 sec to response
	case <-time.After(3 * time.Second):
		fmt.Println("Hello")
		// but timing out after 2 sec, result context deadline exceed
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
	}

}
