package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct{
	httpServer *http.Server 
}

func (s *Server) Run(port string) error {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	s.httpServer = &http.Server{
		Addr:           ":"+port,
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Closedown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}