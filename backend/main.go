package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/timhi/gallery-backend/m/v2/api"
)

func main() {
	api.LoadData()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/random-object", api.GetRandomRoadsideObject)
	e.Logger.Fatal(e.Start(":9007"))
}
