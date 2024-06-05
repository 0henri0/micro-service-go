package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type DB struct {
	*sql.DB
}

func Connect(dns string) (*DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Query logs the query and then executes it.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	//start := time.Now()
	rows, err := db.DB.Query(query, args...)
	//elapsed := time.Since(start)
	//logQuery(query, args, elapsed, err)
	return rows, err
}

// Exec logs the query and then executes it.
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := db.DB.Exec(query, args...)
	elapsed := time.Since(start)
	logQuery(query, args, elapsed, err)
	return result, err
}

// logQuery logs the query details.
func logQuery(query string, args []interface{}, elapsed time.Duration, err error) {
	log.Printf("Query: %s\nArgs: %v\nDuration: %s\nError: %v\n", query, args, elapsed, err)
}
