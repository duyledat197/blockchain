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

type approvalRepository struct {
	databaseName string
	mgoClient    *mongoclient.MongoClient
	collection   *mongo.Collection
}

// NewApprovalRepository initializes a new ApprovalRepository with the provided Mongo client and database name.
func NewApprovalRepository(mgoClient *mongoclient.MongoClient, databaseName string) repositories.ApprovalRepository {
	return &approvalRepository{
		mgoClient:    mgoClient,
		databaseName: databaseName,
	}
}

// getCollection retrieves the collection for approvals.
func (r *approvalRepository) getCollection() *mongo.Collection {
	if r.collection == nil {
		e := &entities.Approval{}
		r.collection = r.mgoClient.Database(r.databaseName).Collection(e.TableName())
	}

	return r.collection
}

// Create creates a new approval in the repository.
func (r *approvalRepository) Create(ctx context.Context, data *entities.Approval) error {
	res, err := r.getCollection().InsertOne(ctx, data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	data.ID = id.(primitive.ObjectID).Hex()

	return nil
}

// FindByOwner retrieves approvals by owner.
func (r *approvalRepository) FindByOwner(ctx context.Context, owner string) ([]*entities.Approval, error) {
	var result []*entities.Approval

	cur, err := r.getCollection().Find(ctx, &entities.Approval{
		Owner: owner,
	})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetList retrieves a list of approvals.
func (r *approvalRepository) GetList(ctx context.Context) ([]*entities.Approval, error) {
	var result []*entities.Approval
	cur, err := r.getCollection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
