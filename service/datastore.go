package service

import (
	"gopkg.in/mgo.v2"
)

var mysession, err = mgo.Dial("localhost")

func session() *mgo.Session {
	return mysession.Copy()
}
