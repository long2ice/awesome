package tasks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopicPeriodic(t *testing.T) {
	err := GetTopicPeriodic(context.Background(), nil)
	assert.Nil(t, err)
}
