package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type SQLAddressRepository struct {
	store *SQLRepository
}

func (S SQLAddressRepository) Retrieve(ctx context.Context, AddressId int64) models.AddressRetrieve {
	//TODO implement me
	panic("implement me")
}
