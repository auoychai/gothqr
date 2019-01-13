package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Merchant struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Status     string        `bson:"status" json:"status"`
	Store      string        `bson:"store" json:"store"`
	MerchantId string        `bson:"merchantId" json:"merchantId"`
	Name       string        `bson:"name" json:"name"`
	Address    string        `bson:"address" json:"address"`
	Tel        string        `bson:"tel" json:"tel"`
	Notifymode string        `bson:"notifymode" json:"notifymode"`
	Urlhook    string        `bson:"urlhook" json:"urlhook"`
	Hook_key   string        `bson:"hook_key" json:"hook_key"`
	UpdateAt   time.Time     `bson:"updateAt" json:"updateAt"`
	CreateAt   time.Time     `bson:"createAt" json:"createAt"`
}
