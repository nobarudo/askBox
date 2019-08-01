package controller

import (
	"github.com/gin-contrib/sessions"
	"net/http"
	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func HomeGet(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("alive") == true {
		name := session.Get("nickName")
		c.HTML(http.StatusOK, "home.html", gin.H{"userName": name})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}