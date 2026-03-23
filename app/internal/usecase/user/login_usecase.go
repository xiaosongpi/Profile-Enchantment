package usecase

import (
	"fmt"
	"os"
	"profile-enchantment/app/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Login(reqBody *domain.LoginInput) (*domain.LoginResult, error) {
	userEmail := u.userRepository.FindByEmail(reqBody.Email)
	if userEmail == nil {
		return nil, fmt.Errorf("email not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userEmail.Password), []byte(reqBody.Password))
	if err != nil {
		return nil, err
	}

	expiredAt := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":    userEmail.ID,
		"email": userEmail.Email,
		"role":  userEmail.Role,
		"exp":   expiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokeString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	loginResult := &domain.LoginResult{
		Token:     tokeString,
		ExpiredAt: expiredAt,
		User:      userEmail,
	}

	return loginResult, nil
}
