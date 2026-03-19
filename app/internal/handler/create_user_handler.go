package handler

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/handler/dto"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) CreateUserHandler(c *fiber.Ctx) error {
	var reqBody domain.CreateUserInput
	err := c.BodyParser(&reqBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.FailedResponse{
			Success: false,
			Message: "invalid input",
		})
	}

	user, err := h.userUsecase.CreateUser(reqBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	response := dto.CreateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      string(user.Role),
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "success to create user",
		Data:    response,
	})
}
