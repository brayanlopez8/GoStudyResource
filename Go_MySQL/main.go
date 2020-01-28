package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//User model user
type User struct {
	Name string `json:"name"`
}

func main() {
	// sampleImsert()
	sampleSelect()
}

func sampleSelect() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Seccessfully connected to database ")

	results, err := db.Query("Select name from user")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

}

func sampleImsert() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Seccessfully connected to database ")

	insert, err := db.Query("insert into user values ('Brayan')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully inserted into user tables")
}
