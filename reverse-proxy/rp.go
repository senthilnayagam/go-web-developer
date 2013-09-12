package main
 
import(
        "log"
        "net/url"
        "net/http"
        "net/http/httputil"
        "os"
        "os/signal"
        "syscall"
        "flag"
        "fmt"
)


func signalCatcher() {
        ch := make(chan os.Signal)
        signal.Notify(ch, syscall.SIGINT)
        <-ch
        log.Println("CTRL-C; exiting")
        os.Exit(0)
}

var localPort *string = flag.String("p", "3000", "local port")
var remoteurl *string = flag.String("u", "public", "http://url_with_port")
 
func main() {
        go signalCatcher()
        flag.Parse()
        remote, err := url.Parse(*remoteurl) //
        if err != nil {
                panic(err)
        }
 
        proxy := httputil.NewSingleHostReverseProxy(remote)
        http.HandleFunc("/", handler(proxy))
        err = http.ListenAndServe(fmt.Sprint(":",*localPort), nil)
        if err != nil {
                panic(err)
        }
}
 
func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
        return func(w http.ResponseWriter, r *http.Request) {
                log.Println(r.URL)
                p.ServeHTTP(w, r)
        }
}