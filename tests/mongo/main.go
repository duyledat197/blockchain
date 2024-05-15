package main

import (
	"context"
	"log"

	"openmyth/blockchain/config"
	"openmyth/blockchain/internal/contract/repositories/mongo"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
)

func main() {
	ctx := context.Background()
	db := &config.Database{
		Schema:   "mongodb",
		Host:     "localhost",
		Port:     "27017",
		Database: "blockchain",
	}
	mgoClient := mongoclient.NewMongoClient(db.Address())

	if err := mgoClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer mgoClient.Close(ctx)

	transferRepo := mongo.NewTransferRepository(mgoClient, db.Database)
	result, _ := transferRepo.GetList(ctx)
	log.Println(len(result))
}
