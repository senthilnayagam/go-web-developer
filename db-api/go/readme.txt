


mysql connection pooling and exposing tables as simple API


install

go get github.com/go-sql-driver/mysql







wrk --timeout 2s --latency -d 2m -c 100 -t 10 http://127.0.0.1:3000/mysql/car/4





todo
to match features with ruby and nodejs version
json
generic json conversion
generic table version
logging



12k(pat) regular

17k (http router)

21k requests/minute ( httprouter + maxprocs)