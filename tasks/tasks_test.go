package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRepos(t *testing.T) {
	err := GetRepos(context.Background(), asynq.NewTask(TypeGetRepos, []byte("1")))
	assert.Nil(t, err)
}
