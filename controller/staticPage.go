package controller

import (
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
	c.HTML(http.StatusOK, "home.html", gin.H{})
}