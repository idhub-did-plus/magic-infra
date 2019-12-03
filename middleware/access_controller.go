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
		if c.FullPath() == "/login" {
			c.Next()
			return
		}
		session := sessions.Default(c)
		if nil == session.Get("claim") {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		claim := session.Get("claim").(string) //从session取出当前用户选中的角色
		if !filter(c, claim) {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "access denied!",
			})
			return
		}
		c.Next()
	}

}
func filter(c *gin.Context, claim string) bool {
	return true
}
