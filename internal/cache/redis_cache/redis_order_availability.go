package redis_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abrbird/logistics/internal/models"
	"time"
)

type RedisOrderAvailabilityCache struct {
	Prefix     string
	redisCache *RedisCache
	expiration time.Duration
}

func (r RedisOrderAvailabilityCache) getIDKey(orderId int64, issuePointId int64) string {
	return fmt.Sprintf("%s_%v_%v", r.Prefix, orderId, issuePointId)
}

func (r RedisOrderAvailabilityCache) Get(ctx context.Context, orderId int64, issuePointId int64) models.OrderAvailabilityRetrieve {
	//
	//	span, ctx := opentracing.StartSpanFromContext(ctx, "cache")
	//	defer span.Finish()
	//

	item, err := r.redisCache.client.Get(ctx, r.getIDKey(orderId, issuePointId)).Result()
	if err != nil {
		return models.OrderAvailabilityRetrieve{OrderAvailability: nil, Error: err}
	}

	var record models.OrderAvailability
	if err = json.Unmarshal([]byte(item), &record); err != nil {
		return models.OrderAvailabilityRetrieve{OrderAvailability: nil, Error: err}
	}

	return models.OrderAvailabilityRetrieve{OrderAvailability: &record, Error: nil}
}

func (r RedisOrderAvailabilityCache) Set(ctx context.Context, record models.OrderAvailability) error {
	//
	//	span, ctx := opentracing.StartSpanFromContext(ctx, "cache")
	//	defer span.Finish()
	//

	data, err := json.Marshal(record)
	if err != nil {
		return err
	}

	return r.redisCache.client.Set(ctx, r.getIDKey(record.OrderId, record.IssuePointId), data, r.expiration).Err()
}
