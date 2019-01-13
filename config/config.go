package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Configuration struct {
	Name     string
	Port     int
	Base_URL string
	DB       struct {
		URI string
		DbName string
	}
	Options struct {
		User                string
		Pass                string
		AuthSource          string
		AutoIndex           bool
		ReconnectTries      int64
		ReconnectInterval   int
		PoolSize            int
		BufferMaxEntries    int
		MaxStalenessSeconds int
	}
	FCM_url string
	FCM_key string
}

var err error
var Cf *Configuration

var tStruct *structs.Struct

func init() {

	fmt.Println("First time config initialize")
	c := flag.String("c", "./config/config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	vConfig := Configuration{}
	err = decoder.Decode(&vConfig)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
	Cf = &vConfig

	tStruct = structs.New(vConfig)

}

func get(config string) string {

	var xs bool = true
	var field *structs.Field

	// yy := strings.Split("127.0.0.1:5432", ":")

	yy := strings.Split(config, ".")
	// fmt.Println("Size-of-token:", yy)

	sLen := len(yy)

	if sLen == 1 {

		//fmt.Println("Token:", yy[0])
		field, xs = tStruct.FieldOk(yy[0])

		if xs == false && field == nil {

			return ""
		}

	} else {
		//fmt.Println("Token:", yy[0])
		field, xs = tStruct.FieldOk(yy[0])
		if xs == false && field == nil {

			fmt.Println("Field not found:",yy[0])
			return ""
		}

		fmt.Println("L1 Result:",field.Value())

		for i := 1; i < sLen; i++ {

			token := yy[i]
			fmt.Println("Field Name:",token)
			
			field, xs = field.FieldOk(token)

			if xs == false && field == nil {

				fmt.Println("Field not found:",token)
				return ""
			}
		}

	}

	ret := field.Value()

	if reflect.TypeOf(ret).Kind() != reflect.TypeOf("string").Kind() {
		//fmt.Println("YYYYY:", ret)
		return ""
	}

	//fmt.Println("BBBBB:", ret)
	return ret.(string)

}
