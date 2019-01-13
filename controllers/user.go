package  controllers

import(
	"github.com/labstack/echo"
	"net/http"
)

// create user
func CreateUser( c echo.Context) error {
	//user := new(models.User)

	user := "I'am User"

	// var err error
	return  c.JSON(http.StatusCreated, user)
}