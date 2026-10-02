package postgres

import (
	"context"

	"github.com/AliAstanov/my_blogs/models"
)

type ContnentRepoI interface {
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	GetCategorys(ctx context.Context, page, limit int32) (*models.GetCategoriesResp, error)
	GetCategory(ctx context.Context, id string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	DeletteCategory(ctx context.Context, id string) error

	CreateSubCategory(ctx context.Context, sub_category *models.SubCategory) (*models.SubCategory, error)
	GetSubCategorys(ctx context.Context, page, limit int32) ([]*models.SubCategory, error)
	GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error)
	UpdateSubCategory(ctx context.Context, sub_category *models.SubCategory) (*models.SubCategory, error)
	DeleteSubCategory(ctx context.Context, id string) error

	CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	GetArticles(ctx context.Context, page, limit int32) ([]*models.Article, error)
	GetArticle(ctx context.Context, id string) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	DeletteArticle(ctx context.Context, id string) error
}

type OwnerRepoI interface {
	Login(ctx context.Context, login *models.LoginOwn)(*models.Owner, error)
}

type CommonRepoI interface {
	CheckIsExists(ctx context.Context, req *models.Common) (bool, error)
}
