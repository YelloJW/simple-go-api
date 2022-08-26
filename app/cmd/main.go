package main

import (
	"fmt"

	"github.com/YelloJW/simple-rest-api/app/packages/controllers"
	"github.com/YelloJW/simple-rest-api/app/packages/db"
	"github.com/YelloJW/simple-rest-api/app/packages/env"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	env := env.Load()
	defer db.DB.Close()
	
	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			fmt.Println(env.Env)
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


