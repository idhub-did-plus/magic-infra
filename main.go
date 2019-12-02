// magic-infra project main.go
package main

import (
	"magic-infra/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(controller.CorsMiddleware())

	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/saveDeployedToken", controller.SaveDeployedToken)
	r.GET("/listDeployedTokens", controller.ListDeployedTokens)
	r.GET("/login", controller.Login)
	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Delete("tizi365")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
