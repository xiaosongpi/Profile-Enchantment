package handler

import (
	"profile-enchantment/app/pkg"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) GetUserDetailHandler(c *fiber.Ctx) error {
	loggedInUserId := ""
	userId := c.Params("id")

	users, err := h.userUsecase.GetUserDetailUsecase(loggedInUserId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(pkg.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Success: true,
		Message: "Sucess to get user detail",
		Data:    users,
	})
}
