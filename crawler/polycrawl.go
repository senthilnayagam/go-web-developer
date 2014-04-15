/*

create fixed number of go routines in the begining, which will get from the queue to crawl

a script will process the new pages and add to queue




*/



package main

import (
	"fmt"
    "strconv"
	"sync"
    "time"
	
	"github.com/PuerkitoBio/goquery"
	"encoding/base64"
	"io/ioutil"
    "crypto/md5"
    "encoding/hex"
	"os"
	"net/url"
	"github.com/foize/go.fifo"
	"strings"

    "flag"
	"github.com/deckarep/golang-set"
	
	
)



func toAbsUrl(fullurl string, weburl string) string {
	baseurl,_ := url.Parse(fullurl) 
	relurl, err := url.Parse(weburl)
	if err != nil {
		return ""
	}
	absurl := baseurl.ResolveReference(relurl)
	return absurl.String()
}


func CrawlableUrl(testurl string)bool{
    u, err := url.Parse(testurl)
       if err != nil {
          // panic(err)
		  fmt.Println( err)
       }
	if u.Scheme == "http" {
		return true
	} else {
		return false
    }
	
	
}


func check(e error) {
    if e != nil {
      //  panic(e)
	  fmt.Println("file write error")
	  
    }
}


func FileExists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}


func write_file(f string, content string) {
	data := []byte(content)
	err := ioutil.WriteFile( f , data, 0644)
	check(err)
	
}


func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}




func GetPage(url string, mustcontain string,qChan chan string) {

	if CrawlableUrl(url) == false {
		fmt.Println(url + " not crawlable" )
		return
	}
	
	
	if strings.Contains(url,mustcontain)== false {
	fmt.Println(url + " does not contain " + mustcontain)
	return
    }

	md5 := GetMD5Hash(url)
	filename := "./output/" + md5 + ".json"	
	
	if FileExists(filename){
		fmt.Println("already crawled: " + filename)
		return
	}
	
//	visited_urls.Add(url)
	
	// create the file now
	write_file(filename,"") // so that possible deadlock is avoided
	

	doc, _ := goquery.NewDocument(url)
	html , _ :=  doc.Html()
	 str := base64.StdEncoding.EncodeToString([]byte(html))

	result := "{ 'urls':["
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url2, _ := s.Attr("href")


absurl := toAbsUrl(url,url2)
_ = absurl
		result= result + "'" + absurl + "',\n"  // url2
		global_url_queue.Add(absurl)
		//qChan <- absurl // queue to channel
		
	})
	result = result + "],\n 'url':'" +url +  "','base64html':'" + str + "'}"
	// fmt.Println(doc.Html())
	// fmt.Println(doc.Text())

	 write_file(filename,result)
	// fmt.Println(result)
	fmt.Println("crawled : " + url + " filename : " + filename )
	
	

}





// parallelizing via channels

func worker(linkChan chan string,qChan chan string, wg *sync.WaitGroup) {
   // Decreasing internal counter for wait-group as soon as goroutine finishes
   defer wg.Done()

   for url := range linkChan {
     // Analyze value and do the job here
	 time.Sleep(1 * time.Second)
	 GetPage(url ,*mustcontain,qChan)
	 		fmt.Printf("Done processing link #%s\n", url)
   }
}



func queuer(linkChan chan string, qChan chan string, wg *sync.WaitGroup) {
   // Decreasing internal counter for wait-group as soon as goroutine finishes
   defer wg.Done()
   /*
   for url := range qChan {
	   linkChan <- url
   }
*/
   
   for  {
           // get a new item from the things queue
           newurl := global_url_queue.Next();

           // check if there was an item
           if newurl == nil {
               fmt.Println("queue is empty")
   			// done.Done()
                return
           }
		
   		// fmt.Println("recursion starts: " + newurl)
		
   		s := fmt.Sprintf("%s", newurl)
   		fmt.Println("recursion starts: " + s)
       linkChan <- s
		
      }
      
   
}



var visited_urls = mapset.NewSet()
var global_url_queue = fifo.NewQueue()
var starturl *string = flag.String("u", "http://railsfactory.com/", "URL")
var mustcontain *string = flag.String("m", "railsfactory", "must contain string in URL")
var outputpath *string = flag.String("p", "output", "folder to save crawled files")

func main() {
    lCh := make(chan string) // initial crawl
	qCh := make(chan string) // queuing channel
    wg := new(sync.WaitGroup)

fmt.Println("start")

parallelCrawls := 10



    // Adding routines to workgroup and running then
    for i := 0; i < parallelCrawls ; i++ {
        wg.Add(1)
        go worker(lCh,qCh, wg)
		fmt.Println("routine :" + strconv.Itoa(i) )
    }
	
	lCh <- *starturl // crawl initialising by inserting in queue
	
	// a dedicated channel for queueing
	for i := 0; i < parallelCrawls ; i++ {
    wg.Add(1)
    go queuer(lCh,qCh, wg)
}
	
	
	
	
	
	/*
	// check if new url's to crawl
    for url := range qCh {
         lCh <- url
 	 		fmt.Printf("queing link: %s\n", url)
    }
	*/
	
	/*
	yourLinksSlice := make([]string, 5000)
		for i := 0; i < 5000; i++ {
			yourLinksSlice[i] = fmt.Sprintf("%d", i+1)
		}
		
    // Processing all links by spreading them to `free` goroutines
    for _, link := range yourLinksSlice {
        lCh <- link
    }
*/
    // Closing channel (waiting in goroutines won't continue any more)
    close(lCh)
	close(qCh)

    // Waiting for all goroutines to finish (otherwise they die as main routine dies)
    wg.Wait()
}




