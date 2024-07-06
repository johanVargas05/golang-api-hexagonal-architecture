package seed_ports

import "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"

type RunSeedRepositoryPort interface {
	Execute([]*entities.Portfolio) error
}

type SeedRunCheckRepositoryPort interface {
	Execute() error
}

type LoadDataSeedRepositoryPort interface {
	Execute() ([]*entities.Portfolio, error)
}

