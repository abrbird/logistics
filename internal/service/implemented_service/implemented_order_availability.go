package implemented_service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type OrderAvailabilityService struct{}

func (o OrderAvailabilityService) Retrieve(ctx context.Context, repository repository.OrderAvailabilityRepository, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {
	return repository.Retrieve(ctx, orderId, issuePointId)
}

func (o OrderAvailabilityService) Update(ctx context.Context, repository repository.OrderAvailabilityRepository, orderAvailability models.OrderAvailability) error {
	return repository.Update(ctx, orderAvailability)
}
