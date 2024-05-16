package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/internal/user-mgnt/repositories"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
	"openmyth/blockchain/pkg/xerror"
)

type userRepository struct {
	databaseName string
	mgoClient    *mongoclient.MongoClient
	collection   *mongo.Collection
}

// NewUserRepository creates a new instance of the UserRepository interface.
//
// It takes a *mongoclient.MongoClient and a string representing the database name as parameters.
// It returns a *userRepository implementing the UserRepository interface.
func NewUserRepository(mgoClient *mongoclient.MongoClient, databaseName string) repositories.UserRepository {
	return &userRepository{
		databaseName: databaseName,
		mgoClient:    mgoClient,
	}
}
func (r *userRepository) getCollection() *mongo.Collection {
	if r.collection == nil {
		e := &entities.User{}
		r.collection = r.mgoClient.Database(r.databaseName).Collection(e.TableName())
	}

	return r.collection
}

// Create inserts a new user into the user repository.
//
// It takes a context.Context and a pointer to an entities.User as parameters.
// It returns an error if there was a problem inserting the user.
func (r *userRepository) Create(ctx context.Context, data *entities.User) error {
	res, err := r.getCollection().InsertOne(ctx, data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	data.ID = id.(primitive.ObjectID)

	return nil
}

// FindUser finds a user by their ID in the user repository.
//
// It takes a context.Context and a string representing the user ID as parameters.
// It returns a pointer to an entities.User and an error. If the user is not found,
// it returns xerror.ErrNotFound. If there is an error retrieving the user, it
// returns the error.
func (r *userRepository) FindUser(ctx context.Context, id string) (*entities.User, error) {
	var result entities.User
	oID, _ := primitive.ObjectIDFromHex(id)
	if err := r.getCollection().FindOne(ctx, &entities.User{
		ID: oID,
	}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerror.ErrNotFound
		}

		return nil, err
	}

	return &result, nil
}

// FindUserByUsername finds a user by their username in the user repository.
//
// It takes a context.Context and a string representing the username as parameters.
// It returns a pointer to an entities.User and an error.
func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var result entities.User

	if err := r.getCollection().FindOne(ctx, &entities.User{
		UserName: username,
	}).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerror.ErrNotFound
		}

		return nil, err
	}

	return &result, nil
}
