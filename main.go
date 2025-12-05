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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New(fiber.Config{
		AppName: "Backend-Go",
	})

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	service := user.NewUserService()
	handler := user.NewUserHandler(service)
	api := app.Group("/backend-go/v1")
	userGroup := api.Group("/user")

	userGroup.Get("", handler.GetAllUser)
	userGroup.Get("/:id", handler.GetUserById)
	userGroup.Post("", handler.CreateUser)
	userGroup.Put("/:id", handler.UpdateUser)
	userGroup.Delete("/:id", handler.DeleteUser)

	log.Fatal(app.Listen(":" + port))
}
