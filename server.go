package main

import (
	"github.com/labstack/echo"
	"synack/controllers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

//go:generate sqlboiler --wipe mysql

//Docker & docker-compose
//Golang testing & benchmarking (testify/assert) --> https://github.com/stretchr/testify
//OpenAPI 3.0.3

func main() {
	e := echo.New()
	routes(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func routes(e *echo.Echo) {
	g := e.Group("/store")
	g.GET("/user/:id", controllers.GetUser)
	g.POST("/user", controllers.SaveUser)
	g.PUT("/user/:id", controllers.UpdateUser)
	g.DELETE("/user/:id", controllers.DeleteUser)
}
