package usecase

import "profile-enchantment/app/internal/domain"

func (u *userUsecase) GetUserList(loggedInUserID string) ([]domain.User, error) {
	return u.userRepository.GetUserList(loggedInUserID)
}
