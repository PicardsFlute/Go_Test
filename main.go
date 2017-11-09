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
	"Starfleet/global"

	"io/ioutil"
)

var (

	db  *gorm.DB
	err error
	globalSessions *session.Manager
)


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
		&model.PartTimeFaculty{},
		&model.FullTimeFaculty{},
		&model.Admin{},
		&model.Researcher{},
		&model.Major{},
		&model.Minor{},
		&model.StudentMajor{},
		&model.StudentMinor{},

		&model.Hold{},
		&model.StudentHolds{},

		&model.Course{},
		&model.Prerequisite{},

		&model.Day{},
		&model.Semester{},
		&model.Period{},
		&model.TimeSlot{},
		&model.Section{},
		&model.Enrollment{},
		&model.StudentHistory{},

	)
}




func main() {

	routes := mux.NewRouter()
	global.Tpl = template.Must(template.ParseGlob("templates/*"))
	routes.PathPrefix("/style").Handler(http.StripPrefix("/style/",http.FileServer(http.Dir("style"))))
	routes.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	routes.HandleFunc("/",index)
	//routes.HandleFunc("/about/{number}", about)
	routes.HandleFunc("/login", loginPage).Methods("GET")
	routes.HandleFunc("/login", loginUser).Methods("POST")
	//routes.Handle("/user",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.Handle("/student",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.Handle("/admin",  checkSessionWrapper(displayAdmin)).Methods("GET")
	routes.Handle("/faculty",  checkSessionWrapper(displayFaculty)).Methods("GET")
	routes.Handle("/researcher", checkSessionWrapper(displayResearcher)).Methods("GET")


	routes.Handle("/admin/student" , checkSessionWrapper(ViewStudentSchedulePage)).Methods("GET")
	routes.HandleFunc("/admin/student/{student}", ViewStudentSchedule).Methods("GET")

	routes.HandleFunc("/admin/holds", ViewStudentHoldsPage)
	routes.HandleFunc("/admin/holds/{user}", ViewStudentHolds).Methods("GET")
	routes.HandleFunc("/admin/holds/{user}/{id}", AdminDeleteHold).Methods("POST")

	//routes.HandleFunc("/admin/student/holds/{student}", ViewStudentHolds)
	routes.Handle("/admin/course",checkSessionWrapper(AdminAddCoursePage))
	routes.HandleFunc("/admin/course/{course}",AdminAddCourse).Methods("POST")
	routes.HandleFunc("/admin/course/search", AdminSearchCoursePage).Methods("GET")
	//routes.HandleFunc("/admin/course/", AdminDeleteCourse)
	routes.HandleFunc("/admin/section", AdminAddSectionPage)
	//routes.Handle("/admin/course/{course}",checkSessionWrapper(AdminAddCoursePage))

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
	global.Tpl.ExecuteTemplate(w, "index", nil)

}

func unauthorized(w http.ResponseWriter, r *http.Request){
	global.Tpl.ExecuteTemplate(w, "index" , "You can not view this page")
}

func loginPage(w http.ResponseWriter, r *http.Request){
	global.Tpl.ExecuteTemplate(w,"login",nil)
}

func redirectPost(w http.ResponseWriter, r *http.Request){
	req, err := http.NewRequest("DELETE", "/admin/holds/{id}", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	respBody , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP RESPONSE FROM DELETE IS", string(respBody))

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

				//Tpl.ExecuteTemplate(w,"user",user)
				checkUserType(user, w, r)
			} else {

				global.Tpl.ExecuteTemplate(w,"login","Error, username or password does not match.")

			}


		} else {
			fmt.Println()
			global.Tpl.ExecuteTemplate(w,"login","User not found")
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
		//Tpl.ExecuteTemplate(w,"admin", "administrative user!")

	case 4:
		fmt.Println("Youre a researcher")
		http.Redirect(w,r,"/researcher", http.StatusFound)
		//Tpl.ExecuteTemplate(w,"admin", "administrative user!")

	default:
		fmt.Println("Not sure your type")
		http.Redirect(w,r,"/", http.StatusFound)
		global.Tpl.ExecuteTemplate(w,"index",nil)
		//return user,user.UserType
	}


}




func CheckLoginStatus(w http.ResponseWriter, r *http.Request) (bool,model.MainUser){
	sess := globalSessions.SessionStart(w,r)
	sess_uid := sess.Get("UserID")
	u := model.MainUser{}
	if sess_uid == nil {
		//http.Redirect(w,r, "/", http.StatusForbidden)
		//Tpl.ExecuteTemplate(w,"index", "You can't access this page")
		return false, u
	} else {
		uID := sess_uid
		db.First(&u, uID)
		fmt.Println("Logged in User, ", uID)
		//Tpl.ExecuteTemplate(w, "user", nil)
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
		isLogged, _ := CheckLoginStatus(w, r)
		if isLogged { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}



func displayStudent(w http.ResponseWriter, r *http.Request){

	_, user := CheckLoginStatus(w, r)
	if user.UserType == 1 {
		global.Tpl.ExecuteTemplate(w, "student", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func displayFaculty(w http.ResponseWriter, r *http.Request){
	_, user := CheckLoginStatus(w,r)

	if user.UserType == 2 {
		global.Tpl.ExecuteTemplate(w, "faculty", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func displayAdmin(w http.ResponseWriter, r *http.Request){
	_, user := CheckLoginStatus(w,r)

	if user.UserType == 3 {
		m := map[string]interface{}{
			"User":user,
		}
		global.Tpl.ExecuteTemplate(w, "admin", m)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}



func displayResearcher(w http.ResponseWriter, r *http.Request){

	_, user := CheckLoginStatus(w,r)

	if user.UserType == 4 {
		global.Tpl.ExecuteTemplate(w, "researcher", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}



/*
func about(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) //returns a mapping responses
	personId := vars["number"] //get map with key id number
	if num, _ := strconv.Atoi(personId); num > 3 {
		p := Person{"Bob", 4}
		Tpl.ExecuteTemplate(w, "index.html", p)
	}else {
		p := Person{"Steve", 2}
		Tpl.ExecuteTemplate(w, "index.html", p)
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