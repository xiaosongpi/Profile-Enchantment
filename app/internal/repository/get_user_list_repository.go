package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) GetUserList(loggedInUserId int) ([]domain.User, error) {
	var models []model.UserModel
	result := r.db.
		Table("users").
		Where("deleted_at IS NULL").
		Find(&models)

	if result.Error != nil {
		return nil, result.Error
	}

	users := make([]domain.User, len(models))
	for i, m := range models {
		users[i] = domain.User{
			ID:        m.ID,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Email:     m.Email,
			Password:  m.Password,
			Role:      domain.Role(m.Role),
			IsActive:  m.IsActive,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			DeletedAt: m.DeletedAt,
		}
	}

	return users, nil
}
