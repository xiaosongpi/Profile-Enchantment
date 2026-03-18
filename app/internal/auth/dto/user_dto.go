package dto

import "time"

type GetUserListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetUserListResponse struct {
	FullName  string     `json:"fullName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	IsActive  bool       `json:"isActived"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type GetUserDetailRequest struct {
	ID string `json:"id"`
}

type GetUserDetailResponse struct {
	ID        string     `json:"id"`
	FirstName string     `json:"firstNam"`
	LastName  string     `json:"lastName"`
	FullName  string     `json:"fullName"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	IsActive  bool       `json:"isActived"`
	IsDeleted bool       `json:"isDeleted"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
