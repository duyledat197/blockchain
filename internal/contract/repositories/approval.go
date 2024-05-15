package repositories

import (
	"context"

	"openmyth/blockchain/internal/contract/entities"
)

type ApprovalRepository interface {
	Create(context.Context, *entities.Approval) error
	FindByOwner(ctx context.Context, owner string) ([]*entities.Approval, error)
	GetList(context.Context) ([]*entities.Approval, error)
}
