package hoge

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var instance *xorm.Engine

func GetInstance() *xorm.Engine {
	// なければ生成
	if instance == nil {
		engine, err := xorm.NewEngine("mysql", "game:game@tcp(localhost:3306)/game_test?charset=utf8")
		checkErr(err, "sql.Open failed")

		instance = engine

		log.Println("instance make!!")
	}
	return instance
}

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
