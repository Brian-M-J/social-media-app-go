package notifications

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/internals/notifications"
	"github.com/Brian-M-J/social-media-app-go/models/users"
	"github.com/Brian-M-J/social-media-app-go/services/friendships"
	"github.com/google/uuid"
)

// store
var Store map[uuid.UUID]chan string

// mutex
var mu sync.Mutex

func InitNotificationsSystem() {
	Store = make(map[uuid.UUID]chan string)
}

func Register(userID uuid.UUID) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := Store[userID]; !ok {
		Store[userID] = make(chan string)
	}
}

func ListenForNotifications(ctx context.Context, userID uuid.UUID) {
	mu.Lock()
	channel, ok := Store[userID]
	mu.Unlock()
	if !ok {
		fmt.Printf("No notification channel registered for user %v\n", userID)
		return
	}

	us := users.New()
	us.User = &dto.User{}
	us.User.ID = userId

	if err := us.Get(ctx); err != nil {
		return
	}

	for {
		select {
		case message := <-channel:
			fmt.Printf("Hey %v, you have a new notification: %v\n", us.User.Name, message)

		case <-ctx.Done():
			fmt.Printf("Stopping notification channel for user %v\n", userID)
			return
		}
	}
}

func NotifyUsers(ctx context.Context, userID uuid.UUID, msg string) {
	// get all friends and notify
	fs := friendships.New()
	fs.UserID = userID
	fs.GetAll(ctx)

	mu.Lock()
	defer mu.Unlock()

	for _, f := range fs.AllFriends {
		if ch, ok := Store[f.FriendID]; ok {
			go func() {
				ch <- msg
			}()
		}
	}
}

func Hydrate() {
	ctx := context.Background()

	us := users.New()
	us.Users = &dto.Users{}
	if err := us.GetAll(ctx); err != nil {
		log.Fatalf("Internal error: %v\n", err)
	}

	for _, u := range us.Users.Users {
		notifications.Register(u.User.ID)

		go notifications.ListenForNotifications(ctx, u.User.ID)

	}

}
