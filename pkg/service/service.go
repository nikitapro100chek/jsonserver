package service

import (
	todo "github.com/nikitapro100chek/jsonserver"
	"github.com/nikitapro100chek/jsonserver/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		//	TodoList:      NewTodoListService(repos.TodoList),
		//	TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),}
	}
}
