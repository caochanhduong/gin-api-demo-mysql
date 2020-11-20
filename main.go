package main

import (
	// "github.com/gin-gonic/gin"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User is a struct containing user info
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// c := gin.Default()
	// c.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "duongcc",
	// 	})
	// })
	// c.Run()

	db, err := sql.Open("mysql", "root:Bikhungha1!@tcp(127.0.0.1:3306)/gin_demo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("connected to mysql")
	// // test insert query
	// insert, err := db.Query("INSERT INTO user VALUES (4, 'hangltt3')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// // test select all query
	// results, err := db.Query("SELECT * FROM user")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// users := make([]User, 0)
	// for results.Next() {
	// 	var user User
	// 	err := results.Scan(&user.ID, &user.Name)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	users = append(users, User{
	// 		ID:   user.ID,
	// 		Name: user.Name,
	// 	})
	// }
	// fmt.Println("--------------------result")
	// fmt.Println(users)

	// // test select one by id
	// var user User
	// err = db.QueryRow("SELECT * FROM user WHERE id = ?", 2).Scan(&user.ID, &user.Name)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(user)

	// // test delete
	// result, err := db.Exec("DELETE FROM user WHERE id = ?", 4)
	// if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	num, err := result.RowsAffected()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	} else {
	// 		fmt.Println("Num rows affected: ", num)
	// 	}
	// }

	// test update by id
	result, err := db.Exec("UPDATE user SET name = ? WHERE id = ?", "hahs", 3)
	if err != nil {
		panic(err.Error())
	} else {
		num, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Println("Num rows affected: ", num)
		}
	}
}
