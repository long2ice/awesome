package main

import (
	"github.com/long2ice/awesome/conf"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := CreateApp()
	log.Fatal(app.Listen(conf.ServerConfig.Listen))
}
