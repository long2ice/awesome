package main

import (
	"github.com/long2ice/awesome/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := CreateApp()
	log.Fatal(app.Listen(config.ServerConfig.Listen))
}
