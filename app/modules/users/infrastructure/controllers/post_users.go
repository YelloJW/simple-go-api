package controllers

import (
	"net/http"

	usecases "github.com/YelloJW/simple-rest-api/app/modules/users/application/use_cases"
	"github.com/gin-gonic/gin"
)

type PostUserPayload struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type PostUserResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type PostUserController struct {
	CreateUserUseCase *usecases.CreateUser
}

func (controller *PostUserController) Handle(c *gin.Context) {
	var payload PostUserPayload

	if err := c.BindJSON(&payload); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.CreateUserUseCase.Execute(payload.FirstName, payload.LastName, payload.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.IndentedJSON(http.StatusOK, &PostUserResponse{Id: user.Id, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})
}
