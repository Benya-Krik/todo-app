package service

import (
	"eduProject"
	"eduProject/pkg/repository"
)

type Authorization interface {
	CreateUser(user eduProject.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list eduProject.TodoList) (int, error)
	GetAll(userId int) ([]eduProject.TodoList, error)
	GetById(userId, listId int) (eduProject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input eduProject.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item eduProject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]eduProject.TodoItem, error)
	GetItemById(userId, itemId int) (eduProject.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input eduProject.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
