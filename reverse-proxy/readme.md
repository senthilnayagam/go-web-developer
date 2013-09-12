reverse proxy

seamless seems like a good proxy with round robin and zero downtime switching
https://bitbucket.org/tebeka/seamless/src

but for smaller tasks a simple reverse proxy would do

initial reverse proxy code taken from https://gist.github.com/JalfResi/6287706




start the sinatra hello world app

ruby sinatra/hello.rb 

== Sinatra/1.4.3 has taken the stage on 4567 for development with backup from Thin
>> Thin web server (v1.5.1 codename Straight Razor)
>> Maximum connections set to 1024
>> Listening on localhost:4567, CTRL+C to stop


start the reverse proxy


run the load test

go run rp.go 


wrk -c 10 -r 10k http://127.0.0.1:8000/hi
Making 10000 requests to http://127.0.0.1:8000/hi
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.80ms    5.02ms  38.88ms   97.39%
    Req/Sec     0.00      0.00     0.00    100.00%
  10000 requests in 8.66s, 2.77MB read
Requests/sec:   1154.64
Transfer/sec:    327.00KB




todo

configure from console




GOMAXPROCS=2 ./rp -u http://127.0.0.1:4567/












faced some issues 

on reverse proxy script console was getting these errors randomly, possibly port allocation issue

2013/09/12 13:06:46 http: proxy error: dial tcp 127.0.0.1:4567: can't assign requested address
2013/09/12 13:06:46 http: proxy error: dial tcp 127.0.0.1:4567: can't assign requested address


wrk -r 100k http://127.0.0.1:3000/hi
Making 100000 requests to http://127.0.0.1:3000/hi
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    81.25ms   53.32ms 685.81ms   63.12%
    Req/Sec   189.77    524.67     2.00k    87.09%
  100000 requests in 2.11m, 12.34MB read
  Non-2xx or 3xx responses: 6157
Requests/sec:    788.56
Transfer/sec:     99.62KB



giving output to console slowed it down to a extent, should disable it

