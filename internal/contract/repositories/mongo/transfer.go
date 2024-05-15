package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"openmyth/blockchain/internal/contract/entities"
	"openmyth/blockchain/internal/contract/repositories"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

type transferRepository struct {
	databaseName string
	mgoClient    *mongoclient.MongoClient
	collection   *mongo.Collection
}

// NewTransferRepository initializes a new TransferRepository with the provided Mongo client and database name.
//
// - mgoClient: The Mongo client for the repository.
// - databaseName: The name of the database for the repository.
// Return type: repositories.TransferRepository.
func NewTransferRepository(mgoClient *mongoclient.MongoClient, databaseName string) repositories.TransferRepository {
	return &transferRepository{
		mgoClient:    mgoClient,
		databaseName: databaseName,
	}
}

// getCollection retrieves the collection for transfers.
//
// No parameters.
// Return type: *mongo.Collection.
func (r *transferRepository) getCollection() *mongo.Collection {
	if r.collection == nil {
		e := &entities.Transfer{}
		r.collection = r.mgoClient.Database(r.databaseName).Collection(e.TableName())
	}

	return r.collection
}

// Create creates a new transfer in the repository.
//
// Parameters:
// - ctx: The context for creating the transfer.
// - data: The transfer data to be inserted.
// Return type: error.
func (r *transferRepository) Create(ctx context.Context, data *entities.Transfer) error {
	res, err := r.getCollection().InsertOne(ctx, data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	data.ID = id.(primitive.ObjectID).Hex()

	return nil
}

// FindByFrom retrieves transfers by the 'from' field.
//
// Parameters:
// - ctx: The context for retrieving the transfers.
// - from: The sender of the transfers to retrieve.
// Return type: A slice of *entities.Transfer and an error.
func (r *transferRepository) FindByFrom(ctx context.Context, from string) ([]*entities.Transfer, error) {
	panic("not implemented") // TODO: Implement
}

// FindByTo retrieves transfers by the 'to' field.
//
// Parameters:
// - ctx: The context for retrieving the transfers.
// - to: The recipient of the transfers to retrieve.
// Return type: A slice of *entities.Transfer and an error.
func (r *transferRepository) FindByTo(ctx context.Context, to string) ([]*entities.Transfer, error) {
	panic("not implemented") // TODO: Implement
}

// GetList retrieves a list of transfers.
//
// - ctx: The context for retrieving the list of transfers.
// Return type: A slice of *entities.Transfer and an error.
func (r *transferRepository) GetList(ctx context.Context) ([]*entities.Transfer, error) {
	var result []*entities.Transfer
	cur, err := r.getCollection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
