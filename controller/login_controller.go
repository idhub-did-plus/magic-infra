// login_controller
package controller

import (
	"magic-infra/middleware"
	"net/http"

	"github.com/google/uuid"

	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username, _ := c.PostForm("username"), c.PostForm("password")
	// Authentication
	// blahblah...

	// Generate random session id
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	// Store current subject in cache
	middleware.GlobalCache.Set(sessionId, []byte(username))
	// Send cache key back to client in cookie
	c.SetCookie("current_subject", sessionId, 30*60, "/resource", "", false, true)
	//c.JSON(200, component.RestResponse{Code: 1, Message: username + " logged in successfully"})
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

//https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0
