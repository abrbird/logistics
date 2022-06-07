package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type SQLIssuePointRepository struct {
	store *SQLRepository
}

func (S SQLIssuePointRepository) Retrieve(ctx context.Context, issuePointId int64) models.IssuePointRetrieve {
	//TODO implement me
	panic("implement me")
}

func (S SQLIssuePointRepository) RetrieveByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve {
	//TODO implement me
	panic("implement me")
}
