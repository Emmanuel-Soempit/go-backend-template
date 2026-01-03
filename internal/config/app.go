package config

import (
	"context"
	"go-backend-template/ent"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
)

func appConfigurations(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://whatsapp-vendor-frontend.vercel.app",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	log.Println("CORS middleware applied")

	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request Origin: %s", c.Get("Origin"))
		return c.Next()
	})

	log.Println("App configurations successful")
}

func databseConfigs() *ent.Client {
	// Initialize database connection
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed opening connection to postgres: %v", err)
	}
	// defer client.Close()
	log.Println("Database Connected")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Migration Successful")

	return client
}
