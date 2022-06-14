package implemented_service

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/cache/redis_cache"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
)

type OrderAvailabilityService struct {
	service *Service
}

func (o OrderAvailabilityService) Retrieve(ctx context.Context, repository repository.OrderAvailabilityRepository, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {

	retrieved := o.service.cache.OrderAvailability().Get(ctx, orderId, issuePointId)
	if errors.Is(retrieved.Error, redis_cache.Nil) {
		retrieved = repository.Retrieve(ctx, orderId, issuePointId)

		if retrieved.Error != nil {
			return retrieved
		}
	}

	if err := o.service.cache.OrderAvailability().Set(ctx, *retrieved.OrderAvailability); err != nil {
		retrieved.OrderAvailability = nil
		retrieved.Error = err
		return retrieved
	}

	return retrieved
}

func (o OrderAvailabilityService) Update(ctx context.Context, repository repository.OrderAvailabilityRepository, orderAvailability models.OrderAvailability) error {
	err := repository.Update(ctx, orderAvailability)
	if err != nil {
		return err
	}

	if err = o.service.cache.OrderAvailability().Set(ctx, orderAvailability); err != nil {
		return err
	}

	return nil
}

func (o OrderAvailabilityService) RemoveOrder(
	ctx context.Context,
	oaRepository repository.OrderAvailabilityRepository,
	ipRepository repository.IssuePointRepository,
	orderId int64,
	addressId int64,
) models.OrderAvailabilityRetrieve {
	issuePointRetrieved := o.service.IssuePoint().RetrieveByAddress(
		ctx,
		ipRepository,
		addressId,
	)

	if issuePointRetrieved.Error != nil {
		return models.OrderAvailabilityRetrieve{
			OrderAvailability: nil,
			Error:             models.NotFoundError(issuePointRetrieved.Error),
		}
	}

	if !issuePointRetrieved.IssuePoint.IsAvailable {
		return models.OrderAvailabilityRetrieve{
			OrderAvailability: nil,
			Error:             models.IssuePointIsUnavailableError(nil),
		}
	}

	orderAvailabilityRetrieved := o.Retrieve(
		ctx,
		oaRepository,
		orderId,
		issuePointRetrieved.IssuePoint.Id,
	)

	if orderAvailabilityRetrieved.Error != nil {
		orderAvailabilityRetrieved.OrderAvailability = nil
		orderAvailabilityRetrieved.Error = models.NotFoundError(orderAvailabilityRetrieved.Error)

		return orderAvailabilityRetrieved
	}

	if orderAvailabilityRetrieved.OrderAvailability.Status == models.Available {
		orderAvailabilityRetrieved.OrderAvailability.Status = models.Issued

		err := o.Update(
			ctx,
			oaRepository,
			*orderAvailabilityRetrieved.OrderAvailability,
		)
		if err != nil {
			retryError := models.NewRetryError(err)
			orderAvailabilityRetrieved.Error = retryError
			return orderAvailabilityRetrieved
		}

		return orderAvailabilityRetrieved
	}

	if orderAvailabilityRetrieved.OrderAvailability.Status == models.Issued {
		return orderAvailabilityRetrieved
	}

	if orderAvailabilityRetrieved.OrderAvailability.Status == models.Moved {
		orderAvailabilityRetrieved.Error = models.OrderIsMovedMessageError(nil)
		return orderAvailabilityRetrieved
	}

	retryError := models.NewRetryError(nil)
	orderAvailabilityRetrieved.Error = retryError
	return orderAvailabilityRetrieved
}
