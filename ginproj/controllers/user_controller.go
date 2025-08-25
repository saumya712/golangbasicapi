package controllers

import (
	"GINPROJ/config"
	"GINPROJ/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Createuser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(200, user)
}

func Getuser(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(200, users)
}
