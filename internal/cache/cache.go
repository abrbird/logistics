package cache

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type Cache interface {
	OrderAvailability() OrderAvailabilityCache
	IssuePoint() IssuePointCache
}

type OrderAvailabilityCache interface {
	Get(ctx context.Context, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve
	Set(ctx context.Context, orderAvailability models.OrderAvailability) error
}

type IssuePointCache interface {
	GetByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve
	SetByAddress(ctx context.Context, issuePoint models.IssuePoint) error
}
