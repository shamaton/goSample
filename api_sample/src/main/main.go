package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type TestInput struct {
	Name string
	Age  int
}

type TestOutput struct {
	Name string
	Age  int
}

func test(w rest.ResponseWriter, r *rest.Request) {
	input := TestInput{}
	r.DecodeJsonPayload(&input)
	fmt.Println(input)
	output := TestOutput{}

	output.Name = input.Name + "さん"
	output.Age = input.Age
	w.WriteJson(output)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	/*
		api.Use(&rest.CorsMiddleware{
			RejectNonCorsRequests: false,
			OriginValidator: func(origin string, request *rest.Request) bool {
				return true
			},
			AllowedMethods: []string{"GET", "POST", "PUT"},
			AllowedHeaders: []string{
				"Accept", "Content-Type", "X-Custom-Header", "Origin"},
			AccessControlAllowCredentials: true,
			AccessControlMaxAge:           3600,
		})
	*/
	router, err := rest.MakeRouter(
		rest.Post("/test", test),
	)

	// 存在しないルート時
	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9999", api.MakeHandler()))
}

/*
package main

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"sample"

	_ "github.com/go-sql-driver/mysql"
)

func selectSQL(db *sql.DB) string {
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var teststr string

	for rows.Next() {
		var id int
		var name string
		var score int
		err := rows.Scan(&id, &name, &score)
		if err != nil {
			panic(err)
		}
		teststr = fmt.Sprintf("id:%d\tname:%s\tscore:%d\n", id, name, score)
	}
	return teststr
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	title := html.EscapeString(req.URL.Path[1:])
	strpost := req.FormValue("strpost")
	strget := req.FormValue("strget")

	db, err := sql.Open("mysql", "game:game@tcp(localhost:3306)/game_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(sample.Message) // hello world
	dbstr := selectSQL(db)

	output := `
<html>
  <head>
   <title>` + title + `</title>
  </head>
  <body>
  <h1>post/getのテスト</h1>
     <h2>post</h2>
        post=` + html.EscapeString(strpost) + `</br>
     <h2>get</h2>
        get=` + html.EscapeString(strget) + `</br>
     <h2>db</h2>
        db=` + html.EscapeString(dbstr) + `</br>
  </body>
</html>
`
	fmt.Fprintf(w, "%s", output)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8080", nil)

}
*/
