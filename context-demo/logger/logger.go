package logger

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func Println(ctx context.Context, msg string) {
	// getting context value, since context interface has value function attached to it.
	// here key : requestID
	// value returns the value associated with this context for key, or nil
	// if no value is associated with key
	v, ok := ctx.Value("requestID").(int64)
	if !ok {
		fmt.Println("request id not found")
		return
	}
	log.Println("%d %s", v, msg)
}

// Decorate  function actually takes our hanlder and in this functiuon it will attach this new context
// with value
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// initialize reqid to be sent
		reqID := rand.Int63()

		// get context from request.context()
		ctx := r.Context()

		//add value context to it
		ctx = context.WithValue(ctx, "requestID", reqID)

		// call the function with context request and writer
		f(w, r.WithContext(ctx))
	}
}
