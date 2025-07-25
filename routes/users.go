package routes

import (
	"github.com/Brian-M-J/social-media-app-go/controllers/users"
	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router) {
	userRoutes := r.Group("/users")
	userRoutes.Post("/", users.Add)
	userRoutes.Get("/:id", users.Get)
	userRoutes.Get("/", users.GetAll)
	userRoutes.Delete("/:id", users.Delete)
}
