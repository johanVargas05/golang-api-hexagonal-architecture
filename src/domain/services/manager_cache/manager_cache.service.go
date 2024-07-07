package manager_cache_services

import (
	"fmt"
	"time"

	dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/manager_cache"
	ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/manager_cache"
)

type managerCacheService struct {
	repository ports.ManagerCacheRepositoryPort
}

func New(repository ports.ManagerCacheRepositoryPort) *managerCacheService {
	return &managerCacheService{
		repository: repository,
	}
}

func (service *managerCacheService) GetData(key string, structure interface{}) (interface{}, error) {
	params := dtos.CacheParams{
		Key:   key,
		Value: structure,
	}

	err := service.repository.GetData(&params)

	if err != nil {
		return nil, err
	}

	return params.Value, nil
}

func (service *managerCacheService) SetData(key string, value interface{}, expiration time.Duration) {
	params := dtos.CacheParams{
		Key:   key,
		Value: value,
	}

	err := service.repository.SetData(&params, expiration)

	if err != nil {
		fmt.Println(err.Error())
	}

}
