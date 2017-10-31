package main
//all go files must be part of package main if you want to use their functionality without importing

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/gorilla/handlers"
	"os"
 	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"Starfleet/session"
	_"Starfleet/memory"
	"Starfleet/model"
	//"os/user"

	//"strconv"
)

var (
	tpl *template.Template
	db *gorm.DB
	err error
	globalSessions *session.Manager
)

type Person struct {
	Email string
	Password string
}


/*TODO:
1. user visits /student
2. app handles route, checks session, if valid, renders page for user id in session
2b. if not valid, redirects (302) to /login

*/



// Then, initialize the session manager
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()

	dbPassword := os.Getenv("PG_DATABASE_PW")
	db, err = gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password="+dbPassword)
	if err != nil {
		fmt.Println("Cannot connect to database...")
		fmt.Println("DB Error: ", err)
	}


	db.SingularTable(true)

	db.AutoMigrate(
		&model.MainUser{},
		&model.Student{},
		&model.PartTimeStudent{},
		&model.FullTimeStudent{},
		&model.Department{},
		&model.Faculty{},
		&model.Admin{},
	)
}




func main() {

	routes := mux.NewRouter()
	tpl = template.Must(template.ParseGlob("templates/*"))
	routes.PathPrefix("/style").Handler(http.StripPrefix("/style/",http.FileServer(http.Dir("style"))))
	routes.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	routes.HandleFunc("/",index)
	//routes.HandleFunc("/about/{number}", about)
	routes.HandleFunc("/login", loginPage).Methods("GET")
	routes.HandleFunc("/login", loginUser).Methods("POST")
	//routes.Handle("/user",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.Handle("/student",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.Handle("/admin",  checkSessionWrapper(displayAdmin)).Methods("GET")
	routes.Handle("/faculty",  checkSessionWrapper(displayFacultyt)).Methods("GET")
	//routes.HandleFunc("/unauthorized", unauthorized)

	routes.HandleFunc("/logout", logout)
	//routes.HandleFunc("/student", AuthHandler(displayUser))



	// USED FOR HEROKU
	//http.ListenAndServe(":" + os.Getenv("PORT"), routes)

	//USED FOR LOCAL, only use one
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout,routes))

	//defer db.Close(), want to keep db connection open
}


// routes for site

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index", nil)

}

func unauthorized(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index" , "You can not view this page")
}

func loginPage(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"login",nil)
}



func loginUser(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		//is a POST
		formEmail := r.FormValue("email")
		formPassword :=	r.FormValue("password")
		// Try to find user in DB
		user := model.MainUser{}
		db.Where(&model.MainUser{UserEmail: formEmail}).First(&user)



		if user.UserEmail != "" {
			dbPassword := user.UserPassword

			if user.CheckPasswordMatch(formPassword) {
				fmt.Println("User found in DB with email:", formEmail, " and password: ", dbPassword)
				sess.Set("username", r.Form["username"])
				sess.Set("UserID", user.UserID)
				//http.Redirect(w,r,"/user/" + strconv.Itoa(int(user.UserID)), http.StatusFound)
				//http.Redirect(w,r,"/user", http.StatusFound)
				checkUserType(user, w, r)
			} else {

				tpl.ExecuteTemplate(w,"login","Error, username or password does not match.")

			}


		} else {
			fmt.Println()
			tpl.ExecuteTemplate(w,"login","User not found")
		}

	}


}

func checkUserType(user model.MainUser, w http.ResponseWriter, r *http.Request){

	//cont := context.Get(r,"user")
	//TODO: 1. user visits /login (GET)
	//2. user submits login form (POST)
	//3. app validates request, logs in, if happy, redirects (302) to /student
	//4. user is redirected to /student, app gets the data it needs, passes to template and renders the student template

	switch user.UserType {

	case 1:
		fmt.Println("You're a student")
		fmt.Println("User data is", user.FirstName)
		http.Redirect(w,r,"/student", http.StatusFound)

		// The data is lost after redirect because it's a new request,
		// now I need to get the student data and render the template, which is a different request
		//since http is stateless, you lose the data structure after the first request.
	case 2:
		fmt.Println("Youre a faculty")
		http.Redirect(w,r,"/faculty", http.StatusFound)


	case 3:
		fmt.Println("Youre an admin")
		http.Redirect(w,r,"/admin", http.StatusFound)
		//tpl.ExecuteTemplate(w,"admin", "administrative user!")

	default:
		fmt.Println("Not sure your type")
		http.Redirect(w,r,"/", http.StatusFound)
		tpl.ExecuteTemplate(w,"index",nil)
		//return user,user.UserType
	}


}




func checkLoginStatus(w http.ResponseWriter, r *http.Request) (bool,model.MainUser){
	sess := globalSessions.SessionStart(w,r)
	sess_uid := sess.Get("UserID")
	u := model.MainUser{}
	if sess_uid == nil {
		//http.Redirect(w,r, "/", http.StatusForbidden)
		//tpl.ExecuteTemplate(w,"index", "You can't access this page")
		return false, u
	} else {
		uID := sess_uid
		db.First(&u, uID)
		fmt.Println("Logged in User, ", uID)
		//tpl.ExecuteTemplate(w, "user", nil)
		return true, u
	}
}
/*
In this snippet we're placing our handler logic in an anonymous function
 and closing-over the message variable to form a closure.
 We're then converting this closure to a handler by using the http.HandlerFunc adapter and returning it.
 */
func checkSessionWrapper(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware")
		isLogged, _ := checkLoginStatus(w, r)
		if isLogged { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/", http.StatusUnauthorized) // redirects route and gives unauthorized link
			tpl.ExecuteTemplate(w,"index", "You are unauthorized to acess that page") //this renders the index template right under it
		}

	})
}



func displayStudent(w http.ResponseWriter, r *http.Request){
	//TODO: Check only user can access correct roles


	_, user := checkLoginStatus(w, r)
	if user.UserType == 1 {
		tpl.ExecuteTemplate(w, "student", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func displayAdmin(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "admin", nil)
}


func displayFacultyt(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "faculty", nil)
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

//func AuthHandler(w http.ResponseWriter, r *http.Request) http.Handler {
//	http.Handler* h
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		loggedIn, _ := checkLoginUser(w,r)
//		if !loggedIn {
//			http.Redirect(w,r,"/login", 200)
//		} else {
//
//			h.ServeHTTP(w, r)
//		}
//	})
//}

func logout(w http.ResponseWriter, r *http.Request){
	sess := globalSessions.SessionStart(w, r)
	//sid := sess.SessionID()
	sess.Delete("UserID")
	sess.Delete("username")
	http.Redirect(w,r,"/login", http.StatusSeeOther)
}