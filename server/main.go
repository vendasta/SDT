package main

import (
	"net/http"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/vendasta/SDT/internal"
)

func main() {
	app := sdt.NewApplication()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		if app.Ready() {
			err := app.Ping()
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			} else {
				fmt.Fprint(writer, "OK")
			}
		} else {
			http.Error(writer, "Not Ready", http.StatusBadRequest)
			return
		}
	})

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello")
	})

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{Addr: ":11000", Handler: mux}

	srv.ListenAndServe()

}
