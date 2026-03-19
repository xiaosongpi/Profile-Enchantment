package handler

import (
	"profile-enchantment/app/internal/handler/dto"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) GetUserList(c *fiber.Ctx) error {
	loggedInUserId := 1

	users, err := h.userUsecase.GetUserList(loggedInUserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	response := make([]dto.GetUserListResponse, len(users))
	for i, user := range users {
		response[i] = dto.GetUserListResponse{
			ID:        user.ID,
			FullName:  user.FirstName + " " + user.LastName,
			Email:     user.Email,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt,
		}
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "success to get user detail",
		Data:    response,
	})
}
