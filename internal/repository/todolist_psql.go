package repository

import (
	model "ToDoList/internal/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodolistPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodolistPostgres {
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

func (r *TodolistPostgres) Update(userId, listId int, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)
	// fmt.Println(args)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodolistPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
