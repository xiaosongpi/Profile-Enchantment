package usecase

import "profile-enchantment/app/internal/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}
