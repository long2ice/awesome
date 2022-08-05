package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
)

func Platform(c *fiber.Ctx) error {
	result := db.Client.Platform.Query().AllX(c.Context())
	return c.JSON(result)
}
