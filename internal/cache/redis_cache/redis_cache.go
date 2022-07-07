package redis_cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/config"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/cache"
	"time"
)

const (
	OrderAvailabilityPrefix = "order_availability"
	IssuePointPrefix        = "issue_point"
)

const (
	Nil = redis.Nil
)

type RedisCache struct {
	client            *redis.Client
	orderAvailability *RedisOrderAvailabilityCache
	issuePoint        *RedisIssuePointCache
}

func New(cfg config.Redis) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0, // use default DB
	})

	redisCache := &RedisCache{
		client: rdb,
	}

	redisCache.orderAvailability = &RedisOrderAvailabilityCache{
		Prefix:     OrderAvailabilityPrefix,
		redisCache: redisCache,
		expiration: time.Minute * 15,
	}
	redisCache.issuePoint = &RedisIssuePointCache{
		Prefix:     IssuePointPrefix,
		redisCache: redisCache,
		expiration: time.Minute * 15,
	}

	return redisCache
}

func (r RedisCache) OrderAvailability() cache.OrderAvailabilityCache {
	return r.orderAvailability
}

func (r RedisCache) IssuePoint() cache.IssuePointCache {
	return r.issuePoint
}
