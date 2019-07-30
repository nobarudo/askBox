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

func Login(c *gin.Context) {
	c.Request.ParseForm()
	name := c.Request.Form["userID"][0]
	password := c.Request.Form["password"][0]
	user := model.SearchUser(name)

	result := passwordVerify(user[0].Password, password)
	if result == true {
		c.Redirect(http.StatusMovedPermanently, "/home")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func passwordHash(pass string) (string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(hash)
}

func passwordVerify(hash, pw string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}