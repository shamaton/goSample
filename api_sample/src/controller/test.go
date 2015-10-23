package controller

import (
	"log"
	"model"

	"gopkg.in/gorp.v1"

	"golang.org/x/net/context"

	"github.com/ant0ine/go-json-rest/rest" // sql_builderとして扱う
)

func Test(w rest.ResponseWriter, r *rest.Request /*ctx context.Context*/) {
	log.Println("hogehogfukga")

	ctx := r.Env["context"].(context.Context)
	str := ctx.Value("test").(string)
	log.Println(str)

	db := ctx.Value("DB").(*gorp.DbMap)
	db.AddTableWithName(model.User{}, "users").SetKeys(false, "Id") // これがないと無理

	// データをselect
	user := model.Find(ctx, db, 3)
	user.Hoge()
	log.Println("fjdksal;gjidopajio")

	// データをupdate : for updateで呼ぶべき
	user.Score += 1
	log.Println(user)
	tx, errr := db.Begin()
	checkErr(errr, "tx error!")
	res, e := tx.Update(&user)
	log.Println(res)
	checkErr(e, "")
	ee := tx.Commit()
	checkErr(ee, "commit error!!")

	tx.Commit()
	w.WriteJson(user)
}

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
