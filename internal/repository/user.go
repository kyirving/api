package repository

import (
	"api/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ? AND is_deleted = 0", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) MaxUserID() (uint64, error) {
	var maxID uint64
	err := r.db.Model(&model.User{}).Select("COALESCE(MAX(user_id), 0)").Scan(&maxID).Error
	return maxID, err
}
