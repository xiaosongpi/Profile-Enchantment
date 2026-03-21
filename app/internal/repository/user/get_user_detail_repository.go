package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) GetUserDetail(loggedInUserId int, userId int) (*domain.User, error) {
	var model model.UserModel
	result := r.db.
		Table("users").
		Where("id = ?", userId).
		Where("deleted_at IS NULL").
		First(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	user := domain.User{
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

	return &user, nil
}
