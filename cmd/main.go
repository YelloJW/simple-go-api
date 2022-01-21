package main

import (
	"github.com/YelloJW/simple-rest-api/pkg/db"
	"github.com/YelloJW/simple-rest-api/pkg/handlers"
	"github.com/gin-gonic/gin"
)




func main() {
	db.Connect()
	router := gin.New()
	router.POST("/users", handlers.PostUsers)
	router.GET("/users", handlers.GetUsers)
	router.Run("localhost:8081")
	db.DB.Close()
}
