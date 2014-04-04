package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var id int
var name string

func main() {
    // go get github.com/go-sql-driver/mysql

    // database connection
    db, err := sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/dbapi?autocommit=true&charset=utf8")
    if err != nil {
        fmt.Println("error conecting to DB")
    } else {
        fmt.Println("connected")
    }
    defer db.Close()
	
	
	
    // run a query
      rows, err := db.Query("SELECT * FROM car")
      if err != nil {
          fmt.Printf("error query: %s", err.Error())
          return
      }
      defer rows.Close()
   
      // print elements
      for rows.Next() {
          rows.Scan(&id, &name)
          fmt.Printf("id = %d, name = %s\n", id, name)
      }



      // run a query
        rows, err = db.Query("SELECT * FROM car where id=1")
        if err != nil {
            fmt.Printf("error query: %s", err.Error())
            return
        }
        defer rows.Close()
   
        // print elements
        for rows.Next() {
            rows.Scan(&id, &name)
            fmt.Printf("id = %d, name = %s\n", id, name)
        }
	
	
	
}