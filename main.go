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
	r.Any("/login", controller.Login)

	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
