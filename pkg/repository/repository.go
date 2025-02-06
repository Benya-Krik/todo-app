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
	Create(userId int, list eduProject.TodoList) (int, error)
	GetAll(userId int) ([]eduProject.TodoList, error)
	GetById(userId, listId int) (eduProject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input eduProject.UpdateListItem) error
}

type TodoItem interface {
	Create(listId int, item eduProject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]eduProject.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
