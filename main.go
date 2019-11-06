// magic-infra project main.go
package main

import (
	_ "fmt"
	"log"
	"net/http"

	"magic-infra/model"
	"magic-infra/service"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func save(token *model.DeployedToken) {
	session := service.Session()

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("InfraRepository").C("DeployedToken")
	err := c.Insert(token)
	if err != nil {
		log.Fatal(err)
	}

}
func listDeployedTokens() *[]model.DeployedToken {
	session := service.Session()

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("InfraRepository").C("DeployedToken")

	result := []model.DeployedToken{}
	err := c.Find(bson.M{"name": "Ale"}).Limit(10).All(&result)
	if err != nil {

	}
	return &result

}
func myHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
