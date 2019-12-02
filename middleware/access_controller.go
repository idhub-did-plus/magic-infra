package middleware

import (
	"net/http"

	"fmt"
	"time"

	"github.com/allegro/bigcache"
	_ "github.com/casbin/casbin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var GlobalCache *bigcache.BigCache

func init() {
	var err error
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute)) // Set expire time to 30 mins
	if err != nil {
		panic(fmt.Sprintf("failed to initialize cahce: %v", err))
	}
}
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
