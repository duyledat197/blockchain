package repositories

import (
	"context"

	"openmyth/blockchain/internal/contract/entities"
)

// TransferRepository is the interface for operations on the transfer repository.
type TransferRepository interface {
	// Create creates a new transfer in the database.
	Create(context.Context, *entities.Transfer) error
	// FindByFrom finds all transfers with a given sender.
	FindByFrom(ctx context.Context, from string) ([]*entities.Transfer, error)
	// FindByTo finds all transfers with a given receiver.
	FindByTo(ctx context.Context, to string) ([]*entities.Transfer, error)
	// GetList finds all transfers in the database.
	GetList(context.Context) ([]*entities.Transfer, error)
}
