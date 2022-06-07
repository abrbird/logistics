package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type OrderAvailabilityService interface {
	Retrieve(ctx context.Context, repository repository.OrderAvailabilityRepository, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve
	Update(ctx context.Context, repository repository.OrderAvailabilityRepository, orderAvailability models.OrderAvailability) error
}
