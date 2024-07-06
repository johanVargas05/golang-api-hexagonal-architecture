package controllers

import (
	"time"

	"github.com/labstack/echo/v4"
)


type HealthController struct{}

func New() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Execute(ctx echo.Context) error {
	response:=map[string]any{
		"code":200,
		"data": map[string]any {
			"status":"UP",
			"current_time":time.Now().Format("2006-01-02 15:04:05"),
		},
		"message":"API is healthy",
	}
	
	return ctx.JSON(200, response)		
}