package service

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
	"github.com/abrbird/logistics/internal/repository"
)

type AddressService interface {
	Retrieve(ctx context.Context, repository repository.AddressRepository, AddressId int64) models.AddressRetrieve
}
