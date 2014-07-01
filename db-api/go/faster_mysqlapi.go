package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime"
)

/*

bytesOfJSON, _ := json.Marshal(myStructOrSlice)

*/

var db *sql.DB

var id string
var name string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use max cores/cpus available from OS
	// err :=  0;
	// database connection
	var err error
	db, err = sql.Open("mysql", "root:root@tcp([127.0.0.1]:3306)/dbapi?autocommit=true&charset=utf8")
	if err != nil {
		fmt.Println("error conecting to DB")
	} else {
		fmt.Println("connected")
	}
	defer db.Close()

	// connection pool
	db.SetMaxIdleConns(100)

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}

	/*
	   mux := pat.New()
	   mux.Get("/mysql/:name/:id", http.HandlerFunc(profile))

	   http.Handle("/", mux)

	   log.Println("Listening port:3000 ")
	   http.ListenAndServe(":3000", nil)

	*/

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/mysql/:name/:id", Hello)

	log.Fatal(http.ListenAndServe(":12345", router))

}

func Index(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	id = vars["id"]
	table := vars["name"]
	_ = table

	rows, err := db.Query("SELECT * FROM car where id=?", id)
	if err != nil {
		fmt.Printf("error query: %s", err.Error())
		return
	}
	defer rows.Close()

	// print elements
	for rows.Next() {
		rows.Scan(&id, &name)
		//     fmt.Printf("id = %d, name = %s\n", id, name)
	}

	fmt.Fprintf(w, "{[\"name\":\""+name+"\",\"id\":"+id+"]}")
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	//  tablename := params.Get(":name")
	//  fmt.Printf(name)

	idn := params.Get(":id")
	//   fmt.Printf(idn)

	// run a query
	//*
	rows, err := db.Query("SELECT * FROM car where id=?", idn)
	if err != nil {
		fmt.Printf("error query: %s", err.Error())
		return
	}
	defer rows.Close()

	// print elements
	for rows.Next() {
		rows.Scan(&id, &name)
		//     fmt.Printf("id = %d, name = %s\n", id, name)
	}
	//*/

	w.Write([]byte("{[\"name\":\"" + name + "\",\"id\":" + idn + "]}")) // hand coded json
}
