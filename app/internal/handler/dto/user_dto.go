package dto

import (
	"time"
)

type GetUserListResponse struct {
	TotalUser          int                       `json:"totalIser"`
	TotalLoggeedInUser int                       `json:"totalLoggedInUser"`
	TotalActiveUser    int                       `json:"totalActiveUser"`
	Data               []GetUserListResponseData `json:"data"`
}

type GetUserListResponseData struct {
	ID        int       `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetUserDetailResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateUserResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	Token     string                 `json:"token"`
	ExpiredAt time.Time              `json:"expiredAt"`
	User      *GetUserDetailResponse `json:"user"`
}

type DeleteUserRequest struct {
	ID []int `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
