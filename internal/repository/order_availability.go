package repository

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
)

type OrderAvailabilityRepository interface {
	Retrieve(ctx context.Context, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve
	Update(ctx context.Context, orderAvailability models.OrderAvailability) error
}
