package tasks

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopicsPeriodic(t *testing.T) {
	err := GetTopicsPeriodic(context.Background(), nil)
	assert.Nil(t, err)
}
