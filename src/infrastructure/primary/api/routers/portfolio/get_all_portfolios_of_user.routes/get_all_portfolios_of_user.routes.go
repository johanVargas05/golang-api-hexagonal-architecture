package get_all_portfolios_of_user_routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	get_all_portfolio_of_user_use_case "github.com/johanVargas05/golang-api-hexagonal-architecture/src/application/get_all_portfolio_of_user"
	all_portfolios_of_user_service "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/services/all_portfolios_of_user"
	manager_cache_services "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/services/manager_cache"
	constants_api "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/constants"
	get_all_portfolios_of_suer_controller "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/controllers/portfolio/get_all_portfolios_of_suer"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/pkg"
	all_port_folio_of_user_repository "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/mongo/portfolio/all_portfolio_of_user"
	manager_cache_repository "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/redis/manager_cache/redis"
)

func InitRouters(e *echo.Echo) {
	controller := constructController()
	patch:= fmt.Sprintf("%s/user/portfolio/:consumerId",constants_api.PREFIX_API)
	e.POST(patch, controller.Execute)
}

func constructController() *get_all_portfolios_of_suer_controller.GetAllPortfoliosOfUserController {

	db := pkg.NewMongoConnection()
	redisDb:=pkg.GetClientRedis()


	managerCacheRepository:=manager_cache_repository.New(redisDb)
	getAllPortfolioRepository:=all_port_folio_of_user_repository.New(db)

	managerCacheService := manager_cache_services.New(managerCacheRepository)
	getAllPortfolioService:= all_portfolios_of_user_service.New(getAllPortfolioRepository)
	managerCachePortfolioService := all_portfolios_of_user_service.NewManagerCache(managerCacheService)

	useCase := get_all_portfolio_of_user_use_case.New(managerCachePortfolioService,getAllPortfolioService)

	return get_all_portfolios_of_suer_controller.New(useCase)
}