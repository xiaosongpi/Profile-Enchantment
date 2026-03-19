package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/domain"
)

func (u *userUsecase) CreateUser(reqBody domain.CreateUserInput) (*domain.User, error) {
	userEmail := u.userRepository.GetUserByEmail(reqBody.Email)
	if userEmail != nil {
		return nil, fmt.Errorf("email already taken")
	}
	return u.userRepository.CreateUser(reqBody)
}
