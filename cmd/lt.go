package main

import (
	"net/http"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			for {
				res, err := http.Get("http://35.202.190.253/")
				if err != nil {
					log.Fatalf("Error %s", err.Error())
				}
				log.Printf("Status: %d", res.StatusCode)
				res.Body.Close()
			}
		}()
	}
	wg.Wait()
}
