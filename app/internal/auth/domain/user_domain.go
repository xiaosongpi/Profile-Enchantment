package domain

import (
	"profile-enchantment/app/internal/auth/dto"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Role      string    `gorm:"column:role"`
	IsActive  bool      `gorm:"column:is_active"`
	IsDeleted bool      `gorm:"column:is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

type UserRepository interface {
	GetUserListRepository(loggedInUserId string) ([]dto.GetUserListResponse, error)
	GetUserDetailRepository(loggedInUserId string, userId string) (*dto.GetUserDetailResponse, error)
}

type UserUsecase interface {
	GetUserListUsecase(loggedInUserId string) ([]dto.GetUserListResponse, error)
	GetUserDetailUsecase(loggedInUserId string, userId string) (*dto.GetUserDetailResponse, error)
}
