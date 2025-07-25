package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Client() *gorm.DB {
	return DB
}

func Connect() {
	dsn := "user=brian database=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to open database, error: %v\n", err)
	}

	sql, err := db.DB()

	if err != nil {
		log.Fatalf("Unable to get sql database from gorm, error: %v\n", err)
	}

	if err := sql.Ping(); err != nil {
		log.Fatalf("Unable to connect to the database, error: %v\n", err)
	}

	fmt.Printf("Successfully connected to the Postgres database")
	DB = db
}
