package service

import (
	"gopkg.in/mgo.v2"
)

var mysession, err = mgo.Dial("localhost")

func Session() *mgo.Session {
	return mysession.Copy()
}
