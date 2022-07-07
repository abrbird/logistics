package service

import (
	"context"
	"github.com/abrbird/logistics/internal/models"
	"github.com/abrbird/logistics/internal/repository"
)

type IssuePointService interface {
	Retrieve(ctx context.Context, repository repository.IssuePointRepository, issuePointId int64) models.IssuePointRetrieve
	RetrieveByAddress(ctx context.Context, repository repository.IssuePointRepository, addressId int64) models.IssuePointRetrieve
}
