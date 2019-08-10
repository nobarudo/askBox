package model

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id			int		`db:"id"`
	TargetUser	string	`db:"targetUser"`
	Text		string	`db:"text"`
	Reply		*string	`db:"reply"`
}

func ShowQuestionsList(userName string) (questions []Question) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	questions = []Question{}
	err = db.Select(&questions, "select * from question where targetUser=?", userName)
	if err != nil {
		log.Println(err)
	}
	log.Println("model", questions)
	return questions
}

func AddQuestion(userName string, text string) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("insert前 name:",userName, "text:", text)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO question(targetUser, text) values($1, $2)", userName, text)
	tx.Commit()
}

func ShowQuestion(id int) (questions []Question){
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	questions = []Question{}
	err = db.Select(&questions, "select * from question where id=?", id)
	if err != nil {
		log.Println(err)
	}
	log.Println(questions)
	return questions
}

func UpdateReply(reply string, id int) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("update前 id:",id, "reply:", reply)
	tx := db.MustBegin()
	tx.MustExec("UPDATE question set reply=$1 where id=$2 ", reply, id)
	tx.Commit()
}