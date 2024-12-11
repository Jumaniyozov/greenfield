package main

import (
	"fmt"
	"github.com/jumaniyozov/greenfield/handlers"
	"github.com/jumaniyozov/greenfield/handlers/rest"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"), "error")
	if addr == ":" {
		addr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
