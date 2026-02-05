package router

import (
	"go-user-auth-service/internal/config"
	"go-user-auth-service/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindBodyWithJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": user})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hessage": "pong",
		})
	})
	return r
}
