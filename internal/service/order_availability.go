package service

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
	"github.com/abrbird/logistics/internal/repository"
)

type OrderAvailabilityService interface {
	Retrieve(ctx context.Context, repository repository.OrderAvailabilityRepository, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve
	Update(ctx context.Context, repository repository.OrderAvailabilityRepository, orderAvailability models.OrderAvailability) error

	RemoveOrder(
		ctx context.Context,
		oaRepository repository.OrderAvailabilityRepository,
		ipRepository repository.IssuePointRepository,
		orderId int64,
		addressId int64,
	) models.OrderAvailabilityRetrieve
}
