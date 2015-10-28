package main

import (
	"controller"
	"hoge"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	"github.com/garyburd/redigo/redis"
)

// global
var (
	ctx context.Context
)

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

func baseHandlerFunc(handler func(c *gin.Context)) gin.HandlerFunc {
	return baseHandler(gin.HandlerFunc(handler))
}

func baseHandler(handler func(c *gin.Context)) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// common
		// log.Println(c)
		c.Set("gContext", ctx)
		handler(c)
	})
}

/*
func SetGontext(c *gin.Context) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL, r.Method)
		c.Set("gContext", ctx)
		c.
		h.ServeHTTP(w, r)
	}
	return gin.HandlerFunc(fn)
}
*/

func main() {
	// context
	ctx = context.Background()

	// db
	hoge.BuildInstances()

	// redis
	redis_pool := newPool()
	ctx = context.WithValue(ctx, "redis", redis_pool)

	router := gin.Default()
	//router.Use(SetGontext)
	// make route
	router.POST("/test", baseHandlerFunc(controller.Test))
	//router.POST("/test", controller.Test)

	err := router.Run(":9999")

	// 存在しないルート時
	if err != nil {
		log.Fatal(err)
	}
}

// エラー表示
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
