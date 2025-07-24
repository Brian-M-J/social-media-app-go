package users

import (
	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/internals/validator"
	"github.com/Brian-M-J/social-media-app-go/services/users"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var user dto.UserCreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}

	if err := validator.Users(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}
	us := users.New()
	us.User = &dto.User{}
	us.User.Name = user.Name
	us.User.Email = user.Email
	us.User.Password = user.Password
	us.Create(ctx)
	return c.Status(fiber.StatusCreated).JSON(us.User)
}

func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userId

	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Server Error")
	}
	return c.Status(fiber.StatusOK).JSON(us.User)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}
	us := users.New()
	us.User.ID = userID
	if err := us.Delete(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Server Error")
	}
	return c.SendStatus(fiber.StatusNoContent)
}
