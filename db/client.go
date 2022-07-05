package db

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/ent"
	"github.com/long2ice/awesome/ent/migrate"
	log "github.com/sirupsen/logrus"
)

var Client *ent.Client

func init() {
	var err error
	drv, err := sql.Open("mysql", conf.DatabaseConfig.Dsn)
	if err != nil {
		panic(err)
	}
	db := drv.DB()
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Hour)
	Client = ent.NewClient(ent.Driver(drv))
	if err != nil {
		log.Fatalf("connect to database error: %v", err)
	}
	err = Client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
