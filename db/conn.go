package db

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func New(uri string, maxopen, maxidle, idleTime int) (*Database, error) {
	url, err := pq.ParseURL(uri)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxopen)
	db.SetMaxIdleConns(maxidle)
	db.SetConnMaxIdleTime(time.Duration(idleTime) * time.Second)

	return &Database{db}, nil
}

func (db *Database) PingContext(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}

//go:embed "schema.sql"
var schema string

func (db *Database) MigrateDb() error {
	if _, err := db.DB.Exec(schema); err != nil {
		return err
	}

	return nil
}
