package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Aditya8840/crud-with-golang/internals/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()
	// Database connection
	// Setup routing
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("http.StatusOK"))
	})
	// Start server

	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("Server started at", cfg.HttpServer.Address)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %s", err)
		}
	}()

	<-done

	fmt.Println("Server stopped")
}