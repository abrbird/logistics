package implemented_service

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
	"github.com/abrbird/logistics/internal/repository"
)

type AddressService struct {
	service *Service
}

func (a AddressService) Retrieve(ctx context.Context, repository repository.AddressRepository, AddressId int64) models.AddressRetrieve {
	//TODO implement me
	panic("implement me")
}
