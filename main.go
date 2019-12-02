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

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
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
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
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

func saveDeployedToken(c *gin.Context) {
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

func main() {
	r := gin.Default()
	r.Use(CorsMiddleware())

	store := cookie.NewStore([]byte("secret11111"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/saveDeployedToken", saveDeployedToken)
	r.GET("/listDeployedTokens", listDeployedTokens)
	r.GET("/hello", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)

		// 通过session.Get读取session值
		// session是键值对格式数据，因此需要通过key查询数据
		if session.Get("hello") != "world" {
			// 设置session数据
			session.Set("hello", "world")
			// 删除session数据
			session.Delete("tizi365")
			// 保存session数据
			session.Save()
			// 删除整个session
			// session.Clear()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务
	// mux := http.NewServeMux()
	// mux.HandleFunc("/saveDeployedToken", saveDeployedToken) //	设置访问路由
	// mux.HandleFunc("/listDeployedTokens", listDeployedTokens)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
