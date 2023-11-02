package repository

import (
	// "database/sql"

	model "ToDoList/internal/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int,error)
}

type ToDoList interface {
	
}

type ToDoItem interface {
	
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}