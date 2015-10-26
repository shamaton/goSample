package controller

import (
	"fmt"
	"hoge"
	"log"
	"model"

	"time"

	"golang.org/x/net/context"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// sql_builderとして扱う
)

func Test(w rest.ResponseWriter, r *rest.Request) {
	log.Println("hogehogfukga")

	ctx := r.Env["context"].(context.Context)

	// データをselect
	test := model.Find(3)
	log.Println(test)

	// use redis
	redisTest(ctx)

	// データをupdate
	var h *xorm.Engine
	h = hoge.GetDBShardConnection("user", 1)

	session := h.NewSession()
	defer session.Close()

	var err error

	err = session.Begin()
	if checkErr(err, "begin error") {
		w.WriteJson(err)
		return
	}

	var u []model.User
	err = session.Where("id = ?", 3).ForUpdate().Find(&u)
	if checkErr(err, "user not found") {
		w.WriteJson(err)
		return
	}

	user := u[0]
	user.Score += 1

	time.Sleep(6 * time.Second)

	//res, e := session.Id(user.Id).Cols("score").Update(&user) // 単一 PK
	_, err = session.Id(core.PK{user.Id, user.Name}).Update(&user) // 複合PK
	if checkErr(err, "update error") {
		w.WriteJson(err)
		return
	}

	err = session.Commit()
	if checkErr(err, "commit error") {
		w.WriteJson(err)
		return
	}

	w.WriteJson(user)
}

func redisTest(ctx context.Context) {

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
}

// エラー表示
func checkErr(err error, msg string) bool {
	if err != nil {
		log.Println(msg, err)
		return true
	}
	return false
}
