package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func HelloWorld(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, gin.H{"data": "Hello World!"})
}

func main() {
	router := gin.Default()
	router.GET("/", HelloWorld)
	router.Run("localhost:8081")
}
