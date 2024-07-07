package manager_cache_ports

import (
	"time"

	dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/manager_cache"
)

type ManagerCacheRepositoryPort interface {
	GetData(params *dtos.CacheParams) error
	SetData(params *dtos.CacheParams, expiration time.Duration) error
}
