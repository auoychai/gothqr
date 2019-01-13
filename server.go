package main

import (
	"fmt"

	"github.com/auoychai/gothqr/config"
	"github.com/auoychai/gothqr/db"
	"github.com/auoychai/gothqr/service"
	_ "github.com/labstack/echo"
)

type Person struct {
	Name  string
	Phone string
}

var err error

func main() {

	fmt.Println("-----------", config.Cf.DB.URI)
	fmt.Println("-----------", config.Cf.DB.DbName)

	dbProp := db.DBProperty{}
	dbProp.DbURI = config.Cf.DB.URI
	dbProp.DbName = config.Cf.DB.DbName

	dbProp.Connect()
	defer db.DisConnect()

	service.Insert()
	service.GetPeople()

	/*
		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		// CORS
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE },
		}))

		e.GET("/",func(c echo.Context) error {

			return c.String(http.StatusOK, "Hello, World!\n")
		})

		e.POST("/users",controllers.CreateUser)

		// Server

		e.Start(":8080")
	*/
}
