package routes

import (
	"github.com/Brian-M-J/social-media-app-go/controllers/friendships"
	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router) {
	friendshipsRoutes := r.Group("/friends")
	friendshipsRoutes.Post("/", friendships.Add)
	friendshipsRoutes.Get("/:id", friendships.Get)
	friendshipsRoutes.Delete("/:id", friendships.Delete)
}
