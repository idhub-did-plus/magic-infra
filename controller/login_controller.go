// login_controller
package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Delete("tizi365")
		session.Save()
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "saved!",
	})

}
