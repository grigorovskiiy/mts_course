package main

import (
	"fmt"
	"net/http"
	"task2/task2/server"
	"time"
)

func main() {
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  0 * time.Second,
		WriteTimeout: 0 * time.Second,
	}

	mux.Handle("/version", http.HandlerFunc(server.VersionHandler))
	mux.Handle("/decode", http.HandlerFunc(server.DecodeHandler))
	mux.Handle("/hard-op", http.HandlerFunc(server.HardopHandler))
	fmt.Println("Ready to serve at", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
