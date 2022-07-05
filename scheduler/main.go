package main

import (
	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/tasks"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	scheduler := asynq.NewScheduler(
		tasks.Option,
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)
	entryID, err := scheduler.Register("0 0 * * *", asynq.NewTask(tasks.TypeGetTopicsPeriodic, nil))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q type: %s \n", entryID, tasks.TypeGetTopicsPeriodic)
	if err = scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
