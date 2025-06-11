package handlerV1

import (
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

type ProductHandler struct{}
type GetProductsBySlugV1Params struct {
	Slug string `uri:"slug" binding:"slug,min=3,max=50"` // Slug must be a string with length between 3 and 50 characters
}
type GetProductsV1Params struct {
	Search string `form:"search" binding:"required,min=3,max=50,search"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}
type PostProductsParams struct {
	Name    string        `json:"name" binding:"required,min=2,max=50"`
	Price   int           `json:"price" binding:"required,gte=0,lte=1000000"` // Price must be a non-negative integer
	Display *bool         `json:"display" binding:"omitempty"`                // Display must be true or false
	Image   ProductsImage `json:"image" binding:"required"`                   // dive dùng để validate các trường bên trong struct
}

type ProductsImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImgaeLink string `json:"image_link" binding:"required,file_ext=jpg png gif"` // ImageLink must be a valid URL
}

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}
func (u *ProductHandler) GetProductsV1(c *gin.Context) {
	var params GetProductsV1Params
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	if params.Limit == 0 {
		params.Limit = 10 // Set default limit if not provided
	}
	if params.Email == "" {
		params.Email = "vixuancu2004@gmail.com"
	}
	if params.Date == "" {
		params.Date = time.Now().Format("2006-01-02") // Set default date if not provided
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "List all products V1",
		"search":  params.Search,
		"limit":   params.Limit,
		"email":   params.Email,
		"date":    params.Date,
	})
}
func (u *ProductHandler) GetProductsBySlugV1(c *gin.Context) {
	var params GetProductsBySlugV1Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID V1",
		"slug":    params.Slug,
	})
}
func (u *ProductHandler) PostProducts(c *gin.Context) {
	var params PostProductsParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	if params.Display == nil {
		defaultDisplay := true // Set default value for Display if not provided
		params.Display = &defaultDisplay
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "create a new product V1",
		"data":    params.Name,
		"price":   params.Price,
		"display": params.Display,
		"image":   params.Image,
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
