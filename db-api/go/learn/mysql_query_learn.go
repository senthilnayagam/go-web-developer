package main

import (
    "database/sql"
    "fmt"
//    "github.com/gorilla/mux"
    _ "github.com/go-sql-driver/mysql"
    "log"
//    "net/http"
)

var db *sql.DB // global variable to share it between main and the HTTP handler

func main() {
    fmt.Println("starting up")

    var err error
//    db, err = sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/dbapi") // this does not really open a new connection
    db, err = sql.Open("mysql", "root:root@/dbapi") //	
	
	
	 
    if err != nil {
        log.Fatalf("Error on initializing database connection: %s", err.Error())
    }

    db.SetMaxIdleConns(100)

    err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
    if err != nil {
        log.Fatalf("Error on opening database connection: %s", err.Error())
    }



    var msg string
    err2 := db.QueryRow("SELECT * FROM car WHERE id=1").Scan(&msg)
    if err2 != nil {
        fmt.Println( "Database Error!")
		  fmt.Println( msg)
    } else {
        fmt.Println( msg)
    }

//    r := mux.NewRouter()
//    r.HandleFunc("/", HomeHandler)

//    http.Handle("/", r)
//    http.ListenAndServe(":8080", nil)
}

/*
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    var msg string
    err := db.QueryRow("SELECT * FROM car WHERE id=?", "1").Scan(&msg)
    if err != nil {
        fmt.Fprintf(w, "Database Error!")
    } else {
        fmt.Fprintf(w, msg)
    }
}
*/