package hoge

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"
)

var instance *gorp.DbMap

func GetInstance() *gorp.DbMap {
	log.Println(instance)
	if instance == nil {
		db, err := sql.Open("mysql", "game:game@tcp(localhost:3306)/game_test")
		checkErr(err, "sql.Open failed")

		log.Println("instance make!!")

		// construct a gorp DbMap
		instance = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}
	return instance
}

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
