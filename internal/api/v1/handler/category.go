package handlerV1

import (
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct{}

var validCategory = map[string]bool{
	"php":        true,
	"javascript": true,
	"golang":     true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (cg *CategoryHandler) GetCategoriesV1(c *gin.Context) {
	category := c.Param("category")
	// Check if the category is valid
	if err := utils.ValidationInList("Category", category, validCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "List all categories V1",
		"category": category,
	})
}
