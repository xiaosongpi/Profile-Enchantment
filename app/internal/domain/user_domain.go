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

type UserStats struct {
	TotalUser          int64
	TotalLoggeedInUser int64
	TotalActiveUser    int64
}

type CreateUserInput struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token     string
	ExpiredAt time.Time
	User      *User
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindStats() (*UserStats, error)
	FindByID(userId int) (*User, error)
	FindByEmail(email string) *User
	Create(user *User) (*User, error)
	Delete(userIds []int) error
}

type UserUsecase interface {
	GetUserList(loggedInUserId int) ([]User, *UserStats, error)
	GetUserDetail(loggedInUserId, userId int) (*User, error)
	CreateUser(reqBody *CreateUserInput) (*User, error)
	DeleteUser(loggedInUserId int, userIds []int) error
	Login(reqBody *LoginInput) (*LoginResult, error)
}
