package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

var users = []User{
	{Id: uuid.New(), Username: "John Go"},
	{Id: uuid.New(), Username: "Jane Fiber"},
}

// @Summary List users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /user [get]
func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

// GetUserByID godoc
// @Summary Get a user by UUID
// @Description Get details of a user by UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {object} User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}
	for _, user := range users {
		if user.Id == id {
			return c.JSON(fiber.Map{
				"data": user,
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
}

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}
