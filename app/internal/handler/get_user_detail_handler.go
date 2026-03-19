package handler

import (
	"profile-enchantment/app/internal/handler/dto"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) GetUserDetailHandler(c *fiber.Ctx) error {
	loggedInUserId := ""
	userId := c.Params("id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.FailedResponse{
			Success: false,
			Message: "invalid input",
		})
	}

	users, err := h.userUsecase.GetUserDetail(loggedInUserId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	response := &dto.GetUserDetailResponse{
		ID:        users.ID,
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Email:     users.Email,
		Role:      string(users.Role),
		IsActive:  users.IsActive,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "success to get user detail",
		Data:    response,
	})
}
