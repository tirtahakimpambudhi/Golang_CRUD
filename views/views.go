package views

import (
	"html/template"
	"net/http"
	"sync"
	"strings"
)

var (
	views     *template.Template
	viewsLayouts     *template.Template
	viewsOnce sync.Once
)
type HTML struct {
  Header string
  Main string
  Footer string
}


func Load() {
	views = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func LoadWithLayouts ( t HTML ) {
  viewsLayouts = template.Must(template.ParseFiles(
    t.Main,
    t.Header,
    t.Footer,
    ))
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	viewsOnce.Do(Load)

	err := views.ExecuteTemplate(w, tmpl+".gohtml", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func RenderWithLayouts(w http.ResponseWriter, t HTML, data interface{}) {
	LoadWithLayouts(t)
	pathMain := strings.Split(t.Main, "/")
	tmpl := pathMain[len(pathMain)-1]
	

	err := viewsLayouts.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
