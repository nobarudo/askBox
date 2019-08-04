package main

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	c "askBox/controller"
)

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	router.Static("/public", "./public")
	router.LoadHTMLGlob("views/*")

	router.GET("/", c.IndexGet)
	router.GET("/login", c.LoginGet)
	router.GET("/logout", c.Logout)
	router.GET("/signup", c.SignupGet)
	router.GET("/home", c.GetQuestionList)
	router.GET("/box/:user", c.UserBox)
	router.POST("/signup", c.Signup)
	router.POST("/login", c.Login)
	router.POST("/box/:user/new", c.NewQuestion)
	router.Run(":8080")
}