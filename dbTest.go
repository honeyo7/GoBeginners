package main

import (
	
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

	fmt.Println("Hello World")

	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/demo")

	if err !=nil{
		panic(err.Error())
	}

	fmt.Println("Connected")

	defer db.Close()

}