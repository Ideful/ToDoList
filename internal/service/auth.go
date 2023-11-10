package service

import (
	model "ToDoList/internal/models"
	"ToDoList/internal/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// user and password creation

const (
	salt      = "h2uhto9qogibgaoibg21"
	signInKey = "lrbrqwi@213"
	TokenTL   = 12 * time.Hour
)

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

// user validation part

type tokenClaims struct {
	jwt.StandardClaims
	Id       int    `json:"-" db:"id"`
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username, s.generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, 
		user.Id,
	})
	return token.SignedString([]byte(signInKey))
}
