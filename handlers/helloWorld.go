package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Hello World!"})
}