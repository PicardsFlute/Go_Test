package main
//all go files must be part of package main if you want to use their functionality without importing

import (
	"net/http"
	"html/template"
	"os"
	"github.com/gorilla/mux"
	"strconv"
)

var (
	tpl *template.Template
)

type Person struct {
	Name string
	Age int
}

func main() {
	routes := mux.NewRouter()
	tpl = template.Must(template.ParseGlob("templates/*"))
	routes.PathPrefix("/style").Handler(http.StripPrefix("/style/",http.FileServer(http.Dir("style"))))
	routes.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	routes.HandleFunc("/",index)
	routes.HandleFunc("/about/{number}", about)
	/*
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	server.ListenAndServe()
	*/
	//http.ListenAndServe(os.Getenv("PORT"), nil)
	http.ListenAndServe(":" + os.Getenv("PORT"), routes)
}


func index(w http.ResponseWriter, r *http.Request){

	tpl.ExecuteTemplate(w, "index.html", nil)

}

func about(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) //returns a mapping responses
	personId := vars["number"] //get map with key id number

	if num, _ := strconv.Atoi(personId); num > 3 {
		p := Person{"Bob", 4}
		tpl.ExecuteTemplate(w, "index.html", p)

	}else {
		p := Person{"Steve", 2}
		tpl.ExecuteTemplate(w, "index.html", p)
	}
}
