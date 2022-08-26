package controllers

import (
	"net/http"

	"github.com/YelloJW/simple-rest-api/app/packages/db"
	"github.com/YelloJW/simple-rest-api/app/packages/db/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context){
	if err := db.DB.Delete(&models.User{}, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}