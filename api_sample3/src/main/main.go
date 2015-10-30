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

func Custom() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// set global context
		c.Set("gContext", ctx)

		// リクエスト前処理

		c.Next()

		// リクエスト後処理
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// context
	ctx = context.Background()

	// db
	hoge.BuildInstances()

	// redis
	redis_pool := newPool()
	ctx = context.WithValue(ctx, "redis", redis_pool)

	router := gin.Default()a
	router.Use(Custom())
	// make route
	router.POST("/test", controller.Test)
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
