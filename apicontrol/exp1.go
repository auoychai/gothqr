package apicontrol

import (
	"net/http"
	"time"

	"github.com/auoychai/gothqr/service"
	"github.com/labstack/echo"

	"encoding/json"
	"fmt"
	"strconv"
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


func QueryData(c echo.Context) error {

	name := c.QueryParam("name")

	tmp := "Name:" + name
	fmt.Println(tmp)

	age := 35

	xx := `{"firstName":"` + name + `","lastName":"Dow" ,"age":` + strconv.Itoa(age) + `}`
	fmt.Println("YY:", xx)
	rawIn := json.RawMessage(xx)

	var v interface{}
	_ = json.Unmarshal(rawIn, &v)

	return c.JSON(http.StatusOK, &v)
}


func PathData(c echo.Context) error {

	name := c.Param("name")

	tmp := "Name:" + name
	fmt.Println(tmp)

	age := 35

	xx := `{"firstName":"` + name + `","lastName":"Dow" ,"age":` + strconv.Itoa(age) + `}`
	fmt.Println("YY:", xx)
	rawIn := json.RawMessage(xx)

	var v interface{}
	_ = json.Unmarshal(rawIn, &v)

	return c.JSON(http.StatusOK, &v)
}


func CustomRetJson(c echo.Context) error {

	name := "Auoychai"
	tmp := "Name:" + name
	fmt.Println(tmp)

	age := 35

	xx := `{"firstName":"` + name + `","lastName":"Dow" ,"age":` + strconv.Itoa(age) + `}`
	fmt.Println("YY:", xx)
	rawIn := json.RawMessage(xx)

	var v interface{}
	_ = json.Unmarshal(rawIn, &v)

	// Let try time data ------------------
	start := time.Now()
	fmt.Println("Time:",start)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time-Elapsed:",elapsed)
	// -------------------------------------

	return c.JSON(http.StatusOK, &v)
}

