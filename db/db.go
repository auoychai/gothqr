package db

import (
	_ "fmt"

	"gopkg.in/mgo.v2"
)

type DBProperty struct {
	DbURI  string
	DbName string
}

// var DBProp *DBProperty

var tsession *mgo.Session
var err error
var DbCon *mgo.Database

func (dbp *DBProperty) Connect() {

	tsession, err = mgo.Dial(dbp.DbURI)
	if err != nil {
		panic(err)
	}

	// fmt.Println("DB-Connect/session:%s", tsession)
	tsession.SetMode(mgo.Monotonic, true)
	DbCon = tsession.DB(dbp.DbName)

}

func DisConnect() {

	tsession.Close()
}
