package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/tasks"
	log "github.com/sirupsen/logrus"
)

type TopicSearch struct {
}

func (t *TopicSearch) Handler(c *fiber.Ctx) error {
	info, err := tasks.Client.Enqueue(asynq.NewTask(tasks.TypeGetTopicsPeriodic, []byte("")))
	if err != nil {
		return err
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
	return err
}
