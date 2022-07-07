package repository

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
)

type AddressRepository interface {
	Retrieve(ctx context.Context, AddressId int64) models.AddressRetrieve
}
