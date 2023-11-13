package main

import (
	server "ToDoList/internal"
	handler "ToDoList/internal/handler"
	"ToDoList/internal/repository"
	"ToDoList/internal/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
)

func main() {
	var cfg repository.Config
	if _, err := toml.DecodeFile(repository.DBpath, &cfg); err != nil {
		log.Println("Error decoding config file:", err)
		return
	}

	db, err := repository.CreatePostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(server.Server)
	go func() {
		if err := s.Run("8080", handlers.InitRoutes()); err != nil {
			log.Printf("error: \t %s", err)
		}
	}()
	log.Println("App has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("App is shutting down")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("an error while shutting down server:%s\n", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("an error while shutting down DB:%s\n", err)
	}
}
