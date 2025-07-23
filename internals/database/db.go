package database

import (
	"fmt"

	"gorm.io/driver/postgres"

	"github.com/Brian-M-J/social-media-app-go/models/friendship"
	"github.com/Brian-M-J/social-media-app-go/models/posts"
	"github.com/Brian-M-J/social-media-app-go/models/users"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Config() {
	dsn := "user=brian database=postgres sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Unable to open database, error: %v\n", err)
		panic(err)
	}

	sql, err := db.DB()

	if err != nil {
		fmt.Printf("Unable to get sql database from gorm, error: %v\n", err)
		panic(err)
	}

	if err := sql.Ping(); err != nil {
		fmt.Printf("Unable to connect to the database, error: %v\n", err)
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&users.Users{}, &friendship.Friendships{}, &posts.Posts{})

}
