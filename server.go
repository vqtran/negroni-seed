package main

import (
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "github.com/vqtran/negroni-seed/app/views"
)

func main() {
  views.Load()
  r := mux.NewRouter()
  r.HandleFunc("/", views.HomeHandler)
  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
