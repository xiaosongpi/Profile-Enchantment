package handler

import "profile-enchantment/app/internal/auth/domain"

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *userHandler {
	return &userHandler{userUsecase: userUsecase}
}
