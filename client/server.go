package main

import (
	"net/http"
	"net"
	"fmt"
	"log"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for {
			res, err := http.Get("http://35.232.119.71")
			if err != nil {
				fmt.Fprintf(writer, "Err %s\n", err.Error())
			} else {
				fmt.Fprintf(writer, "Status Code: %d\n", res.StatusCode)
				res.Body.Close()
			}
			time.Sleep(time.Millisecond * 25)
		}
	})

	lis, err := net.Listen("tcp", ":11001")
	if err != nil {
		log.Fatal("Error listening on port 80")
	}
	http.Serve(lis, mux)
}
