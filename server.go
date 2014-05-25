package main

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
