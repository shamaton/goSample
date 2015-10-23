package controller

import (
	"fmt"
	"log"
	"model"

	"gopkg.in/gorp.v1"

	"golang.org/x/net/context"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/garyburd/redigo/redis"
	// sql_builderとして扱う
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

	// use redis
	redis_pool := ctx.Value("redis").(*redis.Pool)
	redis_conn := redis_pool.Get()

	_, e2 := redis_conn.Do("SET", "message", "this is value")
	if e2 != nil {
		log.Fatalln("set message", e2)
	}
	s, err := redis.String(redis_conn.Do("GET", "message"))
	if err != nil {
		log.Fatalln("get message", err)
	}
	fmt.Printf("%#v\n", s)

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
