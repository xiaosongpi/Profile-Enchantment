package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) CreateUser(reqBody domain.CreateUserInput) (*domain.User, error) {
	models := model.UserModel{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Email:     reqBody.Email,
		Password:  reqBody.Password,
		Role:      string(domain.RoleUser),
		IsActive:  true,
	}

	result := r.db.
		Table("users").
		Create(&models)

	if result.Error != nil || result.RowsAffected == 0 {
		return nil, result.Error
	}

	user := &domain.User{
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

	return user, nil
}
