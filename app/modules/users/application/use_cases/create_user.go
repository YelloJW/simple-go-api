package usecases

import (
	"github.com/YelloJW/simple-rest-api/app/modules/users/domain/models"
	"github.com/YelloJW/simple-rest-api/app/modules/users/infrastructure/repositories"
)

type CreateUser struct {
	UsersRepository *repositories.UsersRepository
}

func (useCase *CreateUser) Execute(firstName, lastName, email string) (*models.User, error) {
	user := &models.User{FirstName: firstName, LastName: lastName, Email: email}
	return useCase.UsersRepository.Save(user)
}
