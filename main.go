// magic-infra project main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"magic-infra/model"
	"magic-infra/service"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func test() {
	session := service.Session()

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err := c.Insert(&model.DeployedToken{"Ale", "+55 53 8116 9639"},
		&model.DeployedToken{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := model.DeployedToken{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bbb").C("jjjj")
	result := model.DeployedToken{}
	c.Find("yyy").One(&result)
	fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
