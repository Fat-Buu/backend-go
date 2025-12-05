package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAllUser godoc
// @Summary List users
// @Tags users
// @Produce json
// @Success 200 {array} UserResponse
// @Router /user [get]
func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	users := h.service.getAllUser()
	return c.JSON(fiber.Map{
		"data": users,
	})
}

// GetUserById godoc
// @Summary Get a user by UUID
// @Tags users
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [get]
func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}
	user, found := h.service.GetUserByID(id)
	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{"data": user})
}

// CreateUser godoc
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequest true "User data"
// @Success 201 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Param user body UserRequest true "User data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": users,
	})
}
