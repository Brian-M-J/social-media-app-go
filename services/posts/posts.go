package posts

import (
	"context"

	"github.com/Brian-M-J/social-media-app-go/services/posts"
	"github.com/google/uuid"
)

type Posts struct {
	ID    uuid.UUID
	Post  *dto.Post
	Posts *dto.Posts
}

func New() *Posts {
	return &Posts{}
}

func (p *Posts) Create(ctx context.Context) {
	m := posts.New()
	m.Post = p.Post
	m.Create(ctx)
	p.Post.UpdatedAt = nil
}

func (p *Posts) GetAll(ctx context.Context) {
	m := posts.New()
	m.UserID = p.UserID
	m.Posts = p.Posts
	m.Get(ctx)
}

func (p *Posts) Delete(ctx context.Context) error {
	m := posts.New()
	m.UserID = p.UserID
	m.ID = p.ID
	if err := m.Delete(ctx); err != nil {
		return err
	}
	return nil
}
