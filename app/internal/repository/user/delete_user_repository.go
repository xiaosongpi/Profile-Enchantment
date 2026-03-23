package repository

import (
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) Delete(userIds []int) error {
	var models model.UserModel
	result := r.db.
		Table("users").
		Where("id IN ?", userIds).
		Delete(&models)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
