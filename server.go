package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/auoychai/gothqr/apicontrol"
	"github.com/auoychai/gothqr/config"
	"github.com/auoychai/gothqr/db"
	"github.com/auoychai/gothqr/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// for test call data operation
	// this call for inital or repeat insert data to mongodb collection
	// You can comment after you don't need to repeate insert data on very next run this api service
	service.Insert()
	// *****************************
	// Start point for Web API , Http Server

	eHttpServer := echo.New()

	// Declare Middleware
	eHttpServer.Use(middleware.Logger())
	eHttpServer.Use(middleware.Recover())
	// CORS
	eHttpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Declare api route
	eHttpServer.POST("/users", apicontrol.CreateUser)
	eHttpServer.GET("/people", apicontrol.GetPeople)

	eHttpServer.GET("/querydata", apicontrol.QueryData)
	eHttpServer.GET("/pathdata/:name", apicontrol.PathData)
	eHttpServer.GET("/customReturnJson", apicontrol.CustomRetJson)

	eHttpServer.POST("/genqr", apicontrol.PayWithThQR)

	// Suparate start http server to sub-process with Goroutine
	// Let it main go routine jump to next section for handle shutting down
	go func() {
		err = eHttpServer.Start(":8080")

		if err != nil {
			if err != http.ErrServerClosed {
				panic("Error starting server! " + err.Error())
			} else {
				fmt.Printf("Goroutine |--> Shutting down...")
			}
		}

	}()

	fmt.Printf("Going to eHttpServer Shutdown handling ...")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = eHttpServer.Shutdown(ctx); err != nil {
		panic("Server did not shut down before timeout: " + err.Error())
	} else {
		fmt.Println("Server shutdown")
	}

	fmt.Println("Finally")

}
