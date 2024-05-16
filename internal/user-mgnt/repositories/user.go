package repositories

import (
	"context"

	"openmyth/blockchain/internal/user-mgnt/entities"
)

// UserRepository defines the interface for the user repository
// It provides methods to create and find users
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, data *entities.User) error
	// FindUser finds a user by its ID
	FindUser(ctx context.Context, id string) (*entities.User, error)
	// FindUserByUsername finds a user by its username
	FindUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
