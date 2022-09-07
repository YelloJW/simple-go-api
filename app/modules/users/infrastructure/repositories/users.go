package repositories

import (
	"context"
	"time"

	domain_models "github.com/YelloJW/simple-rest-api/app/modules/users/domain/models"
	db_models "github.com/YelloJW/simple-rest-api/app/modules/users/infrastructure/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersRepository struct {
	Collection *mongo.Collection
}

func (repo *UsersRepository) Save(newUser *domain_models.User) (*domain_models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	var userDocument db_models.UserDocument
	err := repo.Collection.FindOneAndUpdate(
		ctx,
		bson.M{"email": newUser.Email},
		bson.M{
			"$set": bson.M{
				"email":     newUser.Email,
				"firstName": newUser.FirstName,
				"lastName":  newUser.LastName,
			},
			"$setOnInsert": bson.M{
				"created": time.Now().UTC(),
			},
		},
		&opts,
	).Decode(&userDocument)

	if err != nil {
		return nil, err
	}

	return userDocument.ToModel(), nil
}
