package main
//all go files must be part of package main if you want to use their functionality without importing

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	//"strconv"
	"fmt"
	"github.com/gorilla/handlers"
	"os"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	tpl *template.Template
)

type Person struct {
	Email string
	Password string
}

func main() {

	db, err := gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password=gd62885")
	if err != nil {
		fmt.Println("Cannot connect to database...")
		fmt.Println("DB Error: ", err)
	}

	type User struct {
		UserID uint `gnorm:"primary_key"`
		UserEmail string `gnorm:"type:varchar(20);unique"`
		UserPassword string `gnorm:"type:varchar(300)"`
		FirstName string `gnorm:"type:varchar(50)"`
		LastName string `gnorm:"type:varchar(50)"`
		UserType int
	}

	type Student struct {
		StudentID uint `gnorm:"primary_key"`
		User  User `gorm:"ForeignKey:UserRefer"`
		UserRefer uint
	}
	
	db.AutoMigrate(&User{}, &Student{})

	routes := mux.NewRouter()
	tpl = template.Must(template.ParseGlob("templates/*"))
	routes.PathPrefix("/style").Handler(http.StripPrefix("/style/",http.FileServer(http.Dir("style"))))
	routes.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	routes.HandleFunc("/",index)
	//routes.HandleFunc("/about/{number}", about)
	routes.HandleFunc("/login", loginPage)
	routes.HandleFunc("/login/{num}", loginUser)


	// USED FOR HEROKU
	//http.ListenAndServe(":" + os.Getenv("PORT"), routes)

	//USED FOR LOCAL, only use one
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout,routes))

	defer db.Close()
}


func index(w http.ResponseWriter, r *http.Request){

	tpl.ExecuteTemplate(w, "index", nil)

}

func loginPage(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"login",nil)
}

func loginUser(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	fmt.Println(vars)

	userEmail := r.FormValue("email")
	userPassword :=	r.FormValue("password")

	p := Person{userEmail,userPassword}
	fmt.Println("Email: ", userEmail)
	fmt.Println("Password: ", userPassword)
	tpl.ExecuteTemplate(w,"login",p)
}


/*
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

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.public", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
*/