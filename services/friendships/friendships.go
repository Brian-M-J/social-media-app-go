package friendships

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Brian-M-J/social-media-app-go/internals/cache"
	"github.com/Brian-M-J/social-media-app-go/internals/dto"
	"github.com/Brian-M-J/social-media-app-go/models/friendship"
	"github.com/google/uuid"
)

type Friends struct {
	UserID     uuid.UUID
	FriendID   uuid.UUID
	Friends    *dto.Friends
	AllFriends []dto.AllFriends
}

func New() *Friends {
	return &Friends{}
}

func (f *Friends) Create(ctx context.Context) {
	m := friendship.New()
	m.Friends = f.Friends
	m.Create(ctx)
	f.Friends.UpdatedAt = nil
}

func (f *Friends) GetAll(ctx context.Context) {
	val, err := cache.Client().Get(ctx, f.UserID.String()).Result()
	if val != "" && err == nil {
		json.Unmarshal([]byte(val), &f.AllFriends)
		return
	}
	m := friendship.New()
	m.UserID = f.UserID
	m.Get(ctx)
	f.AllFriends = m.AllFriends
	b, _ := json.Marshal(f.AllFriends)
	if err := cache.Client().Set(ctx, f.UserID.String(), b, 24*time.Hour).Err(); err != nil {
		fmt.Printf("Error setting up cache for friends: %v\n", err)
	}
}

func (f *Friends) Delete(ctx context.Context) error {
	m := friendship.New()
	m.UserID = f.UserID
	m.FriendID = f.FriendID
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}
