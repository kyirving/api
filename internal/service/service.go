package service

import (
	"errors"
	"time"

	"api/config"
	"api/internal/model"
	"api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repo      *repository.Repository
	jwtSecret []byte
}

func New(repo *repository.Repository, cfg *config.JWTConfig) *Service {
	return &Service{repo: repo, jwtSecret: []byte(cfg.Secret)}
}

func (s *Service) Ping() error {
	return s.repo.Ping()
}

func (s *Service) Register(email, password, name string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashed),
		Name:     name,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
