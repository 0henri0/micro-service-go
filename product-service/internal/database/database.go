package database

import (
	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

func ConnectOrm(dns string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	dbConfig, _ := db.DB()
	dbConfig.SetConnMaxLifetime(time.Minute * 3)
	dbConfig.SetMaxOpenConns(125)
	dbConfig.SetMaxIdleConns(125)

	return db
}
