package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/domain"
)

func (u *userUsecase) CreateUser(reqBody *domain.CreateUserInput) (*domain.User, error) {
	userEmail := u.userRepository.FindByEmail(reqBody.Email)
	if userEmail != nil {
		return nil, fmt.Errorf("email already taken")
	}

	user := &domain.User{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Email:     reqBody.Email,
		Password:  reqBody.Password,
	}

	createdUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
