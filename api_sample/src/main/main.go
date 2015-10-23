package main

import (
	"controller"
	"database/sql"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// global
var (
	ctx context.Context
)

/*
func SetContext(handler http.Handler) http.Handler {
    fn := func(w rest.ResponseWriter, r *rest.Request){
        //web.Cにcontext.Contextを埋め込むサードパーティパッケージがあるのでそれを使います。
        gtx.Set(c, ctx)

        r.Env["c"] = ctx

        h.ServeHTTP(w, r)
    }
    return rest.HandlerFunc(fn)
}
*/

func initDb() *gorp.DbMap {
	// MySQLへのハンドラ
	db, err := sql.Open("mysql", "game:game@tcp(localhost:3306)/game_test")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}

// redis ConnectionPooling
func newPool() *redis.Pool {
	hostname := "127.0.0.1"
	port := "6379"

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hostname+":"+port)

			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func baseHandlerFunc(handler func(w rest.ResponseWriter, r *rest.Request)) rest.HandlerFunc {
	return baseHandler(rest.HandlerFunc(handler))
}

func baseHandler(handler func(w rest.ResponseWriter, r *rest.Request)) rest.HandlerFunc {
	return rest.HandlerFunc(func(w rest.ResponseWriter, r *rest.Request) {
		// common
		log.Println(r.URL, r.Method)
		r.Env["context"] = ctx
		handler(w, r)
	})
}

func main() {
	// context
	ctx = context.Background()
	db := initDb()
	ctx = context.WithValue(ctx, "test", "aaabbbccc")
	ctx = context.WithValue(ctx, "DB", db)

	// redis
	redis_pool := newPool()
	ctx = context.WithValue(ctx, "redis", redis_pool)

	str := ctx.Value("test")
	log.Println(str)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/test", baseHandlerFunc(controller.Test)),
	)

	// 存在しないルート時
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
