package handlerV1

import (
	"fmt"
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

type NewsHandler struct{}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

// form-data
type PostNewsV1Params struct {
	Title  string `form:"title" binding:"required,min=2,max=50"` // Name must be a string with length between 2 and 50 characters
	Status string `form:"status" binding:"required,oneof=1 2"`   // Status must be either "active" or "inactive"
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
func (n *NewsHandler) PostNewsV1(c *gin.Context) {
	var params PostNewsV1Params
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Image is required",
		})
		return
	}
	// os.ModePerm = 0777
	// có nghĩa là được phép đọc, ghi và thực thi cho tất cả mọi người
	err = os.MkdirAll("./upload", os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create upload directory",
		})
		return
	}

	dst := fmt.Sprintf("./upload/%s", filepath.Base(image.Filename))
	if err = c.SaveUploadedFile(image, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save uploaded file",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Create new News V1",
		"news": gin.H{
			"title":  params.Title,
			"status": params.Status,
			"image":  image.Filename,
			"path":   dst,
		},
	})
}
