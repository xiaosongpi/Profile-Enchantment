package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/auth/domain"
	"profile-enchantment/app/internal/auth/dto"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) GetUserListUsecase(loggedInUserId string) ([]dto.GetUserListResponse, error) {
	userRepo, err := u.userRepository.GetUserListRepository(loggedInUserId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user list")
	}

	return userRepo, nil
}
