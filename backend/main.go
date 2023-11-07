package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/timhi/gallery-backend/m/v2/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	api.LoadData()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/random-object", api.GetRandomRoadsideObject)
	e.Logger.Fatal(e.Start(":9007"))
}