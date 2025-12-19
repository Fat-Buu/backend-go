package user

import (
	"log"

	"github.com/backend-go/internal/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	userRepository, err := NewUserRepositoryFromFile(config.UsersFilePath())
	if err != nil {
		log.Fatal(err)
	}
	service := NewUserService(userRepository)
	handler := NewUserHandler(service)

	api := app.Group("/backend-go/v1")
	userGroup := api.Group("/user")

	userGroup.Get("", handler.GetAllUser)
	userGroup.Get("/:id", handler.GetUserById)
	userGroup.Post("", handler.CreateUser)
	userGroup.Put("/:id", handler.UpdateUser)
	userGroup.Delete("/:id", handler.DeleteUser)
}
