package service

import (
	model "ToDoList/internal/models"
	"ToDoList/internal/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "h2uhto9qogibgaoibg21"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generateHashPassword(pwd string) string {
	hash := sha1.New()
	hash.Write([]byte(pwd))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
