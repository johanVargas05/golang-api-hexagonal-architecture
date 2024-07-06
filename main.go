package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err:=godotenv.Load()
	if err!=nil{
		fmt.Println("Error loading .env file")
	}

	port:=os.Getenv("PORT")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	if err := e.Start(":"+port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}