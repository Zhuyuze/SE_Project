package login_register

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type T struct {
	username string `json:"username"`
	password string `json:"password"`
}

func register(username string, password string) string {

	// connect database
	cfg := mysql.Config{
		User:   "root",
		Passwd: "UF233333",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query("SELECT username from t where username=" + username)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results.Next() {
		return "already exist!"
	}

	_, err1 := db.Exec("INSERT INTO t (username,password) VALUES (?, ?)", username, password)
	if err1 != nil {
		return "Went wrong"
	}

	return "Success"
}

func login(username string, password string) string {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "UF233333",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query("SELECT * from t where username=" + username + "and" + "password=" + password)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if results.Next() {
		return "Success"
	}
	return "Wrong"
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "UF233333",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "test",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query("SELECT * FROM t WHERE username='5'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println(results.Next())
	for results.Next() {
		var t T
		// for each row, scan the result into our tag composite object
		err = results.Scan(&t.username, &t.password)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		println(t.username)
		println(t.password)
	}
}
