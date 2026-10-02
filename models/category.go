package models

import (
	"time"

	"github.com/google/uuid"
)

// Category represents the categories table in the database.
type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type GetCategoriesResp struct {
	Categories []*Category
	Count      int32
}

// SubCategory represents the sub_categories table in the database.
type SubCategory struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	CategoryID    uuid.UUID `json:"category_id"`
}
