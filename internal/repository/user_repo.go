package repository

import (
	"errors"
	"sync"
	"time"
	"go-auth/internal/model"
)

// Interface
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// Struct Implenentation
type userRepository struct {
	mu sync.RWMutex
	db map[string]*model.User
}

// CONSTRUCTOR
func NewUserRepository() UserRepository {
	return &userRepository{
		db: make(map[string]*model.User),
	}
}
	


// implenentation of Create method
func (r *userRepository) Create(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.db[user.Email]; exists {
		return errors.New("Email already exists")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.db[user.Email] = user
	return nil
}

//  implementasi Method FindByEmail
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.db[email]
	if !exists {
		return nil, errors.New("User not found")
	}
	return user, nil
}