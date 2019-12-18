package service

import (
	"magic-infra/config"

	"gopkg.in/mgo.v2"
)

var mysession, err = mgo.Dial(config.Config.Mysql.Url)

func Session() *mgo.Session {
	return mysession.Copy()
}
