package main

import (
	"log"
	"os"
	"profile-enchantment/app/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Failed to connect .env: %s", err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect database: %s", err)
	}

	_ = db

	log.Println("Success to connect database")

	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Printf("Server run at http://localhost:%s", port)
	app.Listen(":" + port)
}
