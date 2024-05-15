package mongoclient

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	*mongo.Client
	connectString string
}

func NewMongoClient(connectString string) *MongoClient {
	return &MongoClient{
		connectString: connectString,
	}
}

// Connect initializes the MongoClient by connecting to the MongoDB server.
//
// It takes a context.Context parameter for managing the operation's deadline.
// Returns an error if the connection attempt fails.
func (m *MongoClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	log.Println("Connecting to MongoDB...")
	clientOpts := options.Client()
	clientOpts.ApplyURI(m.connectString)
	clientOpts.SetConnectTimeout(10 * time.Second)
	clientOpts.SetCompressors([]string{"snappy", "zlib", "zstd"})

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	log.Print("Connected to MongoDB!!")

	m.Client = client

	return nil
}

// Close closes the MongoClient by disconnecting from the MongoDB server.
//
// It takes a context.Context parameter for managing the operation's deadline.
// Returns an error if the disconnection attempt fails.
func (m *MongoClient) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
