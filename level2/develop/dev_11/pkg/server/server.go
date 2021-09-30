package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server struct implements sever
type Server struct {
	httpServer *http.Server
}

// Run runs server
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr: ":" + port,
		// Handler:        handler, // Юзаем дефолтный хендлер, на который накручиваются все ручки;
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Println("Server is running on port: ", port)
	return s.httpServer.ListenAndServe()
}

// Shutdown shutdowns server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
