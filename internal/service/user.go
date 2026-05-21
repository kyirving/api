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

type UserService struct {
	repo      *repository.UserRepo
	jwtSecret []byte
}

func NewUserService(repo *repository.UserRepo, cfg *config.JWTConfig) *UserService {
	return &UserService{repo: repo, jwtSecret: []byte(cfg.Secret)}
}

func (s *UserService) Register(username, password, nikename string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	maxID, err := s.repo.MaxUserID()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		UserID:       maxID + 1,
		Username:     username,
		Password:     string(hashed),
		Nikename:     nikename,
		RegisterType: 1,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	if !verifyPassword(user.Password, password) {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(s.jwtSecret)
}

func verifyPassword(hashed, plain string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil {
		return true
	}
	return hashed == plain
}
