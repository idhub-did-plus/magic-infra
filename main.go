// magic-infra project main.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"magic-infra/model"
	"magic-infra/service"
	"net/http"
	"strings"

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
func listDeployedTokens() *[]model.DeployedToken {
	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("InfraRepository").C("DeployedToken")
	result := []model.DeployedToken{}
	err := c.Find(bson.M{"name": "Ale"}).Limit(10).All(&result)
	if err != nil {

	}
	return &result

}
func saveDeployedToken(w http.ResponseWriter, r *http.Request) {
	jsonb, _ := ioutil.ReadAll(r.Body)
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func test() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
