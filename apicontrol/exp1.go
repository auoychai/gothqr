package apicontrol

import (
	"net/http"

	"github.com/auoychai/gothqr/service"
	"github.com/labstack/echo"
)

// create user
func CreateUser(c echo.Context) error {
	//user := new(models.User)

	user := "I'am User"

	// var err error
	return c.JSON(http.StatusCreated, user)
}

func GetPeople(c echo.Context) error {

	pJson, err := service.GetPeople()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pJson)
}

/*
	//for test call data operation
	pJson, _ := service.GetPeople()
	b, err := json.Marshal(pJson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
*/
