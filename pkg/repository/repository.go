package repository

import (
	todo "jsonapp"

	"github.com/jmoiron/sqlx"
	//"github.com/zhashkevych/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
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
	return &Repository{}
	//AuthPostgres:NewAuthPostgres(repos.Authorization),

}

//
//"github.com/zhashkevych/todo-app"
