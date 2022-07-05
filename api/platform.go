package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
)

type Platform struct {
}

func (p *Platform) Handler(c *fiber.Ctx) error {
	result := db.Client.Platform.Query().AllX(c.Context())
	return c.JSON(result)
}
