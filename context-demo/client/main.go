package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// set timeout context of 1 sec
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:3000", nil)
	if err != nil {
		panic(err)
	}

	// take req,return resp
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// defer close request call
	defer resp.Body.Close()

	// copy and write to standard output
	io.Copy(os.Stdout, resp.Body)
}
