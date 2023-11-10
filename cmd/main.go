package main

import (
	// "fmt"
	// "html/template"
	// "log"
	server "ToDoList/internal"
	// "fmt"
	handler "ToDoList/internal/handler"
	"ToDoList/internal/repository"
	"ToDoList/internal/service"
	"log"
	// "net/http"
	// "github.com/Ideful/todolist/internal/mi"
)

func main() {
	db, err := repository.CreatePostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	// h := &handler.Handler{}
	// http.HandleFunc("/submit", h.HandleSubmit)
	s := new(server.Server)
	if err := s.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: \t %s", err)
	}
}
