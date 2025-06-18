package main

import (
	"context"
	"log"
	"net"
	"net/http"
)

func main() {
	port := "5000"
	mux := http.NewServeMux()

	ctx := context.Background()

	server := http.Server{
		Addr:        ":" + port,
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error running server: %v", err)
	}
}
