package service

import (
	"errors"
	"go-auth/internal/repository"
	"strconv"
	"time"

	"go-auth/internal/dto"
	"go-auth/internal/model"

	"golang.org/x/crypto/bcrypt"
)

// interface Service
type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) error
}

// struct implementation
type authService struct {
	repo repository.UserRepository
}

// CONSTRUCTOR
func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

// Logika Register
func (s *authService) Register(req dto.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Terjadi Kesalahan Saat Membuat Password")
	}

	// objek User dari Model
	user := &model.User{
		ID:       strconv.FormatInt(time.Now().UnixNano(), 10),
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Panggil Fungsi Create dari layer repository
	return s.repo.Create(user)
}

// Logika Login
func (s *authService) Login(req dto.LoginRequest) error {
	// Cari User berdasarkan Email di database
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return errors.New("Email atau Password Salah")
	}

	// Bandingkan password plaintext dengan hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.New("Email atau Password Salah")
	}
	return nil
}