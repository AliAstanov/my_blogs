package models

import (
	"time"

	"github.com/google/uuid"
)

// Article represents the articles table in the database.
type Article struct {
	ArticleID     uuid.UUID  `json:"article_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	CategoryID    uuid.UUID  `json:"category_id"`
	SubCategoryID uuid.UUID  `json:"sub_category_id"`
}
