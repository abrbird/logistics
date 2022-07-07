package sql_repository

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
)

type SQLAddressRepository struct {
	store *SQLRepository
}

func (S SQLAddressRepository) Retrieve(ctx context.Context, AddressId int64) models.AddressRetrieve {
	const query = `
		SELECT 
    		id,
			address
		FROM logistics_addresses
		WHERE id = $1
	`

	address := &models.Address{}
	if err := S.store.dbConnectionPool.QueryRow(
		ctx,
		query,
		AddressId,
	).Scan(
		&address.Id,
		&address.Address,
	); err != nil {
		return models.AddressRetrieve{Address: nil, Error: models.NotFoundError(err)}
	}
	return models.AddressRetrieve{Address: address, Error: nil}
}
