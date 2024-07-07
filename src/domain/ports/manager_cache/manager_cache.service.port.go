package manager_cache_ports

import (
	"time"
)

type ManagerCacheServicePort interface {
	GetData(key string, structure interface{}) (interface{}, error)
	SetData(key string, value interface{}, expiration time.Duration)
}
