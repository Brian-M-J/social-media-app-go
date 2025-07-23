package posts

import (
	"time"

	"github.com/Brian-M-J/social-media-app-go/models/users"
	"github.com/google/uuid"
)

type Posts struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User users.Users `gorm:"foreignKey:UserID;references:ID" json:"-"`
}
