package route

import (
	"profile-enchantment/app/internal/auth/domain"
	"profile-enchantment/app/internal/auth/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, userUsecase domain.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)

	userRoute := app.Group("/users")

	userRoute.Get("/", userHandler.GetUserListHandler)
	userRoute.Get("/:id", userHandler.GetUserDetailHandler)
}
