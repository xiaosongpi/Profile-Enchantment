package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/auth/dto"
)

func (u *userUsecase) GetUserListUsecase(loggedInUserId string) ([]dto.GetUserListResponse, error) {
	userRepo, err := u.userRepository.GetUserListRepository(loggedInUserId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user list")
	}

	return userRepo, nil
}
