package usecase

import (
	"profile-enchantment/app/internal/domain"
)

func (u *userUsecase) GetUserList(loggedInUserID int) ([]domain.User, *domain.UserStats, error) {
	userStats, err := u.userRepository.FindStats()
	if err != nil {
		return nil, nil, err
	}

	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, nil, err
	}

	return users, userStats, nil
}
