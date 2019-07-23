package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}