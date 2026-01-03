package config

import (
	"log"
	"os"

	"go-backend-template/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func InitApp() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to OS env")
	}
	log.Println(".env loaded successfully")
	app := fiber.New()

	appConfigurations(app)

	client := databseConfigs()
	api.InitializeApiRoutes(app, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
