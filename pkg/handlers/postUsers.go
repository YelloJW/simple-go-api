package handlers

import (
	"net/http"

	"github.com/YelloJW/simple-rest-api/pkg/db"
	"github.com/YelloJW/simple-rest-api/pkg/db/models"
	"github.com/gin-gonic/gin"
)

type PostUserPayload struct {
	Name  string `json:"name" binding:"required"`
}

func PostUsers(c *gin.Context)  {
	var payload PostUserPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var user = models.User{Name: payload.Name}
	db.DB.Create(&user)


	c.IndentedJSON(http.StatusCreated, gin.H{"data": user})
}