unicrawl is a simple single threaded crawler which saves the output in json file, html is base64 encoded

webcrawl is a simple http server which can crawl a url, later want to implement unicrawl into it


go build unicrawl.go

./unicrawl -u http://www.rediff.com -m rediff -p output



go build webcrawl
./webcrawl

open http://127.0.0.1:4040/Get/www.railsfactory.com

done
a) fetch page
b) extract links
c) save as json file
d) relative url to absolute url
e) limit crawling to pattern
f) not try crawling non http link schemes(eg mailto)
g) do not crawl already crawled link
h) recursion
i) get url and pattern and output folder from input(unicrawl)



todo


a) make it concurrent, run each request on a go routine

b) Use a proper HTTP client
* http client config, timeout, retry, user agent,cookies, header etc,
* speed throttling
* delay between requests
* limit concurrency to a fixed number
* use proxy if needed
* handle get, post
* follow redirects ?
* handle url with params
* timeout
* retry

c) load config from config file
d) log in file
e) depth first search / breadth first search
f) save in  database(optional)
g) robots.txt


