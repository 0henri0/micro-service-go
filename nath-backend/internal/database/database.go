package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func ConnectOrm(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("failed to connect database")
	}
	dbConfig, err := db.DB()
	dbConfig.SetConnMaxLifetime(time.Minute * 3)
	dbConfig.SetMaxOpenConns(125)
	dbConfig.SetMaxIdleConns(125)

	return db, nil
}
