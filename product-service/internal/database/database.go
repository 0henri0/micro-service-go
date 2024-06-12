package database

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"time"
)

func Connect(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(125)
	db.SetMaxIdleConns(125)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
