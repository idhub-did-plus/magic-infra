// token_controller
package controller

import (
	"bytes"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/idhub-did-plus/magic-infra/model"
	"github.com/idhub-did-plus/magic-infra/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ListDeployedTokens(c *gin.Context) {
	owner := c.Query("owner")
	tokens := list(owner)
	//result, _ := json.Marshal(&tokens)
	c.JSON(200, gin.H{
		"success": true,
		"data":    tokens,
	})

}
func list(owner string) *[]model.DeployedToken {
	session := service.Session()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("InfraRepository").C("DeployedToken")
	result := []model.DeployedToken{}
	err := c.Find(bson.M{"owneraccount": owner}).Limit(10).All(&result)
	if err != nil {

	}
	defer session.Close()
	return &result

}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func SaveDeployedToken(c *gin.Context) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
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
	result := model.DeployedToken{}
	dec.Decode(&result)
	result.OwnerAccount = result.DeployAccount

	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("InfraRepository").C("DeployedToken")
	collection.Insert(&result)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "saved!",
	})

}
