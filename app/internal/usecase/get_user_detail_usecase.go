package usecase

import "profile-enchantment/app/internal/domain"

func (u *userUsecase) GetUserDetail(loggedInUserId string, userId string) (*domain.User, error) {
	return u.userRepository.GetUserDetail(loggedInUserId, userId)
}
