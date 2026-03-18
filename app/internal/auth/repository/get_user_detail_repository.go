package repository

import (
	"profile-enchantment/app/internal/auth/domain"
	"profile-enchantment/app/internal/auth/dto"
)

func (r *userRepository) GetUserDetailRepository(loggedInUserId string, userId string) (*dto.GetUserDetailResponse, error) {
	var user domain.User
	err := r.db.
		Table("users").
		Where("id = ?", userId).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	var result dto.GetUserDetailResponse
	fullName := user.FirstName + " " + user.LastName

	result = dto.GetUserDetailResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  fullName,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		IsDeleted: user.IsDeleted,
		CreatedAt: user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
		DeletedAt: &user.DeletedAt,
	}

	return &result, nil
}
