package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type AddressService interface {
	Retrieve(ctx context.Context, repository repository.AddressRepository, AddressId int64) models.AddressRetrieve
}