package main

import (
	"net/http"
	"fmt"
	"github.com/vendasta/SDT/internal"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
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
		fmt.Fprintf(writer, "%d", Fibonacci(35))
	})

	srv := &http.Server{Addr: ":11000", Handler: mux}
	end := make(chan os.Signal)
	signal.Notify(end, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		srv.ListenAndServe()
	}()

	<-end

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	srv.Shutdown(ctx)

}

func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
