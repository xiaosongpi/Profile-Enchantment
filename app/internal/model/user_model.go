package model

import "time"

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
