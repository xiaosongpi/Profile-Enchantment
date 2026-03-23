package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	models := model.UserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      domain.RoleUser,
		IsActive:  true,
	}

	result := r.db.Table("users").Create(&models)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	user.ID = models.ID
	user.CreatedAt = models.CreatedAt
	user.UpdatedAt = models.UpdatedAt

	return user, nil
}
