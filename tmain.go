package main

import (
	"fmt"

	"github.com/auoychai/gothqr/config"
)

func main() {

	fmt.Println("-----------", config.Cf.DB.URI)
	fmt.Println("-----------", config.Cf.DB.DbName)

	// fmt.Println("Result:", config.Get("Options.User"))

}
