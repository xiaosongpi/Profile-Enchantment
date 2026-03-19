package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) GetUserDetail(loggedInUserId string, userId string) (*domain.User, error) {
	var models model.UserModel
	err := r.db.
		Table("users").
		Where("id = ?", userId).
		Where("deleted_at IS NULL").
		First(&models).Error

	if err != nil {
		return nil, err
	}

	user := domain.User{
		ID:        models.ID,
		FirstName: models.FirstName,
		LastName:  models.LastName,
		Email:     models.Email,
		Password:  models.Password,
		Role:      domain.Role(models.Role),
		IsActive:  models.IsActive,
		CreatedAt: models.CreatedAt,
		UpdatedAt: models.UpdatedAt,
		DeletedAt: models.DeletedAt,
	}

	return &user, nil
}
