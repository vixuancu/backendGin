package handlerV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (u *ProductHandler) GetProductsV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all products V1",
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
