package logger

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// takes context
func Println(ctx context.Context, msg string) {
	// get value from the context with key
	v, ok := ctx.Value("requestID").(int64)
	fmt.Println(v, "id is")
	if !ok {
		fmt.Println("request id not found")
		return
	}
	log.Printf("%d -> %s", v, msg)
}

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := rand.Int63()
		// get context from request
		ctx := r.Context()

		// set context with key in request context
		ctx = context.WithValue(ctx, "requestID", reqID)

		// WithContext returns a shallow copy of r with its context changed to ctx.
		// The provided ctx must be non-nil.
		f(w, r.WithContext(ctx))
	}
}
