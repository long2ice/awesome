package db

import (
	"context"
	"fmt"
	"github.com/long2ice/awesome/config"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/long2ice/awesome/ent"
	"github.com/long2ice/awesome/ent/migrate"
	log "github.com/sirupsen/logrus"
)

var Client *ent.Client

func init() {
	var err error
	drv, err := sql.Open("mysql", config.DatabaseConfig.Dsn)
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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
