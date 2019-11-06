package service

import (
	"gopkg.in/mgo.v2"
)

var mysession, err = mgo.Dial("mongodb://admin:123456@localhost:27017/admin")

func Session() *mgo.Session {
	return mysession.Copy()
}
