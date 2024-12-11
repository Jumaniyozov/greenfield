package main

import (
	"fmt"
	"github.com/jumaniyozov/greenfield/handlers"
	"github.com/jumaniyozov/greenfield/handlers/rest"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("listening on %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
