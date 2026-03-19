package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) GetUserByEmail(email string) *domain.User {
	var model model.UserModel
	result := r.db.
		Table("users").
		Where("email = ?", email).
		First(&model)

	if result.Error != nil {
		return nil
	}

	user := &domain.User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Password:  model.Password,
		Role:      domain.Role(model.Role),
		IsActive:  model.IsActive,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}

	return user
}
