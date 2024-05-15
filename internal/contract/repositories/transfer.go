package repositories

import (
	"context"

	"openmyth/blockchain/internal/contract/entities"
)

type TransferRepository interface {
	Create(context.Context, *entities.Transfer) error
	FindByFrom(ctx context.Context, from string) ([]*entities.Transfer, error)
	FindByTo(ctx context.Context, to string) ([]*entities.Transfer, error)
	GetList(context.Context) ([]*entities.Transfer, error)
}
