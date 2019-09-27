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
	password := passwordHash(c.Request.Form["password"][0])
	model.CreateUser(name, nickname, password)

	c.HTML(http.StatusOK, "createUser.html", gin.H{})
}

func Login(c *gin.Context) {
	c.Request.ParseForm()
	password := c.Request.Form["password"][0]
	user := model.SearchUser(c.Request.Form["userID"][0])

	result := passwordVerify(user.Password, password)
	if result == true {
		CreateSession(c, user)
		c.Redirect(http.StatusMovedPermanently, "/home")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func Logout(c *gin.Context) {
	log.Println("logout")
	DeleteSession(c)
	IndexGet(c)
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

func UserBox(c *gin.Context) {
	name := c.Param("user")
	c.HTML(http.StatusOK, "newQuestion.html", gin.H{"name": name})
}