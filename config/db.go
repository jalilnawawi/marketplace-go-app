package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return db
}
