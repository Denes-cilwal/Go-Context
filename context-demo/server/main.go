package main

import (
	"context-demo/context-demo/logger"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", logger.Decorate(handler))
	log.Panic(http.ListenAndServe("127.0.0.1:3000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println(w, "hello")
	case <-ctx.Done():
		logger.Println(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
	fmt.Println(w, "hello")
}
