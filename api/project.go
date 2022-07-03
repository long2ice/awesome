package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
)

type ProjectSearch struct {
	Keyword string `query:"keyword" validate:"required" example:"mysql"`
	TopicID int    `uri:"topic_id" validate:"required"`
}

func (p *ProjectSearch) Handler(c *fiber.Ctx) error {
	projects := db.Client.Project.Query().AllX(c.Context())
	return c.JSON(projects)
}
