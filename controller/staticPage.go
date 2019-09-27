package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func IndexGet(c *gin.Context) {
	session := sessions.Default(c)
	alive := session.Get("alive")
	if alive != nil {
		c.Redirect(http.StatusMovedPermanently, "/home")
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}