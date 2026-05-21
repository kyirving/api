package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint64    `gorm:"uniqueIndex:uk_user_id;not null" json:"user_id"`
	Username     string    `gorm:"size:128;index:idx_username;not null;default:''" json:"username"`
	Password     string    `gorm:"size:128;not null;default:''" json:"-"`
	Mobile       string    `gorm:"size:128;index:idx_mobile;default:''" json:"mobile"`
	Nikename     string    `gorm:"size:128;default:''" json:"nikename"`
	RegisterType uint8     `gorm:"default:1" json:"register_type"`
	Balance      string    `gorm:"type:decimal(10,2);not null;default:0.00" json:"balance"`
	IsDeleted    uint8     `gorm:"default:0" json:"is_deleted"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
