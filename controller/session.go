package controller

import (
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