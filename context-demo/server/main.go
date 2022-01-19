package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	log.Panic(http.ListenAndServe("127.0.0.1:3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	// after 2 sec, call func, but ctx timeout is set of 1 sec
	case <-time.After(2 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
	fmt.Println(w, "hello")
}
