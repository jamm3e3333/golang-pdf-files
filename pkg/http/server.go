package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Port uint32

	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration

	Handler http.Handler
}

type Server struct {
	shutdownTimeout time.Duration

	server *http.Server
}

func NewServer(cfg *Config) *Server {
	return &Server{
		server: &http.Server{
			Handler:      cfg.Handler,
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (s *Server) Start() chan error {
	errChan := make(chan error)
	go func() {
		defer close(errChan)

		err := s.server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
