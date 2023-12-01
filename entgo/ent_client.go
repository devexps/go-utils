package entgo

import (
	"errors"
	"fmt"
	"github.com/XSAM/otelsql"
	"time"

	"entgo.io/ent/dialect/sql"
)

type entClientInterface interface {
	Close() error
}

type EntClient[T entClientInterface] struct {
	db  T
	drv *sql.Driver
}

func NewEntClient[T entClientInterface](db T, drv *sql.Driver) *EntClient[T] {
	return &EntClient[T]{
		db:  db,
		drv: drv,
	}
}

func (c *EntClient[T]) Client() T {
	return c.db
}

func (c *EntClient[T]) Driver() *sql.Driver {
	return c.drv
}

func (c *EntClient[T]) Close() {
	_ = c.db.Close()
}

// CreateDriver creates a database driver
func CreateDriver(driverName, dsn string, maxIdleConnections, maxOpenConnections int, connMaxLifetime time.Duration) (*sql.Driver, error) {
	otelDN, err := otelsql.Register(driverName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed registering otelsql: %v", err))
	}
	drv, err := sql.Open(otelDN, dsn)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed opening connection to db: %v", err))
	}
	db := drv.DB()
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connMaxLifetime)

	return drv, nil
}
