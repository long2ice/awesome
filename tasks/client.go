package tasks

import (
	"github.com/go-resty/resty/v2"
	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/config"
)

var Client *asynq.Client
var Resty = resty.New()
var redis = config.RedisConfig
var Option = asynq.RedisClientOpt{Addr: redis.Addr, Password: redis.Password, DB: redis.DB}

func init() {
	Client = asynq.NewClient(Option)
}
