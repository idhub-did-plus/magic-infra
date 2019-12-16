// login_controller
package controller

import (
	"magic-infra/component"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func recover(identity string, timestamp string, claim string, signed string) string {
	return identity
}
func hasClaim(identity string, key string, value string) bool {
	claim, err := component.ClaimService.GetClaim(key)
	if err != nil {
		return false

	}

	_ = claim
	return true
}
func Login(c *gin.Context) {
	action := c.Query("action")
	if action == "reentry" {
		session := sessions.Default(c)
		claim := session.Get("claim")
		if nil != claim {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"claim":   claim,
				"message": "login successful!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "invalid signature!",
			})
		}
		return
	}
	address, timestamp, claim, signed := c.Query("identity"), c.Query("timestamp"), c.Query("claim"), c.Query("signature")
	recaddress := recover(address, timestamp, claim, signed)
	if recaddress != address {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "invalid signature!",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("claim", claim)
	session.Save()
	cl := session.Get("claim").(string)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"claim":   cl,
		"message": "login successful!",
	})

}

//https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0
