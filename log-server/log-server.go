package main

import (
    "os"
    "os/exec"
    "os/signal"
    "syscall"
   // "net/http"
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
fmt.Printf("Tail: \n")
tail()
fmt.Printf("Head: \n") 
head()
search := "controller"
fmt.Printf("grep: %s\n",search) 
grep(search)
}

func tail(){
    app := "tail"
    arg0 := *logpath //*lines
    cmd := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }
print(string(out))

}



func head(){
    app := "head"
    arg0 := *logpath //*lines
    cmd := exec.Command(app, arg0)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }
print(string(out))	
}


func grep(search string){
    app := "grep"
    arg1 := *logpath //*lines
    arg0:= search
    cmd := exec.Command(app, arg0, arg1)
    out, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }
print(string(out))	
}