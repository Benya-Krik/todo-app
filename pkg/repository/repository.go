package repository

import (
	"eduProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user eduProject.User) (int, error)
	GetUser(username, password string) (eduProject.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
