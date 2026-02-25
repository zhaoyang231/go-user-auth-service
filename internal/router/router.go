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
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": user})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []model.User
		if err := config.DB.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"hessage": "pong",
		})
	})
	return r
}
