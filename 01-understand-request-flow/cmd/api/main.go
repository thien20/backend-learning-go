package main

import (
	"log"
	"net/http"
	"time"

	"github.com/thien/backend-learning-go/01-understand-request-flow/internal/app"
)

func main() {
	container := app.NewContainer()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           container.Handler.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("01-understand-request-flow listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("serve http: %v", err)
	}
}
