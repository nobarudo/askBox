package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	c "askBox/controller"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("views/*")

	router.GET("/", indexGet)
	router.GET("/login", c.LoginGet)
	router.GET("/signup", c.SignupGet)
	router.Run(":8080")
}

func indexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}