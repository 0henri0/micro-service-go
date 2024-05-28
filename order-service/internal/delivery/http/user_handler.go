package http

import (
	"github.com/gin-gonic/gin"
	"micro-service-go/internal/services"
	"net/http"
)

func NewUserRoute(handle *UserHandler) *gin.Engine {
	route := gin.Default()
	route.GET("/users", handle.GetUsers)

	return route
}

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
