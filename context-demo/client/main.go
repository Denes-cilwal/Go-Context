package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// set timout context for 1 secs
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//  sending context across http request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:3000/", nil)
	if err != nil {
		panic(err)
	}

	// Do() sends the httpRequest and responds http requests
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// resp, err := http.Get("https://localhost±:3000/")
	// if err != nil {
	// 	panic(err)
	// }

	// close the response body
	defer resp.Body.Close()

	// write to stdout and read from respose body
	io.Copy(os.Stdout, resp.Body)
}
