// Package Ozon_fintech implement http server for generation zip link
package Ozon_fintech

import (
	"context"
	"net/http"
	"time"
)

// A Server is small shell of http.Server
type Server struct {
	httpServer *http.Server
}

// Run starts the server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 28,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// ShutDown need to gracefully shutdown
func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
