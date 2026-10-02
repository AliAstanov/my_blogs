package handlers

import (
	"github.com/AliAstanov/my_blogs/models"
	"github.com/AliAstanov/my_blogs/pkg/halpers"
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

// CreateCategories godoc
// @Summary Create a new category
// @Description Creates a new category based on the request body
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryReq true "Create Category Request Body"
// @Success 201 {object} models.Category "Created category"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/own/category [post]
func (h *handlers) CreateCategories(ctx *gin.Context) {

	var reqBody = &models.CreateCategoryReq{}

	if err := ctx.BindJSON(&reqBody); err != nil {
		h.log.Error("error on binding request body:", logger.Error(err))
		return
	}

	category := &models.Category{}

	halpers.DataParser1(reqBody, category)

	category, err := h.storage.GetContentRepo().CreateCategory(ctx, category)
	if err != nil {
		h.log.Error(":", logger.Error(err))
		return
	}

	ctx.JSON(201, category)
}

// GetCategoriesList godoc
// @Summary Get categories list
// @Description Retrieves a list of categories with pagination
// @Tags categories
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of categories per page" default(10)
// @Success 200 {array} models.Category "List of categories"
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/pb/categories [get]
func (h *handlers) GetCategoriesList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	categories, err := h.storage.GetContentRepo().GetCategorys(ctx, halpers.GetPage(page), halpers.GetLimit(limit))
	if err != nil {
		h.log.Error("error on get contents:", logger.Error(err))
		return
	}
	ctx.JSON(200, categories)
}

// GetCategory godoc
// @Summary Get a single category by ID
// @Description Retrieves a single category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category "Category data"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/pb/category/{id} [get]
func (h *handlers) GetCategory(ctx *gin.Context) {

	id := ctx.Param("id")
	category := &models.Category{}

	category, err := h.storage.GetContentRepo().GetCategory(ctx, id)
	if err != nil {
		h.log.Error("error on getting category:", logger.Error(err))
		return
	}
	ctx.JSON(200, category)

}
