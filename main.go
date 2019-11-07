// magic-infra project main.go
package main

import (
	"bytes"

	"io/ioutil"
	"log"
	"magic-infra/model"
	"magic-infra/service"
	"net/http"

	"encoding/json"

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
func listDeployedTokens(w http.ResponseWriter, r *http.Request) {
	tokens := list()
	result, _ := json.Marshal(&tokens)
	w.Write(result)
}
func list() *[]model.DeployedToken {
	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("InfraRepository").C("DeployedToken")
	result := []model.DeployedToken{}
	err := c.Find(bson.M{"name": "byq"}).Limit(10).All(&result)
	if err != nil {

	}
	return &result

}
func saveDeployedToken(w http.ResponseWriter, r *http.Request) {
	jsonb, _ := ioutil.ReadAll(r.Body)
	jsonss := string(jsonb)
	log.Println(jsonss)
	jsons := bytes.NewReader(jsonb)

	dec := json.NewDecoder(jsons)
	result := model.DeployedToken{}
	dec.Decode(&result)

	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("InfraRepository").C("DeployedToken")
	c.Insert(&result)
}

func main() {
	http.HandleFunc("/saveDeployedToken", saveDeployedToken) //	设置访问路由
	http.HandleFunc("/", listDeployedTokens)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
