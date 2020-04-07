package service

import (
	"fmt"

	"github.com/idhub-did-plus/magic-infra/config"

	"gopkg.in/mgo.v2"
)

var mysession *mgo.Session

func init() {
	var err error
	mysession, err = mgo.Dial(config.Config.Mysql.Url)
	if err != nil {
		fmt.Println(err)
	}
}

func Session() *mgo.Session {
	return mysession.Copy()
}
