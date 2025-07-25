package friendships

import (
	"fmt"

	"github.com/Brian-M-J/social-media-app-go/internals/cache"
	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/internals/validator"
	"github.com/Brian-M-J/social-media-app-go/models/friendship"
	"github.com/Brian-M-J/social-media-app-go/models/users"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var friend dto.FriendsCreate
	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}

	if err := validator.Payload(friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect input body")
	}
	us := users.New()
	us.User = &dto.User{}
	us.User.ID = friend.UserID

	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Server Error")
	}
	us.User = &dto.User{}
	us.User.ID = friend.FriendID

	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found!")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal Server Error")
	}
	fs := friendship.New()
	fs.Friends = &dto.Friends{}
	fs.Friends.UserID = friend.UserID
	fs.Friends.FriendID = friend.FriendID
	fs.Create(ctx)

	if err := cache.Client().Del(ctx, friend.UserID.String()).Err(); err != nil {
		fmt.Println("Error invalidating cache:", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fs.Friends)
}

func Get(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")
	userID, err := uuid.Parse(id)
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
	fs := friendship.New()
	fs.UserID = userID
	fs.GetAll(ctx)
	return c.Status(fiber.StatusOK).JSON(fs.AllFriends)
}

func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Incorrect user id")
	}
	fid := c.Query("f_id")
	friendId, err := uuid.Parse(fid)
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
	us.User = &dto.User{}
	us.User.ID = friendId
	if err := us.Get(ctx); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("User not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}
	fs := friendship.New()
	fs.UserID = userId
	fs.FriendID = friendId
	fs.Delete(ctx)

	if err := cache.Client().Del(ctx, userId.String()).Err(); err != nil {
		fmt.Println("Error invalidating cache:", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
