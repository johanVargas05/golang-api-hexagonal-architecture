package seed_routes

import (
	"github.com/labstack/echo/v4"

	run_seed_use_case "github.com/johanVargas05/golang-api-hexagonal-architecture/src/application/run_seed"
	run_seed_service "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/services/run_seed"
	seed_run_check_service "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/services/seed_run_check"
	run_seed_controller "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/controllers/seed"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/pkg"
	load_data_seed_repository "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/files/seed/load_data_seed"
	run_seed_repository "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/mongo/seed/run_seed"
	seed_run_check_repository "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/mongo/seed/seed_run_check"
)

func InitRouters(e *echo.Echo) {
	controller:=constructController()
	e.GET("/seed", controller.Execute)
}

func constructController()*run_seed_controller.RunSeedController {

	db:=pkg.NewMongoConnection()

	loadDataSeedRepository:=load_data_seed_repository.New()
	runSeedRepository:=run_seed_repository.New(db)
	seedRunCheckRepository:=seed_run_check_repository.New(db)
	
	runSeedService:=run_seed_service.New(runSeedRepository,loadDataSeedRepository)
	checkSeedService:=seed_run_check_service.New(seedRunCheckRepository)
	useCase:=run_seed_use_case.New(runSeedService,checkSeedService)

	return run_seed_controller.New(useCase)
}