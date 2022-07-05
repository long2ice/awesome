package main

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hibiken/asynqmon"
	"github.com/long2ice/awesome/api"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/error"
	"github.com/long2ice/awesome/tasks"
	"github.com/long2ice/fibers"
	"github.com/long2ice/fibers/router"
)

func initRouters(app *fibers.App) {
	var (
		topicSearch = router.New(&api.TopicSearch{}, router.Summary("Search topic"))
		repoSearch  = router.New(&api.Repo{})
	)
	topic := app.Group("/topic", fibers.Tags("Topic"))
	topic.Get("", topicSearch)
	project := topic.Group("/:topic_id/repo")
	project.Get("", repoSearch)
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/asynqmon",
		RedisConnOpt: tasks.Option,
	})
	app.App.All("/asynqmon/*", adaptor.HTTPHandler(h))
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
