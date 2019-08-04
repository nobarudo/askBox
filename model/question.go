package model

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/jmoiron/sqlx"
)

type Question struct {
	Id			int		`db:"id"`
	TargetUser	string	`db:"targetUser"`
	Text		string	`db:"question"`
}

func ShowQuestions(userName string) (questions []Question) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	questions = []Question{}
	err = db.Select(&questions, "select * from question")
	if err != nil {
		log.Println(err)
	}
	log.Println(questions)
	return questions
}

func AddQuestion(userName string, text string) {
	db, err := sqlx.Connect("sqlite3", "./db/askbox.db")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("insert前 name:",userName, "text:", text)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO question(targetUser, question) values($1, $2)", userName, text)
	tx.Commit()
}