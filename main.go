package main
//all go files must be part of package main if you want to use their functionality without importing

import (
	"net/http"
	"html/template"
)

var (
	tpl *template.Template
)

type Person struct {
	Name string
	Age int
}

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/",index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request){

	tpl.ExecuteTemplate(w, "index.html", nil)

}

func about(w http.ResponseWriter, r *http.Request){

	p := Person{"Bob",4}
	tpl.ExecuteTemplate(w, "index.html", p)

}
