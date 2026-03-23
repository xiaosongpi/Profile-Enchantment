package model

import (
	"profile-enchantment/app/internal/domain"
	"time"
)

type UserModel struct {
	ID        int `gorm:"primaryKey;autoIncrement;<-:false"`
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      domain.Role
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type UserModelStats struct {
	TotalUser          int64 `json:"totalIser"`
	TotalLoggeedInUser int64 `json:"totalLoggedInUser"`
	TotalActiveUser    int64 `json:"totalActiveUser"`
}
