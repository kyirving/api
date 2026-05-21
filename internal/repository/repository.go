package repository

import (
	"errors"

	"api/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Ping() error {
	var v int
	if err := r.db.Raw("SELECT 1").Scan(&v).Error; err != nil {
		return err
	}
	if v != 1 {
		return errors.New("database ping returned unexpected value")
	}
	return nil
}

func (r *Repository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
