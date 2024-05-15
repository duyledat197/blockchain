package main

import (
	"context"
	"log"

	"openmyth/blockchain/internal/user-mgnt/entities"
	"openmyth/blockchain/internal/user-mgnt/repositories/mongo"
	mongoclient "openmyth/blockchain/pkg/mongo_client"
	"openmyth/blockchain/util"
)

func main() {
	ctx := context.Background()
	url := "mongodb://localhost:27017"
	mgoClient := mongoclient.NewMongoClient(url)

	if err := mgoClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer mgoClient.Close(ctx)

	userRepo := mongo.NewUserRepository(mgoClient, "test")
	pwd, _ := util.HashPassword("test_password")
	if err := userRepo.Create(ctx, &entities.User{
		UserName:       "test_user",
		HashedPassword: pwd,
	}); err != nil {
		log.Fatalf("unable to create user: %v", err)
	}
}
