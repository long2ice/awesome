package tasks

import (
	"context"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestGetTopicPeriodic(t *testing.T) {
	err := GetTopicPeriodic(context.Background(), nil)
	assert.Nil(t, err)
}
func TestSyncRepo(t *testing.T) {
	err := SyncRepoPeriodic(context.Background(), asynq.NewTask(TypeSyncRepoPeriodic, nil))
	assert.Nil(t, err)
}
func TestSyncTopic(t *testing.T) {
	err := SyncTopicPeriodic(context.Background(), asynq.NewTask(TypeSyncTopicPeriodic, nil))
	assert.Nil(t, err)
}
