log-server

log server will show the n lines of the log file specified at port p , head for top of the log, tail for the bottom of the log, and grep for search result

uses exec to run tail, head or grep command from the operating system, should work on mac, and linux type operating systems




go run log-server.go -l log/production.log 


http://localhost:8000/tail
http://localhost:8000/head

todo
http://localhost:8000/grep/keyword

planned
log-server -p 2000 -f ~/app/log/development.log -n 100

load from a config file/default config file

authentication?

