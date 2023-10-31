package main

import (
	// "fmt"
	// "html/template"
	// "log"
	// "html/template"
	"net/http"
	"time"
)

func main() {
	a:=Server{}
	a.Run("8080")
	// tmpl,_:=template.ParseFiles("web/index.html")
	// tmpl.Execute()
}

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