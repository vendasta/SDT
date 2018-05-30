package main

import (
	"net/http"
	"fmt"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "%d", Fibonacci(35))
	})

	srv := &http.Server{Addr: ":11000", Handler: mux}
	srv.ListenAndServe()
}

func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
