package portfolio_ports

import (
	pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"
	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
)

type AllPortFolioOfUserQueryServicePort interface {
	Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio,*pagination_dtos.PaginationResponseDto, error)
}

type AllPortFolioOfUserCacheServicePort interface {
	Get(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, *pagination_dtos.PaginationResponseDto)
	Set(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto, data []*entities.Portfolio, pagination *pagination_dtos.PaginationResponseDto)
}