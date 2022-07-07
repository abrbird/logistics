package repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type IssuePointRepository interface {
	Retrieve(ctx context.Context, issuePointId int64) models.IssuePointRetrieve
	RetrieveByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve
}
