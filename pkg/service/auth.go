package service

import (
	"crypto/sha1"
	"fmt"

	//todo "jsonapp"
	"jsonapp/pkg/repository"

	"github.com/zhashkevych/todo-app"
)

const salt = "tiurjkfgn675849dkjfgh"

type AuthServise struct {
	repo repository.Authorization
}

func NewAuthServise(repo repository.Repository) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthServise) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
