package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/usmonzodasomon/e-wallet/models"
	"github.com/usmonzodasomon/e-wallet/pkg/repository"
)

const (
	tokenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos}
}

func (s *AuthService) SignUp(user *models.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.SignUp(user)
}

func (s *AuthService) GenerateToken(user models.User) (string, error) {
	user, err := s.repos.GetUser(user.Phone, generatePasswordHash(user.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.ID,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
