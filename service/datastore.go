package service

import (
	"gopkg.in/mgo.v2"
)

var mysession, err = mgo.Dial("mongodb://localhost:27017")

func Session() *mgo.Session {
	return mysession.Copy()
}
