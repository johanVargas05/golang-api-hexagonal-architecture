package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/primary/api/routers"
)

func main() {
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println("Error loading .env file")
	}

	port:=os.Getenv("PORT")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	routers.InitRoutes(e)

	if err := e.Start(":"+port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}