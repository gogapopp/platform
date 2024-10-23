package main

import (
	"log"
	"net/http"
	httpserver "platform/internal/http-server"
	"time"
)

func main() {
	router := httpserver.Routes()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
