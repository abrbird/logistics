package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type SQLOrderAvailabilityRepository struct {
	store *SQLRepository
}

func (S SQLOrderAvailabilityRepository) Retrieve(ctx context.Context, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {
	const query = `
		SELECT 
    		order_id,
			issue_point_id,
			status
		FROM logistics_orders_availability
		WHERE order_id = $1 AND issue_point_id = $2
	`

	orderAvailability := &models.OrderAvailability{}
	if err := S.store.dbConnectionPool.QueryRow(
		ctx,
		query,
		orderId,
		issuePointId,
	).Scan(
		&orderAvailability.OrderId,
		&orderAvailability.IssuePointId,
		&orderAvailability.Status,
	); err != nil {
		return models.OrderAvailabilityRetrieve{OrderAvailability: nil, Error: models.NotFoundError}
	}
	return models.OrderAvailabilityRetrieve{OrderAvailability: orderAvailability, Error: nil}
}

func (S SQLOrderAvailabilityRepository) Update(ctx context.Context, orderAvailability models.OrderAvailability) error {
	const query = `
		UPDATE logistics_orders_availability
		SET (status) = ($3)
		WHERE order_id = $1 AND issue_point_id = $2
	`

	err := S.store.dbConnectionPool.QueryRow(
		ctx,
		query,
		orderAvailability.OrderId,
		orderAvailability.IssuePointId,
		orderAvailability.Status,
	)
	if err != nil {
		return models.NotFoundError
	}
	return nil
}
