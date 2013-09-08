package main

import (
    "os"
    "os/signal"
    "syscall"
    "net/http"
    "log"
    "flag"
    "fmt"
)

var localPort *string = flag.String("p", "3000", "local port")
var publicFolder *string = flag.String("f", "public", "public folder")


func signalCatcher() {
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT)
        <-ch
        log.Println("CTRL-C; exiting")
        os.Exit(0)
}


func main() {
	go signalCatcher()
  flag.Parse()
  fmt.Printf("Listening: %v\nServing: %v\n\n", *localPort, *publicFolder)

// http.Fileserver serves the index.html if it exists and file is not specified
    err := http.ListenAndServe(fmt.Sprint(":",*localPort), http.FileServer(http.Dir(*publicFolder)))
    if err != nil {
        log.Printf("Error running web server for static assets: %v", err)
    }

}


