package main

import (
	// "fmt"
	// "html/template"
	// "log"
	server "ToDoList/internal"
	// "fmt"
	"log"
	"net/http"
	handler "ToDoList/internal/handler"
	// "github.com/Ideful/todolist/internal/mi"
)


func main() {
	s:= new(server.Server)
	h := &handler.Handler{}
	http.HandleFunc("/submit", h.HandleSubmit) 
	if err := s.Run("8080"); err != nil {
		log.Fatalf("error: \t %s",err)
	}
}


