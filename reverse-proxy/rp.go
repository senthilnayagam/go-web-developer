package main
 
import(
        "log"
        "net/url"
        "net/http"
        "net/http/httputil"
)
 
func main() {
        remote, err := url.Parse("http://127.0.0.1:4567/")
        if err != nil {
                panic(err)
        }
 
        proxy := httputil.NewSingleHostReverseProxy(remote)
        http.HandleFunc("/", handler(proxy))
        err = http.ListenAndServe(":8000", nil)
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