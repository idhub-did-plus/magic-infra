// magic-infra project main.go
package main

import (
	"bytes"

	"encoding/json"
	"io/ioutil"
	"log"
	"magic-infra/model"
	"magic-infra/service"
	_ "net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func listDeployedTokens(c *gin.Context) {
	tokens := list()
	//result, _ := json.Marshal(&tokens)
	c.JSON(200, gin.H{
		"success": true,
		"data":    tokens,
	})

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

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 这是允许访问所有域
		// 设置返回格式是json

		c.Next() //  处理请求
	}
}

func saveDeployedToken(c *gin.Context) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body := c.Request.Body
	jsonb, _ := ioutil.ReadAll(body)
	jsonss := string(jsonb)
	if jsonss == "" {
		c.JSON(200, gin.H{
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

	session := service.Session()

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("InfraRepository").C("DeployedToken")
	collection.Insert(&result)

	c.JSON(200, gin.H{
		"success": true,
		"message": "saved!",
	})

}

func main() {
	r := gin.Default()
	r.Use(CorsMiddleware())
	r.POST("/saveDeployedToken", saveDeployedToken)
	r.GET("/listDeployedTokens", listDeployedTokens)

	r.Run() // 在 0.0.0.0:8080 上监听并服务
	// mux := http.NewServeMux()
	// mux.HandleFunc("/saveDeployedToken", saveDeployedToken) //	设置访问路由
	// mux.HandleFunc("/listDeployedTokens", listDeployedTokens)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
