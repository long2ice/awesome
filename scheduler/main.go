package main

import (
	"github.com/long2ice/awesome/config"
	"time"

	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/tasks"
	log "github.com/sirupsen/logrus"
)

func main() {
	loc, err := time.LoadLocation(config.ServerConfig.Timezone)
	if err != nil {
		log.Fatal(err)
	}
	scheduler := asynq.NewScheduler(
		tasks.Option,
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)
	entryID, err := scheduler.Register("0 1 * * *", asynq.NewTask(tasks.TypeGetTopicPeriodic, nil))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q type: %s \n", entryID, tasks.TypeGetTopicPeriodic)
	if err = scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
