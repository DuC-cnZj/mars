package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/duc-cnzj/mars/v4/internal/ent"
	"github.com/duc-cnzj/mars/v4/internal/mlog"
	"github.com/duc-cnzj/mars/v4/internal/util/timer"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() (*ent.Client, error) {
	client, _ := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1&loc=Local")
	_ = client.Schema.Create(context.TODO())

	return client, nil
}

type slowLogDriver struct {
	dialect.Driver
	slowThreshold time.Duration
	logger        mlog.Logger
	timer         timer.Timer
}

func (d *slowLogDriver) Exec(ctx context.Context, query string, args, v any) error {
	start := d.timer.Now()
	err := d.Driver.Exec(ctx, query, args, v)
	elapsed := time.Since(start)
	if elapsed > d.slowThreshold {
		d.logger.Infof("slow query: %s, args: %v, took: %s", query, args, elapsed)
	}
	return err
}

func (d *slowLogDriver) Query(ctx context.Context, query string, args, v any) error {
	start := d.timer.Now()
	err := d.Driver.Query(ctx, query, args, v)
	elapsed := time.Since(start)
	if elapsed > d.slowThreshold {
		d.logger.Infof("slow query: %s, args: %v, took: %s", query, args, elapsed)
	}
	return err
}

func InitMysqlDB(dsn string, logger mlog.Logger, slogLogEnabled bool, slowLogThreshold time.Duration, timer timer.Timer) (*ent.Client, error) {
	var (
		drv dialect.Driver
		err error
	)
	drv, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.(*sql.Driver).DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	if slogLogEnabled {
		drv = &slowLogDriver{
			timer:         timer,
			Driver:        drv,
			slowThreshold: slowLogThreshold,
			logger:        logger.WithModule("SlowLog"),
		}
	}
	dbCli := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			logger.Debug(a...)
		}),
	)

	return dbCli, nil
}
