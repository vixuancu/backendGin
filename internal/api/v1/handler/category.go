package handlerV1

import "github.com/gin-gonic/gin"

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
	if validCategory[category] == false {
		c.JSON(400, gin.H{
			"error": "Invalid category",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "List all categories V1",
		"category": category,
	})
}
