package repository

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/model"
)

func (r *userRepository) FindStats() (*domain.UserStats, error) {
	var userModel model.UserModelStats

	err := r.db.
		Table("users").
		Model(&model.UserModel{}).
		Where("deleted_at IS NULL").
		Count(&userModel.TotalUser).Error
	if err != nil {
		return nil, err
	}

	err = r.db.
		Table("users").
		Model(&model.UserModel{}).
		Where("deleted_at IS NULL").
		Count(&userModel.TotalLoggeedInUser).Error
	if err != nil {
		return nil, err
	}

	err = r.db.
		Table("users").
		Model(&model.UserModel{}).
		Where("deleted_at IS NULL").
		Count(&userModel.TotalActiveUser).Error
	if err != nil {
		return nil, err
	}

	user := &domain.UserStats{
		TotalUser:          userModel.TotalUser,
		TotalLoggeedInUser: userModel.TotalLoggeedInUser,
		TotalActiveUser:    userModel.TotalActiveUser,
	}

	return user, nil
}
