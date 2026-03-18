package repository

import (
	"profile-enchantment/app/internal/auth/domain"
	"profile-enchantment/app/internal/auth/dto"
)

func (r *userRepository) GetUserListRepository(loggedInUserId string) ([]dto.GetUserListResponse, error) {
	var users []domain.User
	err := r.db.
		Table("users").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	var result []dto.GetUserListResponse
	for _, user := range users {
		fullName := user.FirstName + " " + user.LastName

		result = append(result, dto.GetUserListResponse{
			FullName:  fullName,
			Email:     user.Email,
			Role:      user.Role,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		})
	}

	return result, nil
}
