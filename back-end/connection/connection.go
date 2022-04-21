package main

import (
	"net/http"

	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

var global_username string

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
	r.GET("/get_info", func(c *gin.Context) {
		name := user_info()
		c.String(http.StatusOK, name)
	})
	r.GET("/sign_in/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		message := sign_in(username, password)
		c.String(http.StatusOK, "Hello %s", message)
	})

	r.GET("/sign_up/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		message := sign_up(username, password)
		c.String(http.StatusOK, "%s", message)
	})
	r.GET("/order", func(c *gin.Context) {

		message := order()
		c.String(http.StatusOK, "%s", message)
	})

	r.GET("/get_num", func(c *gin.Context) {
		message := prev_orders()
		c.String(http.StatusOK, "%d", message)
	})
	r.GET("/order_food/:username", func(c *gin.Context) {
		username := c.Param("username")
		message := order_food(username)
		c.String(http.StatusOK, "%d", message)
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func sign_up(username string, password string) string {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	checkErr(err)

	// see if the username already exists
	rows, err := db.Query("SELECT * FROM userinfo where username=" + "'" + username + "'")
	//checkErr(err)
	fmt.Println(rows.Next())
	if err != nil {
		fmt.Println("database success")

	}

	if rows.Next() {
		return "username already exists"

	} else {
		stmt, err := db.Prepare("INSERT INTO userinfo(username, password) values(?,?)")
		checkErr(err)

		res, err := stmt.Exec(username, password)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
		return "successs"
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
	rows, err := db.Query("SELECT username,password FROM userinfo where username=" + "'" + username + "'")
	checkErr(err)
	if !rows.Next() {
		return "Wrong username or password"
	} else {
		var temp1 string
		var temp2 string
		err = rows.Scan(&temp1, &temp2)
		checkErr(err)
		fmt.Println(temp1)
		fmt.Println(temp2)
		if temp2 != password {
			return "Wrong password"
		} else {
			global_username = username
			return "Success"
		}
	}
	return "Username Doesnt exist"
}

func order() string {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	//db.SetMaxOpenConns(1)
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO orders(username, food_name,quantity) values(?,?,?)")
	checkErr(err)
	//time.Sleep(5 * time.Second)

	res, err := stmt.Exec(global_username, "Ribs with brown sause", 1)
	checkErr(err)
	res, err = stmt.Exec(global_username, "Brased Pork in brown sause", 1)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	return "successfully ordered"
}

func prev_orders() int {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	//db.SetMaxOpenConns(1)
	checkErr(err)
	rows, err := db.Query("SELECT quantity from orders where username='" + global_username + "'")
	checkErr(err)
	count := 0
	if !rows.Next() {
		return 0
	} else {
		rows, err := db.Query("SELECT quantity from orders where username='" + global_username + "'")
		checkErr(err)
		for rows.Next() {
			count = count + 1
		}
	}
	return count
}

func order_food(username string) string {
	db, err := sql.Open("sqlite3", "CRUDtest.db")
	//db.SetMaxOpenConns(1)
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO orders(username, food_name,quantity) values(?,?,?)")
	res, err := stmt.Exec(username, "Spareribs with brown sauce", 1)
	stmt, err = db.Prepare("INSERT INTO orders(username, food_name,quantity) values(?,?,?)")
	res, err = stmt.Exec(username, "Braised pork in brown sauce", 1)
	checkErr(err)
	fmt.Println(res)
	return "Success"
}

func user_info() string {
	if global_username == "" {
		return "please login first"
	}
	return global_username
}
