package model

import (
	"log"

	"golang.org/x/net/context"

	"gopkg.in/gorp.v1"
)

type User struct {
	Id    int32
	Name  string
	Score int32
	//Hoge int32   //`db:"score, [primarykey, autoincrement]"` 変数名とカラム名が異なる場合JSON的に書ける
}

type baseModel struct {
	*User
	*Base
}

func Find(ctx context.Context, db *gorp.DbMap, userId int) User {
	//db := ctx.Value("DB").(*gorp.DbMap)
	var test baseModel
	//test.SetText("test text")
	test.Hoge()

	// データをselect
	var user User // user := User{}
	err := db.SelectOne(&user, "select * from users where id = ?", userId)
	checkErr(err, "not found data!")
	return user

}

func (u *User) Hoge() {
	log.Println("hogeeeeeeeeeeeee")
}

// データをupdate : for updateで呼ぶべき
/*
   user.Score += 1
   log.Println(user)
   tx, errr := dbmap.Begin()
   checkErr(errr, "tx error!")
   res, e := tx.Update(&user)
   log.Println(res)
   checkErr(e, "")
   ee := tx.Commit()
   checkErr(ee, "commit error!!")

   tx.Commit()
*/

/*

  // パターン 1
    dbmap.AddTableWithName(User{}, "users").SetKeys(false, "Id")
    obj, err := dbmap.Get(User{}, 1)
    checkErr(err, "not found data!")

    u := obj.(*User)
    log.Printf("id : %d, name %s, score %d", u.Id, u.Name, u.Score)

  // パターン 2 (こちらの場合はSQLを書くのでAddTable不要)
  var user User // user := User{}
  err2 := dbmap.SelectOne(&user, "select * from users where id = 2")
  checkErr(err2, "not found data!")
  log.Printf("id : %d, name %s, score %d", user.Id, user.Name, user.Score)

  // パターン 3 (squirrelでSQL生成)
  sb := sql_builder.Select("*").From("users")
  sb = sb.Where(sql_builder.Eq{"id": 3})
  sql, args, sql_err := sb.ToSql()
  log.Println(sql)

  checkErr(sql_err, "SQL error!!")

  var user3 User // user := User{}
  err3 := dbmap.SelectOne(&user3, sql, args[0])
  checkErr(err3, "not found data!")
  log.Printf("id : %d, name %s, score %d", user3.Id, user3.Name, user3.Score)
*/

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
