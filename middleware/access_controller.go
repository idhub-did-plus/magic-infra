package middleware

import (
	"net/http"

	_ "github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AclMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		if nil == session.Get("current_role_code") {
			//session中取不到当前选中的角色,未登录?
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		current_user_role := session.Get("current_role_code").(string) //从session取出当前用户选中的角色
		sub := current_user_role

		if sub == "" {

			c.Abort()
			return
		}
		c.Next()
	}
}
