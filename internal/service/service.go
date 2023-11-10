package service

import (
	"ToDoList/internal/repository"
	"ToDoList/internal/models"
)

type Authorization interface {
	CreateUser(user model.User) (int,error)
	GenerateToken(username,password string) (string,error)
}

type ToDoList interface {
	
}

type ToDoItem interface {
	
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}