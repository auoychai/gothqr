package apicontrol

import (
	"net/http"
	"fmt"
	_ "github.com/auoychai/gothqr/service"
	"github.com/labstack/echo"
)

func PayWithThQR(c echo.Context) error {

	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}

	fmt.Println("JSon Request Data:", m)
	fmt.Println("Data-merchantId:", m["merchantId"])
	fmt.Println("Data-storeId:", m["storeId"])
	fmt.Println("Data-terminalId:", m["terminalId"])

	return c.JSON(http.StatusOK, m)

}
