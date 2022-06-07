package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type IssuePointService interface {
	Retrieve(ctx context.Context, repository repository.IssuePointRepository, issuePointId int64) models.IssuePointRetrieve
	RetrieveByAddress(ctx context.Context, repository repository.IssuePointRepository, addressId int64) models.IssuePointRetrieve
}
