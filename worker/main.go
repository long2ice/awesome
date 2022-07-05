package main

import (
	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/tasks"
	log "github.com/sirupsen/logrus"
)

func main() {
	srv := asynq.NewServer(
		tasks.Option,
		asynq.Config{Concurrency: 1},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeGetTopicsPeriodic, tasks.GetTopicsPeriodic)
	mux.HandleFunc(tasks.TypeGetRepos, tasks.GetRepos)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
