package main

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/hoisie/mustache"
  "net/http"
  "fmt"
)

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string { "Name": "World" }
	html := mustache.RenderFileInLayout("app/views/index.html",
													"app/views/base.html",
													data)
	fmt.Fprintf(w, html)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
