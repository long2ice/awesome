package main

import (
	"github.com/long2ice/awesome/api"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hibiken/asynqmon"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/error"
	"github.com/long2ice/awesome/tasks"
	"github.com/long2ice/fibers"
	"github.com/long2ice/fibers/router"
)

func initRouters(app *fibers.App) {
	var (
		topicSearch = router.New(api.TopicSearch, router.Summary("Search topic"))
		repoSearch  = router.New(api.Repo, router.Summary("Search repo"))
		subtopics   = router.New(api.RepoSubTopic, router.Summary("Get repo subtopics"))

		platform = router.NewX(
			api.Platform,
			router.Summary("Platform"),
			router.Tags("Platform"),
		)
	)
	topic := app.Group("/topic", fibers.Tags("Topic"))
	topic.Get("", topicSearch)
	topic.Get("/:topic_id/repo", repoSearch)
	topic.Get("/:topic_id/subtopics", subtopics)

	app.Get("/platform", platform)
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/asynqmon",
		RedisConnOpt: tasks.Option,
	})
	app.All("/asynqmon/*", adaptor.HTTPHandler(h))
}

func initMiddlewares(app *fibers.App) {
	app.Use(
		logger.New(logger.Config{
			TimeFormat: conf.ServerConfig.LogTimeFormat,
			TimeZone:   conf.ServerConfig.Timezone,
		}),
		recover.New(),
	)

}

func CreateApp() *fibers.App {
	app := fibers.New(NewSwagger(), fiber.Config{ErrorHandler: error.Handler})
	initMiddlewares(app)
	initRouters(app)
	return app
}
