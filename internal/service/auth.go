package service

import (
	"crud-app/internal/auth"
	"crud-app/internal/domain"
	"crud-app/internal/repository"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users repository.UserRepoInterface
	jwt   *auth.JWTManager
}

func NewAuthService(users repository.UserRepoInterface, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{
		users: users,
		jwt:   jwtManager,
	}
}

func (s *AuthService) Register(username, password, role string) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:       uuid.NewString(),
		Username: username,
		Password: string(hash),
		Role:     role,
	}
	err = s.users.Create(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.users.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}

	generatedToken, err := s.jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}
	return generatedToken, nil
}
