package portfolio_routes

import (
	"github.com/labstack/echo/v4"

	get_all_portfolios_of_user_routes "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers/portfolio/get_all_portfolios_of_user.routes"
)

func InitRouters(e *echo.Echo) {
	get_all_portfolios_of_user_routes.InitRouters(e)
}
