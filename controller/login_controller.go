// login_controller
package controller

import (
	"magic-infra/contract"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func recover(identity string, timestamp string, claim string, signed string) string {
	return identity
}
func hasClaim(identity string, key string, value string) bool {
	has := contract.ContractService.HasClaim(identity, key, value)
	return has

	// var vhash = utils.Sha3(value)
	// bytes.Equal(vhash[:], claim[:])
	// return true
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
				"message": "not logged!",
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
func Logout(c *gin.Context) {

	session := sessions.Default(c)
	session.Delete("claim")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "logout success!",
	})

}

//https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0
