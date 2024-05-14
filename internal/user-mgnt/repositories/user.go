package repositories

import (
	"context"

	"openmyth/blockchain/internal/user-mgnt/entities"
)

type UserRepository interface {
	Create(ctx context.Context, data *entities.User) error
	FindUser(ctx context.Context, id string) (*entities.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entities.User, error)
}
