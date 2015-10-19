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
