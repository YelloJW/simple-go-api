package main

import (
	"fmt"
	"log"

	usecases "github.com/YelloJW/simple-rest-api/app/modules/users/application/use_cases"
	"github.com/YelloJW/simple-rest-api/app/modules/users/infrastructure/controllers"
	"github.com/YelloJW/simple-rest-api/app/modules/users/infrastructure/repositories"
	"github.com/YelloJW/simple-rest-api/app/shared/config"
	"github.com/YelloJW/simple-rest-api/app/shared/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("Failed to load config")
	}
	mongoClient, err := db.Init(config.DbConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}

	usersCollection := mongoClient.Database("users").Collection("users")
	usersRepository := &repositories.UsersRepository{Collection: usersCollection}
	createUserUseCase := &usecases.CreateUser{UsersRepository: usersRepository}
	postUserController := &controllers.PostUserController{CreateUserUseCase: createUserUseCase}

	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			fmt.Println("Hello")
		})
		users := v1.Group("/users")
		{
			users.POST("", postUserController.Handle)
		}
	}

	router.Run(":8081")
}
