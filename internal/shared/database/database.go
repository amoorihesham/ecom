package database

import (
	"context"
	"database/sql"
	"time"

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
	defer db.Close()

	db.SetMaxOpenConns(dbCfg.MaxConns)
	db.SetMaxIdleConns(dbCfg.MinConns)

	pCtx, cancel := context.WithTimeout(ctx, dbCfg.PingTimeout)
	defer cancel()
	if err := db.PingContext(pCtx); err != nil {
		return nil, err
	}

	return db, nil
}
