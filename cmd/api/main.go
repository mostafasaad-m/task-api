package main

import (
	"log"
	"net/http"

	"github.com/mostafasaad-m/task-api/internal/config"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Printf("Starting Task API on port %s", cfg.AppPort)

	if err := http.ListenAndServe("127.0.0.1:"+cfg.AppPort, mux); err != nil {
		log.Fatal(err)
	}
}
