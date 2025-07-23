package friendship

import (
	"github.com/Brian-M-J/social-media-app-go/models/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendships struct {
	gorm.Model
	UserID   uuid.UUID   `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendID uuid.UUID   `gorm:"uniqueIndex:idx_user_friend" json:"friend_id"`
	User     users.Users `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Friend   users.Users `gorm:"foreignKey:FriendID;references:ID" json:"-"`
}
