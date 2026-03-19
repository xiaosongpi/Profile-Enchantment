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
	ID        int
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

// DTO
type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Repository
type UserRepository interface {
	GetUserList(loggedInUserId string) ([]User, error)
	GetUserDetail(loggedInUserId string, userId string) (*User, error)
	CreateUser(reqBody CreateUserInput) (*User, error)
}

// Usecase
type UserUsecase interface {
	GetUserList(loggedInUserId string) ([]User, error)
	GetUserDetail(loggedInUserId string, userId string) (*User, error)
	CreateUser(reqBody CreateUserInput) (*User, error)
}
