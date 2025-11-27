package main

import (
	"log"
	"os"

	_ "github.com/backend-go/docs"
	"github.com/backend-go/user"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Fiber User API
// @version 1.0
// @description This is a sample Fiber API.
// @host localhost:3000
// @BasePath /backend-go/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load env")
	}

	app := fiber.New(fiber.Config{
		AppName: "Backend-Go",
	})
	port := os.Getenv("PORT")

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	api := app.Group("/backend-go/v1")

	userGroup := api.Group("/user")

	userGroup.Get("", user.GetAllUser)
	userGroup.Get("/:id", user.GetUserById)
	userGroup.Post("", user.CreateUser)
	userGroup.Put("/:id", user.UpdateUser)
	userGroup.Delete("/:id", user.DeleteUser)

	log.Fatal(app.Listen(":" + port))
}
