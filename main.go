package main

import (
	"github.com/YelloJW/simple-rest-api/handlers"
	"github.com/gin-gonic/gin"
)




func main() {
	router := gin.Default()
	router.GET("/", handlers.HelloWorld)
	router.Run("localhost:8081")
}
