package main

import (
	"github.com/YelloJW/simple-rest-api/pkg/controllers"
	"github.com/YelloJW/simple-rest-api/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", controllers.GetUsers)
			users.GET("/:id", controllers.GetUserById)
			users.POST("", controllers.PostUsers)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}

	router.Run("localhost:8081")
	// db.DB.Close()
}
