package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

type User struct {
	Id    int32
	Name  string
	Score int32
	//Hoge int32   //`db:"score, [primarykey, autoincrement]"` 変数名とカラム名が異なる場合JSON的に書ける
}

func initDb() *gorp.DbMap {
	// MySQLへのハンドラ
	db, err := sql.Open("mysql", "game:game@tcp(localhost:3306)/game_test")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}

func main() {
	// 初期化
	dbmap := initDb()
	defer dbmap.Db.Close()

	// データをselect
	// パターン 1
	dbmap.AddTableWithName(User{}, "users").SetKeys(false, "Id")
	obj, err := dbmap.Get(User{}, 1)
	checkErr(err, "not found data!")

	u := obj.(*User)
	log.Printf("id : %d, name %s, score %d", u.Id, u.Name, u.Score)

	// パターン 2
	var user User // user := User{}
	err2 := dbmap.SelectOne(&user, "select * from users where id = 1")
	checkErr(err2, "not found data!")
	log.Printf("id : %d, name %s, score %d", user.Id, user.Name, user.Score)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
