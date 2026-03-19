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
	GetUserList(loggedInUserId int) ([]User, error)
	GetUserDetail(loggedInUserId int, userId int) (*User, error)
	GetUserByEmail(email string) *User
	CreateUser(reqBody CreateUserInput) (*User, error)
	DeleteUser(loggedInUserId int, userIds []int) error
}

// Usecase
type UserUsecase interface {
	GetUserList(loggedInUserId int) ([]User, error)
	GetUserDetail(loggedInUserId int, userId int) (*User, error)
	CreateUser(reqBody CreateUserInput) (*User, error)
	DeleteUser(loggedInUserId int, userIds []int) error
}
