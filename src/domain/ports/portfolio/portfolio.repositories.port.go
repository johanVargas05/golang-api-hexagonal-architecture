package portfolio_ports

import (
	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
)

type AllPortFolioOfUserQueryRepositoryPort interface {
	Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, int, error)
}