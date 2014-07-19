package views

import (
   "github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
)

var r *render.Render = render.New(render.Options{
	Directory: "app/views/templates",
	Layout: "layouts/layout",
	Extensions: []string{".html"},
	IsDevelopment: true,
})

func HomeHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := map[string]string { "Name": "World" }
	r.HTML(w, http.StatusOK, "index", data)
}


