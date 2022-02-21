package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func sign_up(username string, password string) string {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	checkErr(err)

	// see if the username already exists
	rows, err := db.Query("SELECT username,password FROM userinfo where username=" + username)

	if rows.Next() {
		return "username already exist"
	} else {
		stmt, err := db.Prepare("INSERT INTO userinfo(username, password) values(?,?)")
		checkErr(err)

		res, err := stmt.Exec(username, password)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		return string(id) + "successs"
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sign_in(username string, password string) string {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	checkErr(err)
	rows, err := db.Query("SELECT username,password FROM userinfo where username=" + username)
	checkErr(err)
	for rows.Next() {
		var temp1 string
		var temp2 string
		err = rows.Scan(&temp1, &temp2)
		checkErr(err)
		fmt.Println(temp1)
		fmt.Println(temp2)
		if temp2 != password {
			return "Wrong password"
		} else {
			return "Success"
		}
	}
	return "Username Doesnt exist"
}

func main() {
	message := sign_in("12317312", "dgadsfasf")
	fmt.Println(message)
}
