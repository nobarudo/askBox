package controller

import (
	"strconv"
	"askBox/model"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetQuestionList(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("alive") != nil {
		name := session.Get("nickName")
		question := model.ShowQuestionsList(name.(string))
		c.HTML(http.StatusOK, "home.html", gin.H{"userName": name, "question": question })
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func NewQuestion(c *gin.Context) {
	c.Request.ParseForm()
	name := c.Param("user")
	text := c.Request.Form["question"][0]
	model.AddQuestion(name, text)
	c.HTML(http.StatusOK, "success.html", gin.H{})
}

func GetReplyPage(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	question := model.ShowQuestion(id)
	c.HTML(http.StatusOK, "reply.html", gin.H{"question": question[0]})
}

func Reply(c *gin.Context) {
	c.Request.ParseForm()
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	reply := c.Request.Form["reply"][0]
	model.UpdateReply(reply, id)
	c.Redirect(http.StatusMovedPermanently, "/home")
}