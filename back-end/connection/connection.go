package main

import (
	"net/http"

	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello ")
	})

	r.GET("/haha", func(c *gin.Context) {
		c.String(http.StatusOK, "hahahahha")
	})

	r.GET("/sign_in/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		message := sign_in(username, password)
		c.String(http.StatusOK, "Hello %s", message)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

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
