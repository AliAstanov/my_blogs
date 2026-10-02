package postgres

import (
	"context"

	"github.com/AliAstanov/my_blogs/models"
	log "github.com/AliAstanov/my_blogs/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/saidamir98/udevs_pkg/logger"
)

type contentRepo struct {
	db  *pgxpool.Pool
	log log.Log
}

func NewContentRepo(db *pgxpool.Pool, log log.Log) ContnentRepoI {
	return &contentRepo{db, log}
}

func (c *contentRepo) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	c.log.Debug("request in CreateCategory")

	category.CategoryID = uuid.New()

	query := `
		INSERT INTO
			categories(
				category_id,
				name
				) VALUES(
				 $1,$2
				 )
				 RETURNING category_id, name, created_at`
	var result models.Category
	err := c.db.QueryRow(
		ctx, query,
		category.CategoryID,
		category.Name,
	).Scan(
		&result.CategoryID,
		&result.Name,
		&result.CreatedAt,
	)
	if err != nil {
		c.log.Error("error on creating new category.", logger.Error(err))
		return nil, err
	}

	// cat, err := c.GetCategory(ctx, category.CategoryID.String())
	// if err != nil {
	// 	c.log.Error("error on getting new category.", logger.Error(err))
	// 	return nil, err
	// }
	return &result, nil
}

func (c *contentRepo) GetCategorys(ctx context.Context, page, limit int32) (*models.GetCategoriesResp, error) {
	c.log.Debug("request in GetCategorys")

	query := `
		SELECT 
			*
		FROM
			categories
		LIMIT 
			$1
		OFFSET
			$2`
	offset := (page-1) * limit 

	rows ,err := c.db.Query(
		ctx,query,
		limit,
		offset,
	)
	if err != nil {
		c.log.Error("error on getting all categorys.", logger.Error(err))
		return nil, err
	}
	defer rows.Close()

	var categories  []*models.Category

	for rows.Next(){
		var category models.Category

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.CreatedAt,
		)
		if err != nil {
			c.log.Error("error on scanning category.", logger.Error(err))
			return nil, err
		}
		categories = append(categories,&category)
	}

	var count int32

	err = c.db.QueryRow(ctx,"SELECT count(*) FROM categories").Scan(&count)
	if err != nil {
		c.log.Error("error on scanning category count.", logger.Error(err))
		return nil, err
	}

	return &models.GetCategoriesResp{
		Categories: categories,
		Count: count,
	}, nil
}

func (c *contentRepo) GetCategory(ctx context.Context, id string) (*models.Category, error) {
	c.log.Debug("request in GetCategory")

	var category models.Category

	query := `
			SELECT 
				*
			FROM
				categories
			WHERE 
				category_id = $1
	`
	err := c.db.QueryRow(ctx, query, id).Scan(
		&category.CategoryID,
		&category.Name,
		&category.CreatedAt,
	)

	if err != nil {
		c.log.Error("error ongetting category.", logger.Error(err))
		return nil, err
	}

	return &category, nil
}

func (c *contentRepo) UpdateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return nil, nil
}
func (c *contentRepo) DeletteCategory(ctx context.Context, id string) error {
	return nil
}

func (c *contentRepo) CreateSubCategory(ctx context.Context, sub_category *models.SubCategory) (*models.SubCategory, error) {
	return nil, nil
}
func (c *contentRepo) GetSubCategorys(ctx context.Context, page, limit int32) ([]*models.SubCategory, error) {
	return nil, nil
}
func (c *contentRepo) GetSubCategory(ctx context.Context, id string) (*models.SubCategory, error) {
	return nil, nil
}
func (c *contentRepo) UpdateSubCategory(ctx context.Context, sub_category *models.SubCategory) (*models.SubCategory, error) {
	return nil, nil
}
func (c *contentRepo) DeleteSubCategory(ctx context.Context, id string) error {
	return nil
}

func (c *contentRepo) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) GetArticles(ctx context.Context, page, limit int32) ([]*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	return nil, nil
}
func (c *contentRepo) DeletteArticle(ctx context.Context, id string) error {
	return nil
}
