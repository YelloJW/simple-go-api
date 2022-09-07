package models

import (
	"time"

	"github.com/YelloJW/simple-rest-api/app/modules/users/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	Id        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Email     string             `bson:"email"`
	Created   time.Time          `bson:"created"`
}

func (doc *UserDocument) ToModel() *models.User {
	return &models.User{Id: doc.Id.Hex(), FirstName: doc.FirstName, LastName: doc.LastName, Email: doc.Email}
}
