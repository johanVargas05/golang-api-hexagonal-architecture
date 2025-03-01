package routers

import (
	"github.com/labstack/echo/v4"

	routersHealth "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers/health"
	portfolio_routes "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers/portfolio"
	seed_routes "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers/seed"
)

func InitRoutes(e *echo.Echo) {
	routersHealth.InitRouters(e)
	seed_routes.InitRouters(e)
	portfolio_routes.InitRouters(e)
}