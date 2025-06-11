package handlerV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewsHandler struct{}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}
func (n *NewsHandler) GetNewsV1(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "List all news V1",
			"slug":    "No news",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "List all news V1",
			"slug":    slug,
		})
	}
}
