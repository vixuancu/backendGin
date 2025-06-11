package handlerV1

import (
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserHandler struct{}
type GetUsersByIdV1Params struct {
	Id int `uri:"id" binding:"required,gt=0"` // ID must be a positive integer
}
type GetUsersByUidV1Params struct {
	Uid string `uri:"uid" binding:"required,uuid"` // UUID must be a valid UUID
}

var slugRegex = regexp.MustCompile("^[a-z0-9]+(?:[-.][a-z0-9]+)*$")

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUsersV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all users V1",
	})
}
func (u *UserHandler) GetUsersByIdV1(c *gin.Context) {

	var params GetUsersByIdV1Params

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"id":      params.Id,
	})
}
func (u *UserHandler) GetUsersByUidV1(c *gin.Context) {

	var params GetUsersByUidV1Params
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"uuid":    params.Uid,
	})
}
func (u *UserHandler) GetUsersBySlugV1(c *gin.Context) {

	slug := c.Param("slug")
	// slug chỉ được phép chứa chữ cái, số và dấu gạch ngang hoặc dấu .
	if err := utils.ValidationRegex("Slug", slug, "chỉ được phép chứa chữ cái, số và dấu gạch ngang hoặc dấu .", slugRegex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"slug":    slug,
	})
}
func (u *UserHandler) PostUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "create a new user V1",
	})
}
func (u *UserHandler) PutUsers(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "update user by ID V1",
		"id":      id,
	})
}
func (u *UserHandler) DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete user by ID V1",
		"id":      id,
	})
}
