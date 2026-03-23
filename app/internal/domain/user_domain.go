package domain

import "time"

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

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

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(userId int) (*User, error)
	FindByEmail(email string) *User
	Create(user *User) (*User, error)
	Delete(userIds []int) error
}

type UserUsecase interface {
	GetUserList(loggedInUserId int) ([]User, error)
	GetUserDetail(loggedInUserId, userId int) (*User, error)
	CreateUser(reqBody *CreateUserInput) (*User, error)
	DeleteUser(loggedInUserId int, userIds []int) error
}
