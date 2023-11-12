package repository

import (
	model "ToDoList/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TodolistPostgres struct {
	db *sqlx.DB
}

func NewTodolistPostgres(db *sqlx.DB) *TodolistPostgres {
	return &TodolistPostgres{
		db: db,
	}
}

func (r *TodolistPostgres) Create(userId int, list model.TodoList) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := transaction.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		transaction.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = transaction.Exec(createUsersListQuery, userId, id)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	return id, transaction.Commit()
}

func (r *TodolistPostgres) GetAll(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)

	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodolistPostgres) GetById(userId, listId int) (model.TodoList, error) {
	var list model.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
						 INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)

	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
