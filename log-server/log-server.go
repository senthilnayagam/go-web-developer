package main

import (
    "os"
    "os/exec"
    "os/signal"
    "syscall"
    "net/http"
    "log"
    "flag"
    "fmt"
)

var localPort *string = flag.String("p", "3000", "local port")
var logpath *string = flag.String("l", "path", "log file path")
var lines *string = flag.String("n", "100", "log file path")

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
  fmt.Printf("Listening: %v Serving: %v\n", *localPort, *logpath)

 http.HandleFunc("/", handler)
 http.HandleFunc("/tail", tailhandler)
 http.HandleFunc("/head", headhandler)
 http.HandleFunc("/grep", grephandler)
 err := http.ListenAndServe(":8000", nil)
        if err != nil {
                panic(err)
        }
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "log server")
}

func tailhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, tail())
}

func headhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, head())
}

func grephandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, grep("controller"))
}

func tail() string{
    app := "tail"
    arg0 := *logpath //*lines
    cmd := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return "error"
    }
return string(out)

}



func head() string{
    app := "head"
    arg0 := *logpath //*lines
    cmd := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return "error"
    }
return string(out)
}


func grep(search string) string{
    app := "grep"
    arg1 := *logpath //*lines
    arg0:= search
    cmd := exec.Command(app, arg0, arg1)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return "error"
    }
return string(out)	
}