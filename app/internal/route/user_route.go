package route

import (
	"profile-enchantment/app/internal/domain"
	"profile-enchantment/app/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, userUsecase domain.UserUsecase) {
	userHandler := handler.NewUserHandler(userUsecase)

	userRoute := app.Group("/users")

	userRoute.Get("/", userHandler.GetUserList)
	userRoute.Get("/:id", userHandler.GetUserDetail)
	userRoute.Put("/", userHandler.DeleteUser)
	userRoute.Post("/", userHandler.CreateUser)
}
