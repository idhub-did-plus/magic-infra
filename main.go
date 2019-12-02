// magic-infra project main.go
package main

import (
	"bytes"

	"encoding/json"
	"io/ioutil"
	"log"
	"magic-infra/model"
	"magic-infra/service"
	"net/http"

	_ "github.com/cloudflare/roughtime"

	_ "github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	jsonb, _ := ioutil.ReadAll(r.Body)
	jsonss := string(jsonb)
	if jsonss == "" {
		w.Write([]byte("{\"success\":true, \"message\":\"saved!\"}"))
		return
	}

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
	w.Write([]byte("{\"success\":true, \"message\":\"saved!\"}"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/saveDeployedToken", saveDeployedToken) //	设置访问路由
	mux.HandleFunc("/listDeployedTokens", listDeployedTokens)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
