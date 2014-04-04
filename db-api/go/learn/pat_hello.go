package main

import (
  "github.com/bmizerany/pat"
  "log"
  "net/http"
)

func main() {
  mux := pat.New()
  mux.Get("/user/:name/profile", http.HandlerFunc(profile))

  http.Handle("/", mux)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  name := params.Get(":name")
  w.Write([]byte("Hello " + name))
}