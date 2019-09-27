package controller

import (
	"log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"askBox/model"
)

func CreateSession(c *gin.Context, user model.User) {
	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("userID", user.UserID)
	session.Set("nickName", user.NickName)
	session.Save()
}

func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	log.Println("clear session")
	session.Save()
}