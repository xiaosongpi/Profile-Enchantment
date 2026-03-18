package domain

import "time"

// Constant
type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

// Original Table
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      Role
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Repository
type UserRepository interface {
	GetUserList(loggedInUser string) ([]User, error)
	GetUserDetail(loggedInUser string, userId string) (*User, error)
}

// Usecase
type UserUsecase interface {
	GetUserList(loggedInUser string) ([]User, error)
	GetUserDetail(loggedInUser string, userId string) (*User, error)
}
