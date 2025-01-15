package handlers

import (
	"github.com/feynmaz/fiberg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (h *Handler) UserList(c *fiber.Ctx) error {
	users := h.db.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

func (h *Handler) UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		Name: utils.CopyString(c.FormValue("name")),
	}
	h.db.Insert(user)


	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}
