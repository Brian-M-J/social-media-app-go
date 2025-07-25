package dto

type PostCreate struct {
	Content string `json:"content" validate:"required,max=2000"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}
