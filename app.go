package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/long2ice/awesome/api"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/error"
	"github.com/long2ice/fibers"
	"github.com/long2ice/fibers/router"
)

func initRouters(app *fibers.App) {
	var (
		topicSearch   = router.New(&api.TopicSearch{}, router.Summary("Search topic"))
		projectSearch = router.New(&api.ProjectSearch{})
	)
	topic := app.Group("/topic", fibers.Tags("Topic"))
	topic.Get("", topicSearch)
	project := topic.Group("/:topic_id/project")
	project.Get("", projectSearch)
}

func initMiddlewares(app *fibers.App) {
	app.Use(
		logger.New(logger.Config{
			TimeFormat: conf.ServerConfig.LogTimeFormat,
			TimeZone:   conf.ServerConfig.LogTimezone,
		}),
		recover.New(),
		cors.New(),
	)
}

func CreateApp() *fibers.App {
	app := fibers.New(NewSwagger(), fiber.Config{ErrorHandler: error.Handler})
	initMiddlewares(app)
	initRouters(app)
	return app
}
