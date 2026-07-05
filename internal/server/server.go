package server

import (
	"log"
	"net/http"

	"github.com/mostafasaad-m/task-api/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:    "127.0.0.1:" + cfg.AppPort,
			Handler: mux,
		},
	}
}

func (s *Server) Run() error {
	log.Printf("Starting Task API on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
