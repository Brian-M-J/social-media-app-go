package routes

import "github.com/gofiber/fiber/v2"

func Posts(r fiber.Router) {
	postsRoutes := r.Group("/posts")
	postsRoutes.Post("/", nil)
	postsRoutes.Get("/", nil)
	postsRoutes.Get("/:id", nil)
	postsRoutes.Delete("/:id", nil)
}
