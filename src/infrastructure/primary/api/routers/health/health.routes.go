package routers

import (
	"github.com/labstack/echo/v4"

	controllers "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/controllers/health"
)

func InitRouters(e *echo.Echo) {
	controller:=controllers.New()
	e.GET("/health", controller.Execute)
}