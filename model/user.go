package model

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
	"log"
)

type User struct {
	UserID		int		`db:"id"`
	Name		string	`db:"name"`
	NickName	string	`db:"nickName"`
	Password	string	`db:"pass"`
}

func CreateUser(name string, nickName string, password string) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO user(name, nickName, pass) values($1, $2, $3)", name, nickName, password)
	tx.Commit()
}

func SearchUser(name string) (user User) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	users := []User{}

	err = db.Select(&users, "select * from user where name=?", name)
	if err != nil {
		log.Println(err)
	}

	return users[0]
}