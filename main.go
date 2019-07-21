package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", indexGet)
	router.Run(":8080")
}

func indexGet(c *gin.Context) {
	c.String(http.StatusOK, "Helo World")
}