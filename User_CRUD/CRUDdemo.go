package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./CRUDtest.db")
	checkErr(err)

	fmt.Println("User management")
	for {

		fmt.Println("Please enter an option: ")
		fmt.Println("1. Show all users")
		fmt.Println("2. Create new user")
		fmt.Println("3. Delete user info")
		fmt.Println("4. Update user info")
		fmt.Println("5. Exit")
		fmt.Println("--------------------")
		var n int
		fmt.Scan(&n)

		if n == 1 {
			// search data
			rows, err := db.Query("SELECT * FROM userinfo")
			checkErr(err)

			for rows.Next() {
				var uid int
				var username string
				var password int
				var usertype int
				err = rows.Scan(&uid, &username, &password, &usertype)
				checkErr(err)
				fmt.Printf("Uid=%d Username=%s Password=%d Usertype=%d \n", uid, username, password, usertype)
			}
		} else if n == 2 {
			// create user
			stmt, err := db.Prepare("INSERT INTO userinfo(username, password, usertype) values(?,?,?)")
			checkErr(err)

			fmt.Println("Please enter user name: ")
			var uName string
			fmt.Scan(&uName)

			fmt.Println("Please enter password: ")
			var uPasswd int
			fmt.Scan(&uPasswd)

			fmt.Println("Please enter user type(0 for customer, 1 for administrator): ")
			var uType int
			fmt.Scan(&uType)

			res, err := stmt.Exec(uName, uPasswd, uType)
			checkErr(err)

			id, err := res.LastInsertId()
			checkErr(err)

			fmt.Printf("The new Uid is: %d \n", id)
		} else if n == 3 {
			// delete user
			stmt, err := db.Prepare("delete from userinfo where uid=?")
			checkErr(err)

			fmt.Println("Please enter the Uid you want to delete: ")
			var uid int
			fmt.Scan(&uid)

			res, err := stmt.Exec(uid)
			checkErr(err)

			affect, err := res.RowsAffected()
			checkErr(err)

			fmt.Printf("User(s) deleted:%d \n", affect)
		} else if n == 4 {
			// update data
			stmt, err := db.Prepare("update userinfo set password=? where uid=?")
			checkErr(err)

			fmt.Println("Please enter the Uid: ")
			var id int
			fmt.Scan(&id)

			fmt.Println("Please enter the new password: ")
			var uPasswd int
			fmt.Scan(&uPasswd)

			res, err := stmt.Exec(uPasswd, id)
			checkErr(err)

			affect, err := res.RowsAffected()
			checkErr(err)

			fmt.Printf("User(s) updated:%d \n", affect)
		} else if n == 5 {
			db.Close()
			break
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
