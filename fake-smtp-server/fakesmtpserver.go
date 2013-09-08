package main


import (
    "os"
    "net"
    "os/signal"
    "bufio"
    "syscall"
    "log"
    "flag"
    "fmt"
    "strings"	
	)
	
func signalCatcher() {
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT)
        <-ch
        log.Println("CTRL-C; exiting")
        os.Exit(0)
}

var localPort *string = flag.String("p", "2525", "local port")


func main() {
  go signalCatcher()
  flag.Parse()
  fmt.Printf("fake smtp server \n Listening: %v \n", *localPort)	

listener, err := net.Listen("tcp", fmt.Sprint(":",*localPort))
    if err != nil {
		fmt.Fprintf(os.Stderr, "Error.. %s", err.Error())
	}
	
for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// run as goroutine
		go FakeSmtp(conn)
	}


	
}


func FakeSmtp(conn net.Conn) {
defer conn.Close()
var eol = "\r\n"

reader := bufio.NewReader(conn)

fmt.Fprintf(conn, "220 senthil.com Senthil SMTP server 0.1 " + eol) // , host

    for {
		// Reads a line from the client
		raw_line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error!!" + err.Error())
			return
		}

        // Parses the command
		cmd, _ := getCommand(raw_line)
		
switch cmd {
                case "HELO", "EHLO": //
                         fmt.Fprintf(conn, "250 localhost Hello" + eol)
                         println("command HELO/EHLO")

                case "QUIT":
                        fmt.Fprintf(conn, "221 2.0.0 Bye" + eol)
                        println("command QUIT")
                        return
                        
                case "RSET":
                       fmt.Fprintf(conn, "250 OK" + eol)
                      // println("command RSET")
                case "NOOP":
                       //  println("command NOOP")
                case "MAIL":
                      fmt.Fprintf(conn, "250 OK" + eol)
                        //   println("command MAIL")
/*
                        arg := line.Arg() // "From:<foo@bar.com>"
                        m := mailFromRE.FindStringSubmatch(arg)
                        if m == nil {
                                log.Printf("invalid MAIL arg: %q", arg)
                                s.sendlinef("501 5.1.7 Bad sender address syntax")
                                continue
                        }
                       s.handleMailFrom(m[1])
*/
                case "RCPT":
                            fmt.Fprintf(conn, "250 Accepted" + eol)
                         //   println("command RCPT")
                case "DATA":
                          fmt.Fprintf(conn, "354 Enter message, ending with \".\" on a line by itself" + eol)
                        //  println("command DATA")
                default:
fmt.Fprintf(conn, "500 unrecognized command" + eol)
//println("default command")
 println(cmd)

                }
println(raw_line)


   } 

}	



func getCommand(line string) (string, []string) {
	line = strings.Trim(line, "\r\n")
	cmd := strings.Split(line, " ")
	return cmd[0], cmd[1:]
}
