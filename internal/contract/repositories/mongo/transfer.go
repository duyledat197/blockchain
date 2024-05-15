package mongo

import (
	"context"

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

func NewTransferRepository(mgoClient *mongoclient.MongoClient, databaseName string) repositories.TransferRepository {
	return &transferRepository{
		mgoClient:    mgoClient,
		databaseName: databaseName,
	}
}

func (r *transferRepository) getCollection() *mongo.Collection {
	if r.collection == nil {
		e := &entities.Transfer{}
		r.collection = r.mgoClient.Database(r.databaseName).Collection(e.TableName())
	}

	return r.collection
}

func (r *transferRepository) Create(ctx context.Context, data *entities.Transfer) error {
	res, err := r.getCollection().InsertOne(ctx, data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	data.ID = id.(primitive.ObjectID).Hex()

	return nil
}

func (r *transferRepository) FindByFrom(ctx context.Context, from string) ([]*entities.Transfer, error) {
	panic("not implemented") // TODO: Implement
}

func (r *transferRepository) FindByTo(ctx context.Context, to string) ([]*entities.Transfer, error) {
	panic("not implemented") // TODO: Implement
}

func (r *transferRepository) GetList(ctx context.Context) ([]*entities.Transfer, error) {
	var result []*entities.Transfer
	cur, err := r.getCollection().Find(ctx, &entities.Transfer{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
