package sql_repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
)

type SQLIssuePointRepository struct {
	store *SQLRepository
}

func (S SQLIssuePointRepository) Retrieve(ctx context.Context, issuePointId int64) models.IssuePointRetrieve {
	const query = `
		SELECT 
    		id,
			address_id,
			is_available
		FROM logistics_issue_points
		WHERE id = $1
	`

	issuePointRecord := &models.IssuePoint{}
	if err := S.store.dbConnectionPool.QueryRow(
		ctx,
		query,
		issuePointId,
	).Scan(
		&issuePointRecord.Id,
		&issuePointRecord.AddressId,
		&issuePointRecord.IsAvailable,
	); err != nil {
		return models.IssuePointRetrieve{IssuePoint: nil, Error: models.NotFoundError(err)}
	}
	return models.IssuePointRetrieve{IssuePoint: issuePointRecord, Error: nil}
}

func (S SQLIssuePointRepository) RetrieveByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve {
	const query = `
		SELECT 
    		id,
			address_id,
			is_available
		FROM logistics_issue_points
		WHERE address_id = $1
	`

	issuePointRecord := &models.IssuePoint{}
	if err := S.store.dbConnectionPool.QueryRow(
		ctx,
		query,
		addressId,
	).Scan(
		&issuePointRecord.Id,
		&issuePointRecord.AddressId,
		&issuePointRecord.IsAvailable,
	); err != nil {
		return models.IssuePointRetrieve{IssuePoint: nil, Error: models.NotFoundError(err)}
	}
	return models.IssuePointRetrieve{IssuePoint: issuePointRecord, Error: nil}
}
