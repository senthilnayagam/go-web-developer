

simple api for geoip



with single IP address load test it does about 11k requests/second, but with appropriate caching should be possible to take it 20k requests/second


installation

clone the repo, change to geoip-api-server directory

download the go library

go get github.com/oschwald/geoip2-golang

go build geoip-api-server.go 


download the Maximind database, extract it, change the path in code(os simply use the current directory)


run the app

./geoip-api-server
2014/05/08 16:18:22 Listening port:3000






curl http://127.0.0.1:3000/ip/202.54.1.5
{["ISO country code":"IN","ip":202.54.1.5","timezone":Asia/Kolkata]}

curl http://127.0.0.1:3000/ip/81.2.69.142
{["ISO country code":"GB","ip":81.2.69.142","timezone":Europe/London]}


load test

wrk http://127.0.0.1:3000/ip/81.2.69.142
Running 10s test @ http://127.0.0.1:3000/ip/81.2.69.142
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.88ms  473.14us   2.43ms   61.64%
    Req/Sec     5.93k   379.81     6.89k    66.63%
  112711 requests in 10.00s, 20.10MB read
Requests/sec:  11271.02
Transfer/sec:      2.01MB




todo

object should be converted into valid json and returned
additional fields need to be shared

