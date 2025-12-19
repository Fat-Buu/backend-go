package app

import (
	"log"
	"os"

	"github.com/backend-go/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupFiberApp() *fiber.App {
	user.LoadUsers()
	app := fiber.New(fiber.Config{
		AppName: "Backend-Go",
	})
	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	user.RegisterRoutes(app)
	return app
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found.")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
