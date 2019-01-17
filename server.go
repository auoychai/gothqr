package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/auoychai/gothqr/config"
	"github.com/auoychai/gothqr/controllers"
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

	//for test call data operation 
	service.Insert()
	service.GetPeople()
	// *****************************



	// Start point for Web API , Http Server 
	httpServer := echo.New()
	// Middleware
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())

	// CORS
	httpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	httpServer.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!\n")
	})
	httpServer.POST("/users", controllers.CreateUser)

	// Suparate start http server to sub-process with Goroutine
	// Let it main go routine jump to next section for handle shutting down
	go func() {
		err = httpServer.Start(":8080")

		if err != nil {
			if err != http.ErrServerClosed {
				panic("Error starting server! " + err.Error())
			} else {
				fmt.Printf("Goroutine |--> Shutting down...")
			}
		}

	}()

	fmt.Printf("Going to httpServer Shutdown handling ...")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = httpServer.Shutdown(ctx); err != nil {
		panic("Server did not shut down before timeout: " + err.Error())
	} else {
		fmt.Println("Server shutdown")
	}

	fmt.Println("Finally")
}
