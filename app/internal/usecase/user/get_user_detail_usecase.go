package usecase

import "profile-enchantment/app/internal/domain"

func (u *userUsecase) GetUserDetail(loggedInUserId int, userId int) (*domain.User, error) {
	return u.userRepository.FindByID(userId)
}
