// login_controller
package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"io/ioutil"
	"log"

	"github.com/idhub-did-plus/magic-infra/model"
	"github.com/idhub-did-plus/magic-infra/service"
	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

func claimList(c *gin.Context) {
	body := c.Request.Body
	jsonb, _ := ioutil.ReadAll(body)
	jsonss := string(jsonb)
	if jsonss == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "no data!",
		})

		return
	}

	log.Println(jsonss)
	jsons := bytes.NewReader(jsonb)

	dec := json.NewDecoder(jsons)
	result := model.IssueAgent{}
	dec.Decode(&result)

	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("InfraRepository").C("IssueAgent")
	collection.Insert(&result)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "register successful!",
	})

}

func issue_claim(c *gin.Context) {

}
