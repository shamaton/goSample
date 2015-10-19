package main

import (
	"database/sql"
	"fmt"
	"sample"

	_ "github.com/go-sql-driver/mysql"
)

func selectSQL(db *sql.DB) {
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var score int
		err := rows.Scan(&id, &name, &score)
		if err != nil {
			panic(err)
		}
		fmt.Printf("id:%d\tname:%s\tscore:%d\n", id, name, score)
	}
}

func main() {
	db, err := sql.Open("mysql", "game:game@tcp(localhost:3306)/game_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(sample.Message) // hello world
	selectSQL(db)
}
