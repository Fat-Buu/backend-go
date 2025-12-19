package main

import (
	"log"
	"os"

	_ "github.com/backend-go/docs"
	"github.com/backend-go/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Fiber User API
// @version 1.0
// @description Fiber CRUD API example
// @host localhost:3000
// @BasePath /backend-go/v1
func main() {
	loadEnv()
	user.LoadUsers()

	port := getPort()
	app := setupFiberApp()

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	user.RegisterRoutes(app)

	log.Fatal(app.Listen(":" + port))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found.")
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func setupFiberApp() *fiber.App {
	return fiber.New(fiber.Config{
		AppName: "Backend-Go",
	})
}
