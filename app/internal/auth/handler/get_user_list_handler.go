package handler

import (
	"profile-enchantment/app/internal/auth/domain"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *userHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (h userHandler) GetUserListHandler(c *fiber.Ctx) error {
	loggedInUserId := ""

	users, err := h.userUsecase.GetUserListUsecase(loggedInUserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "Sucess to get user list",
		Data:    users,
	})
}
