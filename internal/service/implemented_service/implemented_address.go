package implemented_service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type AddressService struct{}

func (a AddressService) Retrieve(ctx context.Context, repository repository.AddressRepository, AddressId int64) models.AddressRetrieve {
	//TODO implement me
	panic("implement me")
}
