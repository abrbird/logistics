package implemented_service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type IssuePointService struct{}

func (i IssuePointService) Retrieve(ctx context.Context, repository repository.IssuePointRepository, issuePointId int64) models.IssuePointRetrieve {
	//TODO implement me
	panic("implement me")
}

func (i IssuePointService) RetrieveByAddress(ctx context.Context, repository repository.IssuePointRepository, addressId int64) models.IssuePointRetrieve {
	//TODO implement me
	panic("implement me")
}
