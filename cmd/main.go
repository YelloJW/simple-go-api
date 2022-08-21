package main

import (
	"fmt"

	"github.com/YelloJW/simple-rest-api/pkg/controllers"
	"github.com/YelloJW/simple-rest-api/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	defer db.DB.Close()
	
	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			fmt.Println("Hi everyone")
		})
		users := v1.Group("/users")
		{
			users.POST("", controllers.PostUsers)
			users.GET("/:id", controllers.GetUserById)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}

	router.Run("localhost:8081")
}


