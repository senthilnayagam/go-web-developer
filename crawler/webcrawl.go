package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request, _ map[string]string) {
    fmt.Fprint(w, "Welcome!\n")
}

func Get(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	response, err := http.Get("http://" + vars["url"])
	_ =  err

	  if err != nil {
	        fmt.Printf("%s", err)
	       // os.Exit(1)
	    } else {
	        defer response.Body.Close()
	        contents, err := ioutil.ReadAll(response.Body)
	        if err != nil {
	            fmt.Printf("%s", err)
	          //  os.Exit(1)
	        }
	        // fmt.Printf("%s\n", string(contents))
			fmt.Fprintf(w,"%s", contents)
	    }
   // fmt.Fprintf(w, "hello, %s\n", vars["url"])
	 
}

func Post(w http.ResponseWriter, r *http.Request, vars map[string]string) {
    fmt.Fprintf(w, "hello, %s\n", vars["url"])
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/Get/:url", Get)
	router.GET("/Post/:url", Post)

    log.Fatal(http.ListenAndServe(":4040", router))
}