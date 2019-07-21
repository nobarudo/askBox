package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/public", "./public")
	router.LoadHTMLGlob("views/*")

	router.GET("/", indexGet)
	router.Run(":8080")
}

func indexGet(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}