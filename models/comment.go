package models

import (
	"time"

	"github.com/google/uuid"
)

// Comment represents the comments table in the database.
type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID uuid.UUID `json:"article_id"`
	ViewerID  uuid.UUID `json:"viewr_id"`
}
