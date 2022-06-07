package implemented_service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type OrderAvailabilityService struct{}

func (o OrderAvailabilityService) Retrieve(ctx context.Context, repository repository.OrderAvailabilityRepository, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {
	//TODO implement me
	panic("implement me")
}

func (o OrderAvailabilityService) Update(ctx context.Context, repository repository.OrderAvailabilityRepository, orderAvailability models.OrderAvailability) error {
	//TODO implement me
	panic("implement me")
}
