package main

import (
    "fmt"
    "github.com/oschwald/geoip2-golang"
    "net"
)

func main() {
    db, err := geoip2.Open("GeoLite2-City.mmdb")
    if err != nil {
            panic(err)
    }
    // If you are using strings that may be invalid, check that ip is not nil
    ip := net.ParseIP("81.2.69.142")
    record, err := db.City(ip)
    if err != nil {
            panic(err)
    }
    fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
    fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
    fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
    fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
    fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
    fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)

    db.Close()
    // Output:
    // Portuguese (BR) city name: Londres
    // English subdivision name: England
    // Russian country name: Великобритания
    // ISO country code: GB
    // Time zone: Europe/London
    // Coordinates: 51.5142, -0.0931
}