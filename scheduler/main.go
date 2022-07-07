package main

import (
	"time"

	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/tasks"
	log "github.com/sirupsen/logrus"
)

func main() {
	loc, err := time.LoadLocation(conf.ServerConfig.Timezone)
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
	entryID, err = scheduler.Register("0 2 * * *", asynq.NewTask(tasks.TypeSyncTopicPeriodic, nil))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q type: %s \n", entryID, tasks.TypeSyncTopicPeriodic)
	entryID, err = scheduler.Register("0 3 * * *", asynq.NewTask(tasks.TypeSyncRepoPeriodic, nil))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q type: %s \n", entryID, tasks.TypeSyncRepoPeriodic)
	if err = scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
