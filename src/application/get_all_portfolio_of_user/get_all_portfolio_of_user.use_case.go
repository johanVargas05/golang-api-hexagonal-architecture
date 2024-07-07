package get_all_portfolio_of_user_use_case

import (
	pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"
	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	portfolio_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/portfolio"
)



type GetAllPortfolioOfUserUseCase struct {
	managerCacheService portfolio_ports.AllPortFolioOfUserCacheServicePort
	getAllPortFolioOfUserService portfolio_ports.AllPortFolioOfUserQueryServicePort
}

func New(managerCacheService portfolio_ports.AllPortFolioOfUserCacheServicePort, getAllPortFolioOfUserService portfolio_ports.AllPortFolioOfUserQueryServicePort) *GetAllPortfolioOfUserUseCase {
	return &GetAllPortfolioOfUserUseCase{
		managerCacheService,
		getAllPortFolioOfUserService,
	}
}

func (useCase *GetAllPortfolioOfUserUseCase) Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, *pagination_dtos.PaginationResponseDto, error) {
	paramsCache:=*params
	data,pagination := useCase.managerCacheService.Get(&paramsCache)
	if data != nil {
		return data, pagination, nil
	}

	data, pagination, err := useCase.getAllPortFolioOfUserService.Execute(params)

	if err != nil {
		return nil, nil, err
	}

	useCase.managerCacheService.Set(&paramsCache, data, pagination)

	return data, pagination, nil
}