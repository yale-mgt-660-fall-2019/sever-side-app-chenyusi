package main

import (
	"fmt"
	"net/http"
	"os"
)

// This is a simple function that gets a variable
// from the environment and returns a default if
// none is found. On Heroku and Cloud9 your HTTP
// server needs to listen on the port provided in
// the PORT environment variable in order to be
// accessible.
func getEnv(key, fallback string) string {
	value, foundValue := os.LookupEnv(key)
	if foundValue {
		return value
	}
	return fallback
}

// indexHandler - a "handler" or "controller" that accepts
// an http response and an http request. Information, like
// messages, data, or HTML is written to the response.
// For more: https://golang.org/doc/articles/wiki/#tmp_3
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Write this clever message to w, which implements
	// the Writer interface https://golang.org/pkg/io/#Writer
	fmt.Fprintf(w, "Hellow MGT656/660 FTW!@$#")
}

func main() {
	// Say that when we receive a request for the '/' (or "root") URL
	// we want the function `indexHandler` to handle it.
	http.HandleFunc("/", indexHandler)

	// Start listening for HTTP requests.
	http.ListenAndServe(":"+getEnv("PORT", "8080"), nil)
}
