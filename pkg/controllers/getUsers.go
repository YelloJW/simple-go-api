package controllers

import (
	"net/http"

	"github.com/YelloJW/simple-rest-api/pkg/db"
	"github.com/YelloJW/simple-rest-api/pkg/db/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context)  {
	var users []models.User
	db.DB.Find(&users)
	c.IndentedJSON(http.StatusOK, gin.H{"data": users})
}