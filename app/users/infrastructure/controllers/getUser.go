package controllers

import (
	"fmt"
	"net/http"

	"github.com/YelloJW/simple-rest-api/app/shared/db"
	"github.com/YelloJW/simple-rest-api/app/users/infrastructure/models"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context){
	var user models.User
	fmt.Println(c.Param("id"))
	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}