package repository

import (
	model "ToDoList/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres (db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db :db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int,error) {
	var id int
	query:=fmt.Sprintf("INSERT INTO %s (username,password) VALUES ($1, $2) RETURNING id",usersTable)

	row:=r.db.QueryRow(query,user.Username,user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	// fmt.Println(1212)
	
	return id,nil
}