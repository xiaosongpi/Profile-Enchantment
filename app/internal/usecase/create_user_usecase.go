package usecase

import "profile-enchantment/app/internal/domain"

func (u *userUsecase) CreateUser(reqBody domain.CreateUserInput) (*domain.User, error) {
	return u.userRepository.CreateUser(reqBody)
}
