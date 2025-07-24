package config

import (
	"github.com/Brian-M-J/social-media-app-go/internals/database"
	"github.com/Brian-M-J/social-media-app-go/models/friendship"
	"github.com/Brian-M-J/social-media-app-go/models/posts"
	"github.com/Brian-M-J/social-media-app-go/models/users"
)

func Automigration() {
	database.Client().AutoMigrate(&users.Users{}, &friendship.Friendships{}, &posts.Posts{})
}
