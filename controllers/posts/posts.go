package posts

import (
	"fmt"

	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/internals/notifications"
	"github.com/Brian-M-J/social-media-app-go/internals/validator"
	"github.com/Brian-M-J/social-media-app-go/models/posts"
	"github.com/Brian-M-J/social-media-app-go/services/users"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var post dto.PostCreate
	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}

	if err := validator.Payload(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}
	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userId
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}
	ps := posts.New()
	ps.Post = &dto.Post{}
	ps.Post.UserID = userId
	ps.Post.Content = post.Content
	ps.Create(ctx)

	msg := fmt.Sprintf("Hello, your friend %v has created a new post", us.User.Name)
	notifications.NotifyUsers(ctx, userId, msg)

	return c.Status(fiber.StatusCreated).JSON(ps.Post)
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
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}
	ps := posts.New()
	ps.Posts = &dto.Posts{}
	ps.UserID = userId
	ps.GetAll(ctx)
	return c.Status(fiber.StatusOK).JSON(ps.Posts)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}

	postID, err := uuid.Parse(c.Params("post_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userID
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}

	ps := posts.New()
	ps.Posts = &dto.Posts{}

	ps.UserID = userID
	ps.ID = postID
	ps.Delete(ctx)
	return c.SendStatus(fiber.StatusNoContent)
}
