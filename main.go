// magic-infra project main.go
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"magic-infra/controller"
	"magic-infra/middleware"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(controller.CorsMiddleware())
	r.Use(middleware.AclMiddleware())

	r.POST("/saveDeployedToken", controller.SaveDeployedToken)
	r.GET("/listDeployedTokens", controller.ListDeployedTokens)
	r.Any("/login", controller.Login)

	r.Run() // 在 0.0.0.0:8080 上监听并服务
}
