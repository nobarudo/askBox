package controller

import (
	"golang.org/x/crypto/bcrypt"
	"askBox/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func Signup(c *gin.Context) {
	c.Request.ParseForm()
	name := c.Request.Form["userID"][0]
	nickname := c.Request.Form["userName"][0]
	pass := c.Request.Form["password"][0]
	password := passwordHash(pass)
	log.Println("userID:", name, "userName", nickname,"password", password)
	model.CreateUser(name, nickname, password)

	c.HTML(http.StatusOK, "createUser.html", gin.H{})
}

func passwordHash(pass string) (string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}