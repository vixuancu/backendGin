package handlerV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) GetUsersV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List all users V1",
	})
}
func (u *UserHandler) GetUsersByIdV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID V1",
		"id":      id,
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
