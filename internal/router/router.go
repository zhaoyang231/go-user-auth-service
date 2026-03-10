package router

import (
	"go-user-auth-service/internal/config"
	"go-user-auth-service/internal/model"
	"go-user-auth-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userService := service.UserService{}
	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userService.CreateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": user})
	})

	r.GET("/users", func(c *gin.Context) {
		users, err := userService.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})
	return r
}
