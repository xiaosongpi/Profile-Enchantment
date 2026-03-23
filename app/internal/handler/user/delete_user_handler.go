package handler

import (
	"profile-enchantment/app/internal/handler/dto"
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	loggedInUserId := 1

	var reqBody dto.DeleteUserRequest
	err := c.BodyParser(&reqBody)
	if err != nil || len(reqBody.ID) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.FailedResponse{
			Success: false,
			Message: "invalid input",
		})
	}

	err = h.userUsecase.DeleteUser(loggedInUserId, reqBody.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "success to delete user",
		Data:    nil,
	})
}
