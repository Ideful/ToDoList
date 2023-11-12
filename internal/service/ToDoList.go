package service

import (
	model "ToDoList/internal/models"
	"ToDoList/internal/repository"
)

type TodolistService struct {
	repo repository.ToDoList
}

func NewTodolistService(repo repository.ToDoList) *TodolistService {
	return &TodolistService{repo: repo}
}

func (s *TodolistService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodolistService) GetAll(userId int) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodolistService) GetById(userId, listId int) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
