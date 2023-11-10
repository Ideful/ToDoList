package service

import (
	model "ToDoList/internal/models"
	"ToDoList/internal/repository"
	"crypto/sha1"
	"errors"
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
	UserId int `json:"user_id"`
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

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid Sign-in method")
		}
		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims,ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token Claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
