package handlerV1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"strconv"
)

type UserHandler struct{}

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

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID của người dùng phải là số nguyên dương",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"id":      id,
	})
}
func (u *UserHandler) GetUsersByUidV1(c *gin.Context) {

	uid := c.Param("uid")
	uuid, err := uuid.Parse(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user UUID",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"uuid":    uuid,
	})
}
func (u *UserHandler) GetUsersBySlugV1(c *gin.Context) {

	slug := c.Param("slug")
	// slug chỉ được phép chứa chữ cái, số và dấu gạch ngang hoặc dấu .
	if !slugRegex.MatchString(slug) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user slug",
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
