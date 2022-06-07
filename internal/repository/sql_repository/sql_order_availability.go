package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type SQLOrderAvailabilityRepository struct {
	store *SQLRepository
}

func (S SQLOrderAvailabilityRepository) Retrieve(ctx context.Context, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {
	//TODO implement me
	panic("implement me")
}

func (S SQLOrderAvailabilityRepository) Update(ctx context.Context, orderAvailability models.OrderAvailability) error {
	//TODO implement me
	panic("implement me")
}
