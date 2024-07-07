package run_seed_controller

import (
	"errors"

	"github.com/labstack/echo/v4"

	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
	seed_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/seed"
)

type RunSeedController struct {
	useCase seed_ports.RunSeedUseCasePort
}

func New(useCase seed_ports.RunSeedUseCasePort) *RunSeedController {
	return &RunSeedController{
		useCase,
	}
}

func (c *RunSeedController) Execute(ctx echo.Context) error {
	err:=c.useCase.Execute()
	
	if errors.Is(err, seed_errors.ErrSeedAlreadyExecuted){
		return ctx.JSON(400, map[string]interface{}{
			"code":    400,
			"message": "Seed already executed",
		})
	}

	if errors.Is(err, seed_errors.ErrSeedNotExecuted){
		return ctx.JSON(500, map[string]interface{}{
			"code":    500,
			"message": "Seed not executed",
		})
	}

	if errors.Is(err, seed_errors.ErrLoadDataSeed){
		return ctx.JSON(500, map[string]interface{}{
			"code":    500,
			"message": "Error loading data seed",
		})
	}

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"code":    500,
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	ctx.JSON(200, map[string]interface{}{
			"code":    200,
			"message": "Seed executed",
	})

	return nil
}