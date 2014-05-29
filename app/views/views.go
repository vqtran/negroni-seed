package views

import (
	"github.com/vqtran/amber"
	"html/template"
	"net/http"
	"log"
)

var templates map[string]*template.Template

func Load() {
	dopt, opt := amber.DefaultDirOptions, amber.DefaultOptions
	t, err := amber.CompileDir("app/views/templates", dopt, opt)
	if err != nil {
		log.Fatal(err)
	}
	templates = t
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string { "Name": "World" }
	err := templates["index"].Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
