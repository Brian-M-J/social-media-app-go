package routes

import "github.com/gofiber/fiber/v2"

func Friendships(r fiber.Router) {
	friendshipsRoutes := r.Group("/posts")
	friendshipsRoutes.Post("/", nil)
	friendshipsRoutes.Get("/:id", nil)
	friendshipsRoutes.Delete("/", nil)
}
