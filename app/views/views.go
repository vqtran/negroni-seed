package views

import (
	"github.com/vqtran/tea"
	"net/http"
)

func Load() {
	tea.SetEngine("amber")
	tea.MustCompile("app/views/templates/", tea.Options{".amber", true})
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string { "Name": "World" }
	tea.Render(w, "index", data)
}


