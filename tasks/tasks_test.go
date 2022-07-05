package tasks

import (
	"context"
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestGetRepos(t *testing.T) {
	err := GetRepos(context.Background(), asynq.NewTask(TypeGetRepos, []byte("1")))
	assert.Nil(t, err)
}
