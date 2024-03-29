package config

import (
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	Listen        string `yaml:"listen"`
	Timezone      string `yaml:"timezone"`
	LogTimeFormat string `yaml:"logTimeFormat"`
}
type Database struct {
	Dsn string `yaml:"dsn"`
}
type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
type MeiliSearch struct {
	Server    string `yaml:"server"`
	MasterKey string `yaml:"masterKey"`
	BatchSize int    `yaml:"batchSize"`
}
type Config struct {
	Server      *Server      `yaml:"server"`
	Database    *Database    `yaml:"database"`
	MeiliSearch *MeiliSearch `yaml:"meilisearch"`
	Redis       *Redis       `yaml:"redis"`
	Sentry      *Sentry      `yaml:"sentry"`
}
type Sentry struct {
	Dsn         string `yaml:"dsn"`
	Environment string `yaml:"environment"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v \n", err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	ServerConfig = c.Server
	DatabaseConfig = c.Database
	MeiliSearchConfig = c.MeiliSearch
	RedisConfig = c.Redis
	SentryConfig = c.Sentry
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              SentryConfig.Dsn,
		Environment:      SentryConfig.Environment,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

var (
	ServerConfig      *Server
	DatabaseConfig    *Database
	MeiliSearchConfig *MeiliSearch
	RedisConfig       *Redis
	SentryConfig      *Sentry
)
