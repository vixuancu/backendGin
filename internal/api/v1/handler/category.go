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
