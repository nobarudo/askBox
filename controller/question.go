package controller

import (
	"log"
	"askBox/model"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetQuestionList(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("alive") == true {
		name := session.Get("nickName")
		question := model.ShowQuestions(name.(string))
		log.Println(question)
		c.HTML(http.StatusOK, "home.html", gin.H{"userName": name, "question": question })
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}