package routers

import (
	"github.com/labstack/echo/v4"

	routersHealth "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers/health"
)

func InitRoutes(e *echo.Echo) {
	routersHealth.InitRouters(e)
}