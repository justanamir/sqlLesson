package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:5227@tcp(127.0.0.1:3307)/test")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY NOT NULL AUTO_INCREMENT, username TEXT, email TEXT, password TEXT)")

	insert, err := db.Query("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", "jun", "jun@g.com", "1234")
	if err != nil {
		log.Fatalln(err)
	}

	defer insert.Close()

	rows, err := db.Query("SELECT	* FROM user")
	if err != nil {
		log.Fatalln(err)
	}

	var id int
	var username string
	var email string
	var password string

	for rows.Next() {
		rows.Scan(&id, &username, &email, &password)
		fmt.Println(id, username, email, password)
	}
}
