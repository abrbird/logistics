package redis_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abrbird/logistics/internal/models"
	"time"
)

type RedisIssuePointCache struct {
	Prefix     string
	redisCache *RedisCache
	expiration time.Duration
}

func (r RedisIssuePointCache) getAddressIdKey(addressId int64) string {
	return fmt.Sprintf("%s_address_id_%v", r.Prefix, addressId)
}

func (r RedisIssuePointCache) GetByAddress(ctx context.Context, addressId int64) models.IssuePointRetrieve {
	//
	//	span, ctx := opentracing.StartSpanFromContext(ctx, "cache")
	//	defer span.Finish()
	//

	item, err := r.redisCache.client.Get(ctx, r.getAddressIdKey(addressId)).Result()
	if err != nil {
		return models.IssuePointRetrieve{IssuePoint: nil, Error: err}
	}

	var record models.IssuePoint
	if err = json.Unmarshal([]byte(item), &record); err != nil {
		return models.IssuePointRetrieve{IssuePoint: nil, Error: err}
	}

	return models.IssuePointRetrieve{IssuePoint: &record, Error: nil}
}

func (r RedisIssuePointCache) SetByAddress(ctx context.Context, issuePoint models.IssuePoint) error {
	//
	//	span, ctx := opentracing.StartSpanFromContext(ctx, "cache")
	//	defer span.Finish()
	//

	data, err := json.Marshal(issuePoint)
	if err != nil {
		return err
	}

	return r.redisCache.client.Set(ctx, r.getAddressIdKey(issuePoint.AddressId), data, r.expiration).Err()
}
