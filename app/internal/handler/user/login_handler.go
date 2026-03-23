package handler

import (
	"fmt"
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/handler/dto"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) Login(c *fiber.Ctx) error {
	var reqBody domain.LoginInput
	err := c.BodyParser(&reqBody)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(pkg.FailedResponse{
			Success: false,
			Message: "invalid input",
		})
	}

	user, err := h.userUsecase.Login(&reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: "failed to login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "login success",
		Data: dto.LoginResponse{
			Token:     user.Token,
			ExpiredAt: user.ExpiredAt,
			User: &dto.GetUserDetailResponse{
				ID:        user.User.ID,
				FirstName: user.User.FirstName,
				LastName:  user.User.LastName,
				Email:     user.User.Email,
				Role:      string(user.User.Role),
				IsActive:  user.User.IsActive,
				CreatedAt: user.User.CreatedAt,
				UpdatedAt: user.User.UpdatedAt,
			},
		},
	})
}
