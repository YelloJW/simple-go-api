package main

import (
	"fmt"

	"github.com/YelloJW/simple-rest-api/app/shared/db"
	"github.com/YelloJW/simple-rest-api/app/shared/env"
	"github.com/YelloJW/simple-rest-api/app/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.DB.Close()
	env := env.New()
	
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

	router.Run(":8081")
}


