package repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type AddressRepository interface {
	Retrieve(ctx context.Context, AddressId int64) models.AddressRetrieve
}
