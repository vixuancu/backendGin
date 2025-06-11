package handlerV1

import (
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type ProductHandler struct{}

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (u *ProductHandler) GetProductsV1(c *gin.Context) {
	search := c.Query("search")
	limitStr := c.DefaultQuery("limit", "10")

	if err := utils.ValidationRequied("Search", search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := utils.ValidationStringLength("Search", search, 3, 50); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := utils.ValidationRegex("Search", search, "chỉ chứa kí tự,số và khoảng trắng", searchRegex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "limit must be a positive integer",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "List all products V1",
		"search":  search,
		"limit":   limit,
	})
}
func (u *ProductHandler) GetProductsByIdV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID V1",
		"id":      id,
	})
}
func (u *ProductHandler) PostProducts(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "create a new product V1",
	})
}
func (u *ProductHandler) PutProducts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "update product by ID V1",
		"id":      id,
	})
}
func (u *ProductHandler) DeleteProducts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete product by ID V1",
		"id":      id,
	})
}
