package usecase

import "profile-enchantment/app/internal/domain"

func (u *userUsecase) GetUserList(loggedInUserID int) ([]domain.User, error) {
	return u.userRepository.GetUserList(loggedInUserID)
}
