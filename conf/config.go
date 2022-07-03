package conf

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	LogTimezone   string `yaml:"logTimezone"`
	LogTimeFormat string `yaml:"logTimeFormat"`
}
type Database struct {
	Dsn string `yaml:"dsn"`
}
type MeiliSearch struct {
	Server    string `yaml:"server"`
	MasterKey string `yaml:"master_key"`
}
type Config struct {
	Server   *Server
	Database *Database
	Meili    *MeiliSearch
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v \n", err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	ServerConfig = c.Server
	DatabaseConfig = c.Database
	MeiliConfig = c.Meili
}

var (
	ServerConfig   *Server
	DatabaseConfig *Database
	MeiliConfig    *MeiliSearch
)
