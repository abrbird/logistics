package implemented_service

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/cache/redis_cache"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type IssuePointService struct {
	service *Service
}

func (i IssuePointService) Retrieve(ctx context.Context, repository repository.IssuePointRepository, issuePointId int64) models.IssuePointRetrieve {
	//TODO implement me
	panic("implement me")
}

func (i IssuePointService) RetrieveByAddress(ctx context.Context, repository repository.IssuePointRepository, addressId int64) models.IssuePointRetrieve {

	retrieved := i.service.cache.IssuePoint().GetByAddress(ctx, addressId)
	if errors.Is(retrieved.Error, redis_cache.Nil) {
		retrieved = repository.RetrieveByAddress(ctx, addressId)

		if retrieved.Error != nil {
			return retrieved
		}
	}

	if err := i.service.cache.IssuePoint().SetByAddress(ctx, *retrieved.IssuePoint); err != nil {
		retrieved.IssuePoint = nil
		retrieved.Error = err
		return retrieved
	}

	return retrieved
}
