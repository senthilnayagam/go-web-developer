package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
)

var geodb *geoip2.Reader

func main() {
	var err error
	geodb, err = geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}
	defer geodb.Close()

	/*

	   // If you are using strings that may be invalid, check that ip is not nil
	   ip := net.ParseIP("81.2.69.142")
	   record, err := geodb.City(ip)
	   if err != nil {
	           panic(err)
	   }
	   fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	   fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	   fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	   fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	   fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	   fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	*/

	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931

	mux := pat.New()
	mux.Get("/ip/:ip", http.HandlerFunc(ip2location))

	http.Handle("/", mux)

	log.Println("Listening port:3000 ")
	http.ListenAndServe(":3000", nil)

}

func ip2location(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	//  tablename := params.Get(":name")
	//  fmt.Printf(name)

	rawip := params.Get(":ip")
	ip := net.ParseIP(rawip)
	_ = ip
	record, err := geodb.City(ip)
	if err != nil {
		// panic(err)
		log.Println(err)
	}
	fmt.Println(record) //debbugging

	iso := record.Country.IsoCode
	timezone := record.Location.TimeZone

	w.Write([]byte("{[\"ISO country code\":\"" + iso + "\",\"ip\":" + rawip + "\",\"timezone\":" + timezone + "]}")) // hand coded json
}
