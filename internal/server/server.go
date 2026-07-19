package server

import (
	"log"
	"net/http"

	"github.com/mostafasaad-m/task-api/internal/config"
	"github.com/mostafasaad-m/task-api/internal/handlers"
)

type Server struct {
	cfg        *config.Config
	auth       *handlers.AuthHandler
	httpServer *http.Server
}

func New(cfg *config.Config, auth *handlers.AuthHandler) *Server {
	srv := &Server{
		cfg:  cfg,
		auth: auth,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", srv.auth.Register)
	mux.HandleFunc("/login", srv.auth.Login)
	srv.httpServer = &http.Server{
		Addr:    "127.0.0.1:" + cfg.AppPort,
		Handler: mux,
	}

	return srv
}

func (s *Server) Run() error {
	log.Printf("Starting Task API on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
