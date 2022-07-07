package repository

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
)

type IssuePointRepository interface {
	Retrieve(ctx context.Context, issuePointId int64) models.IssuePointRetrieve
	RetrieveByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve
}
