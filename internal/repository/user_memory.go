package repository

import (
	"crud-app/internal/domain"
	"errors"
	"sync"
	"time"
)

type UserRepoInterface interface {
	Create(user *domain.User) error
	GetByUsername(username string) (*domain.User, error)
}

type UserMemoryRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *UserMemoryRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.users[user.Username]; ok {
		return errors.New("user already exists")
	}
	user.CreatedAt = time.Now().UTC()
	r.users[user.Username] = user
	return nil
}

func (r *UserMemoryRepository) GetByUsername(username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
