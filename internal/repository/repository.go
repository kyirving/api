package repository

import (
	"errors"

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
