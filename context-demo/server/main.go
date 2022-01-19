package main

import (
	"context-demo/context-demo/logger.go"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	// http.HandleFunc("/", handler)
	http.HandleFunc("/", logger.Decorator(handler))
	log.Panic(http.ListenAndServe("127.0.0.1:3000", nil))
}

// every time when we get | recieve request in handler we need to make a request id and then value of
// request id in the context.value

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	// after 2 sec, call func, but ctx timeout is set of 1 sec
	case <-time.After(2 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		// log.Println(ctx.Err().Error())
		logger.Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
	fmt.Println(w, "hello")
}
