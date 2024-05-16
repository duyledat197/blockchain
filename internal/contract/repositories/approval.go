package repositories

import (
	"context"

	"openmyth/blockchain/internal/contract/entities"
)

// ApprovalRepository defines the contract for approval repository.
type ApprovalRepository interface {
	// Create creates a new approval in the database.
	Create(context.Context, *entities.Approval) error

	// FindByOwner returns a list of approvals for a given owner.
	FindByOwner(ctx context.Context, owner string) ([]*entities.Approval, error)

	// GetList returns a list of all approvals in the database.
	GetList(context.Context) ([]*entities.Approval, error)
}
