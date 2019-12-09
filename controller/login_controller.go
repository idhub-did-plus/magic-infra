// login_controller
package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func recover(identity string, timestamp string, claim string, signed string) string {
	return identity
}
func Login(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"claim": claim,
		"message": "login successful!",
	})

}

//https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0
