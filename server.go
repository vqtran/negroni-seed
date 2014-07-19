package main

import (
  "github.com/codegangsta/negroni"
  "github.com/julienschmidt/httprouter"
  "github.com/vqtran/negroni-seed/app/views"
)

func main() {
  r := httprouter.New()
  r.GET("/", views.HomeHandler)
  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
