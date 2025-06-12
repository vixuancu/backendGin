package handlerV1

import (
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct{}
type GetCategoriesV1Params struct {
	Category string `uri:"category" binding:"required,oneof=php javascript golang"` // Category must be one of the valid categories
}

// x-www-form-urlencoded
type PostCategoriesV1Params struct {
	Name   string `form:"name" binding:"required,min=2,max=50"` // Name must be a string with length between 2 and 50 characters
	Status string `form:"status" binding:"required,oneof=1 2"`  // Status must be either "active" or "inactive"
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (cg *CategoryHandler) GetCategoriesV1(c *gin.Context) {
	var params GetCategoriesV1Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(200, gin.H{
		"message":  "List all categories V1",
		"category": params.Category,
	})
}
func (cg *CategoryHandler) PostCategoriesV1(c *gin.Context) {
	var params PostCategoriesV1Params
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}

	c.JSON(200, gin.H{
		"message": "Create new category V1",
		"category": gin.H{
			"name":   params.Name,
			"status": params.Status,
		},
	})
}
