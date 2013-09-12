

go run helloweb.go 

 wrk -r 10k http://127.0.0.1:4567/hi
Making 10000 requests to http://127.0.0.1:4567/hi
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   175.43us   93.84us 312.00us   71.43%
    Req/Sec    10.00k     0.00    10.00k   100.00%
  10000 requests in 461.55ms, 1.23MB read
Requests/sec:  21666.22
Transfer/sec:      2.67MB


go build

GOMAXPROCS=2 ./go-hello

wrk -r 10k http://127.0.0.1:4567/hi
Making 10000 requests to http://127.0.0.1:4567/hi
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   279.25us  114.20us 443.00us   75.00%
    Req/Sec    18.00k     0.00    18.00k   100.00%
  10000 requests in 266.72ms, 1.23MB read
Requests/sec:  37492.78
Transfer/sec:      4.61MB





