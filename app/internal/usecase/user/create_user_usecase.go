package usecase

import (
	"fmt"
	"profile-enchantment/app/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) CreateUser(reqBody *domain.CreateUserInput) (*domain.User, error) {
	userEmail := u.userRepository.FindByEmail(reqBody.Email)
	if userEmail != nil {
		return nil, fmt.Errorf("email already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Email:     reqBody.Email,
		Password:  string(hashedPassword),
	}

	createdUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
