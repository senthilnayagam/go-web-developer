package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"encoding/base64"
	"io/ioutil"
    "crypto/md5"
    "encoding/hex"
	"os"
	"net/url"
	"github.com/foize/go.fifo"
	"strings"
//	"sync"
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




func GetPage(url string, mustcontain string) {
	// defer done.Done()
	if CrawlableUrl(url) == false {
		fmt.Println(url + " not crawlable" )
		// done.Done()
		return
	}
	
	
	if strings.Contains(url,mustcontain)== false {
	fmt.Println(url + " does not contain " + mustcontain)
	// done.Done()
	return
}

	md5 := GetMD5Hash(url)
	filename := "./output/" + md5 + ".json"	
	
	if FileExists(filename){
		fmt.Println("already crawled: " + filename)
	// 	done.Done()
		return
	}
	
	visited_urls.Add(url)
	
	// create the file now
	write_file(filename,"")
	
	
	

	
	doc, _ := goquery.NewDocument(url)
	html , _ :=  doc.Html()
	 str := base64.StdEncoding.EncodeToString([]byte(html))
	// fmt.Println(str)
	
	//defer doc.Body.Close()
	result := "{ 'urls':["
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url2, _ := s.Attr("href")


absurl := toAbsUrl(url,url2)
_ = absurl
		result= result + "'" + absurl + "',\n"  // url2
		global_url_queue.Add(absurl)
		
	})
	result = result + "],\n 'url':'" +url +  "','base64html':'" + str + "'}"
	// fmt.Println(doc.Html())
	// fmt.Println(doc.Text())

	 write_file(filename,result)
	// fmt.Println(result)
	fmt.Println("crawled : " + url + " filename : " + filename )
	
	

}

// var done sync.WaitGroup
var visited_urls = mapset.NewSet()
var global_url_queue = fifo.NewQueue()

var starturl *string = flag.String("u", "http://railsfactory.com/", "URL")
var mustcontain *string = flag.String("m", "railsfactory", "must contain string in URL")
var outputpath *string = flag.String("p", "output", "folder to save crawled files")

func main() {

	  flag.Parse()
  fmt.Printf("url: %v must contain: %v and save output to: %v", *starturl, *mustcontain,*outputpath)
	GetPage(*starturl,*mustcontain)



for  {
        // get a new item from the things queue
        newurl := global_url_queue.Next();

        // check if there was an item
        if newurl == nil {
            fmt.Println("queue is empty")
			// done.Done()
             return
			os.Exit(0)
        }
		
		// fmt.Println("recursion starts: " + newurl)
		
		s := fmt.Sprintf("%s", newurl)
		fmt.Println("recursion starts: " + s)
		// done.Add(1)
		// go GetPage(s,mustcontain)
		GetPage(s,*mustcontain)
		
   }


	
	// done.Wait()
	// os.Exit(0)
}