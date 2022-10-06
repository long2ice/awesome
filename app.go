package main

import (
	"embed"
	"github.com/long2ice/awesome/config"
	"net/http"

	"github.com/long2ice/awesome/api"

	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hibiken/asynqmon"
	"github.com/long2ice/awesome/error"
	"github.com/long2ice/awesome/tasks"
	"github.com/long2ice/fibers"
	"github.com/long2ice/fibers/router"
)

func initRouters(app *fibers.App) {
	apiGroup := app.Group("/api")

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
	topic := apiGroup.Group("/topic", fibers.Tags("Topic"))
	topic.Get("", topicSearch)
	topic.Get("/:topic_id/repo", repoSearch)
	topic.Get("/:topic_id/subtopics", subtopics)
	apiGroup.Get("/platform", platform)

	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/asynqmon",
		RedisConnOpt: tasks.Option,
	})
	app.All("/asynqmon/*", adaptor.HTTPHandler(h))
}

func initMiddlewares(app *fibers.App) {
	app.Use(
		logger.New(logger.Config{
			TimeFormat: config.ServerConfig.LogTimeFormat,
			TimeZone:   config.ServerConfig.Timezone,
		}),
		recover.New(),
	)

}

//go:embed static/*
var static embed.FS

func CreateApp() *fibers.App {
	app := fibers.New(NewSwagger(), fiber.Config{ErrorHandler: error.Handler})
	app.AfterInit(func() {
		app.Use("/", filesystem.New(filesystem.Config{
			Root:       http.FS(static),
			PathPrefix: "static",
			Browse:     true,
		}))
	})
	initMiddlewares(app)
	initRouters(app)
	return app
}
