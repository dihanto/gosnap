package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=gosnap sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(20)

	return db
}
