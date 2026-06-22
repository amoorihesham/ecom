package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	ConnectionString string
	MaxConns         int
	MinConns         int
	PingTimeout      time.Duration
}

func NewDatabase(ctx context.Context, dbCfg *DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbCfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbCfg.MaxConns)
	db.SetMaxIdleConns(dbCfg.MinConns)

	pCtx, cancel := context.WithTimeout(ctx, dbCfg.PingTimeout)
	defer cancel()
	if err := db.PingContext(pCtx); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(ctx context.Context, connectionString string) error {
	m, err := migrate.New("file://internal/shared/database/migrations", connectionString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
