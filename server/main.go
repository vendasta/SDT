package main

import (
	"net/http"
	"fmt"
	"net"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})

	lis, err := net.Listen("tcp", ":11000")
	if err != nil {
		log.Fatal("Error listening on port 80")
	}

	http.Serve(lis, mux)
}
