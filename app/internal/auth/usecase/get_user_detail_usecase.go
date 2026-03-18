package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/auth/dto"
)

func (u *userUsecase) GetUserDetailUsecase(loggedInUserId string, userId string) (*dto.GetUserDetailResponse, error) {
	userRepo, err := u.userRepository.GetUserDetailRepository(loggedInUserId, userId)
	if err != nil {
		return nil, fmt.Errorf("Failed to get user detail")
	}

	return userRepo, nil
}
