package main

import (
	"fmt"
	"github.com/long2ice/awesome/conf"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := CreateApp()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", conf.ServerConfig.Host, conf.ServerConfig.Port)))
}
