package manager_cache_repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/manager_cache"
)

type ManagerCacheRepository struct {
	redisClient *redis.Client
}

func New(redisClient *redis.Client) *ManagerCacheRepository {
	return &ManagerCacheRepository{redisClient: redisClient}
}

func (repository *ManagerCacheRepository) GetData(params *dtos.CacheParams) error {
	value, err := repository.redisClient.Get(context.Background(), params.Key).Result()

	if err != nil && err != redis.Nil {
		return err
	}

	if err == redis.Nil {
		return err
	}

	bytes := []byte(value)
	err = json.Unmarshal(bytes, &params.Value)

	if err != nil {
		return err
	}

	return nil
}

func (repository *ManagerCacheRepository) SetData(params *dtos.CacheParams, expiration time.Duration) error {
	bytes, err := json.Marshal(params.Value)

	if err != nil {
		return err
	}

	data := string(bytes)

	response := repository.redisClient.Set(context.Background(), params.Key, data, expiration)

	if response.Err() != nil {
		return response.Err()
	}

	return nil
}
