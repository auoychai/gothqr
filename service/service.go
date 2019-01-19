package service

import (
	"fmt"
	"log"

	"github.com/auoychai/gothqr/db"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

var err error

func Insert() {

	c := db.DbCon.C("people")

	err = c.Insert(&Person{"Ale-XX", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})

	if err != nil {
		log.Fatal(err)
	}

}

func GetPeople() ([]Person,error) {

	var person []Person

	tDB := db.DbCon
	fmt.Println(tDB)

	c := tDB.C("people")

	result := Person{}

	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	err = c.Find(bson.M{}).All(&person)
/*
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Phone:", result.Phone)

	for _, info := range person {

		fmt.Println("Name , Phone", info.Name, " ", info.Phone)

	}
*/
	return person,err
}
