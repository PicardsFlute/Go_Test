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
	"strconv"
	"io/ioutil"
	"strings"
	//"time"
	//"log"
	//"encoding/csv"

	"time"
	"log"
	//"encoding/csv"
	"github.com/mohae/struct2csv"

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

	//for heroku
	//db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))


	//dbPassword := os.Getenv("PG_DATABASE_PW")
	//dbConnectString := os.Getenv("DATABASE_URL")
	//db, err = gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password="+dbPassword)



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
		&model.Advises{},

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
	routes.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	routes.HandleFunc("/",index)
	//routes.HandleFunc("/about/{number}", about)
	routes.HandleFunc("/login", loginPage).Methods("GET")
	routes.HandleFunc("/login", loginUser).Methods("POST")
	//routes.Handle("/user",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.Handle("/faculty",  checkSessionWrapper(displayFaculty)).Methods("GET")
	routes.Handle("/researcher", checkSessionWrapper(displayResearcher)).Methods("GET")

	/* Master schedule searching*/
	routes.HandleFunc("/course", searchMasterScheduleForm).Methods("GET")
	routes.HandleFunc("/course/", searchMasterSchedule).Methods("GET")


	/*Admin routes */
	
	routes.Handle("/admin",  checkSessionWrapper(displayAdmin)).Methods("GET")
	routes.HandleFunc("/admin/student" , ViewStudentSchedulePage).Methods("GET")
	routes.HandleFunc("/admin/student/{student}", ViewStudentSchedule).Methods("GET")

	routes.HandleFunc("/admin/transcript", viewStudentTranscriptPage).Methods("GET")
	routes.HandleFunc("/admin/transcript/{student}", viewStudentTranscript).Methods("GET")

	routes.HandleFunc("/admin/holds", ViewStudentHoldsPage)
	routes.HandleFunc("/admin/holds/{user}", ViewStudentHolds).Methods("GET")
	routes.HandleFunc("/admin/holds/{user}/{id}", AdminDeleteHold).Methods("POST")

	//routes.HandleFunc("/admin/student/holds/{student}", ViewStudentHolds)
	routes.Handle("/admin/course",checkSessionWrapper(AdminAddCoursePage))
	routes.HandleFunc("/admin/course/{course}",AdminAddCourse).Methods("POST")
	//routes.HandleFunc("/admin/course/{course}/{pre-req}/",AddCoursePreRequisit).Methods("POST")
	routes.HandleFunc("/admin/courses/prereq",AddCoursePreRequisit).Methods("POST")

	routes.HandleFunc("/admin/course/search", AdminSearchCoursePage).Methods("GET")
	routes.HandleFunc("/admin/course/search/{course}", AdminSearchCourse).Methods("GET")
	routes.HandleFunc("/course/update", UpdateCourse).Methods("POST")
	//routes.HandleFunc("/admin/course/", AdminDeleteCourse)
	routes.HandleFunc("/admin/section", AdminAddSectionPage)
	routes.HandleFunc("/section/update", AdminUpdateSectionForm).Methods("GET")
	routes.HandleFunc("/section/update", AdminUpdateSection).Methods("POST")

	routes.HandleFunc("/admin/section/{section}", AdminAddSection)
	routes.HandleFunc("/admin/section/room/{id}", GetRoomsForBuilding)
	routes.HandleFunc("/admin/section/department/{id}", GetDepartmentsForSections).Methods("GET")
	routes.Handle("/admin/user", checkSessionWrapper(newUserForm)).Methods("GET")
	routes.Handle("/admin/user", checkSessionWrapper(createUser)).Methods("POST")
	routes.Handle("/admin/user/student", checkSessionWrapper(createStudent)).Methods("POST")
	routes.Handle("/admin/user/faculty", checkSessionWrapper(createFaculty)).Methods("POST")
	routes.Handle("/admin/user/search" , checkSessionWrapper(searchUser)).Methods("GET")
	routes.Handle("/admin/user/{userID}/delete", checkSessionWrapper(deleteUser)).Methods("POST")
	routes.HandleFunc("/admin/semester" , changeSemesterStatusForm).Methods("GET")
	routes.HandleFunc("/admin/semester" , changeSemesterStatus).Methods("POST")
	routes.HandleFunc("/admin/time", addSectionTime).Methods("POST")


	/* Student Routes*/
	routes.Handle("/student",  checkSessionWrapper(displayStudent)).Methods("GET")
	routes.HandleFunc("/student/schedule", ViewSchedule).Methods("GET")
	routes.HandleFunc("/student/registered", ViewRegisteredCourses).Methods("GET")
	routes.HandleFunc("/student/registered", DropRegisteredCourse).Methods("POST")
	routes.HandleFunc("/student/holds", ViewHolds).Methods("GET")
	routes.HandleFunc("/student/advisor", ViewAdvisor).Methods("GET")
	routes.HandleFunc("/student/transcript", ViewTranscript).Methods("GET")
	routes.HandleFunc("/student/search", AddCoursePage).Methods("GET")
	//routes.HandleFunc("/student/register", StudentSearchCourseResults).Methods("GET")
	routes.HandleFunc("/student/register", RegisterForSection).Methods("POST")




	//TODO: Custom auth middlewear for each user type

	//routes.HandleFunc("/unauthorized", unauthorized)

	/*Faculty Routes */
	routes.HandleFunc("/faculty/schedule", facultyViewSchedule).Methods("GET")
	routes.HandleFunc("/faculty/grades", giveStudentGradesPage).Methods("GET")
	routes.HandleFunc("/faculty/grades/{sectionID}", giveStudentGradesForm).Methods("GET")
	routes.HandleFunc("/faculty/grades/{sectionID}", submitGrades).Methods("POST")


	//routes.Handle("/course/search", checkSessionWrapper(searchMasterScheduleForm)).Methods("GET")

	routes.HandleFunc("/research/students/grades", getStudentsReportByGrade).Methods("GET")
	routes.HandleFunc("/research/students/grades", genReportStudentsByGrade).Methods("POST")

	routes.HandleFunc("/logout", logout)
	//routes.HandleFunc("/student", AuthHandler(displayUser))



	// USED FOR HEROKU
	//http.ListenAndServe(":" + os.Getenv("PORT"),handlers.LoggingHandler(os.Stdout,routes))
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout,routes))

	//defer db.Close(), want to keep db connectioT"), routes)

	//USED FOR LOCAL, only use onen open


}


// routes for site

func index(w http.ResponseWriter, r *http.Request){
	global.Tpl.ExecuteTemplate(w, "index", nil)

}

func loginPage(w http.ResponseWriter, r *http.Request){
	// use session information to determine the user, and if they are already logged in
	logged, u := CheckLoginStatus(w, r)
	// a user already logged in will be sent to the page of their respective role
	if logged{
		checkUserType(u,w,r)
	}else {
		global.Tpl.ExecuteTemplate(w, "login", nil)
	}
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

	switch user.UserType {

	case 1:
		fmt.Println("You're a student")
		fmt.Println("User data is", user.FirstName)
		http.Redirect(w,r,"/student", http.StatusFound)

		// The data is lost after redirect because it's a new request,
		// now I need to get the student data and render the template, which is a different request
		//since http is stateless, you l;ose the data structure after the first request.
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

func checkStudent(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for student")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType ==1 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkFaculty(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for faculty")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 2 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkAdmin(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for admin")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 3 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkResearcher(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for researcher")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 5 { //if check user is true, execute the handle that's inside
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


func logout(w http.ResponseWriter, r *http.Request){
	sess := globalSessions.SessionStart(w, r)
	//sid := sess.SessionID()
	sess.Delete("UserID")
	sess.Delete("username")
	http.Redirect(w,r,"/login", http.StatusSeeOther)
}


/* CRUD for users */

func newUserForm(w http.ResponseWriter, r *http.Request) {
	res := global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", nil)
	if res != nil{
		println("newUserForm: ", res.Error())
	}
}

func createUser (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formEmail := r.FormValue("email")

	userDB := model.MainUser{}
	userDB.UserEmail = formEmail

	count := 1
	db.Model(&model.MainUser{}).Where("user_email = ?", formEmail).Count(&count)
	if count == 0 {
		userDB.FirstName = r.FormValue("first-name")
		userDB.LastName = r.FormValue("last-name")
		userDB.UserPassword = r.FormValue("password")
		userType, _ := strconv.Atoi(r.FormValue("user-type"))
		userDB.UserType = userType
		valid, err := userDB.ValidateData()
		if valid {
			db.Create(&userDB)

			switch userDB.UserType {

			case 1:
				fmt.Println("You're a student")
				student := model.Student{}
				m := map[string]interface{}{
					"User":    userDB,
					"student": student,
				}
				res := global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", m)
				if res != nil{
					println("newUserForm: ", res.Error())
				}
				//return


			case 2:
				fmt.Println("Youre a faculty")
				faculty := model.Faculty{}
				m := map[string]interface{}{
					"User":    userDB,
					"faculty": faculty,
				}
				global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", m)

			case 3:
				fmt.Println("Youre an admin")
				admin := model.Admin{AdminID: userDB.UserID}
				db.Create(&admin)
				http.Redirect(w, r, "/admin", http.StatusFound)
				displayAdmin(w, r)

			case 4:
				fmt.Println("Youre a researcher")
				researcher := model.Researcher{ResearcherID: userDB.UserID}
				db.Create(&researcher)
				//global.Tpl.ExecuteTemplate(w, "admin-new-user-generic", userDB)
				http.Redirect(w, r, "/admin", http.StatusFound)
				displayAdmin(w, r)

			default:
				fmt.Println("Not sure your type")
				global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", userDB)
			}

		} else {
			// validation failed
			m := map[string]interface{}{
				"error": err,
			}
			global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", m)
		}


	} else {
		// add to the err - email already taken
		m := map[string]interface{}{
			"error": "Email Already Taken",
		}
		global.Tpl.ExecuteTemplate(w, "viewNewUserAdmin", m)
		//return
	}
}


func createStudent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formEmail := r.FormValue("email")
	credits,_ := strconv.Atoi(r.FormValue("credits"))
	studentType, _ := strconv.Atoi(r.FormValue("student-type"))
	mu := model.MainUser{}
	db.Where(&model.MainUser{UserEmail: formEmail}).First(&mu)
	stu := model.Student{StudentID:mu.UserID, StudentType:studentType}
	db.Create(&stu)

	if studentType == 1{
		stuFT := model.FullTimeStudent{FullTimeStudentID: stu.StudentID, NumCredits:credits}
		db.Create(&stuFT)
	} else {
		stuPT := model.PartTimeStudent{PartTimeStudentID: stu.StudentID, NumCredits:credits}
		db.Create(&stuPT)
	}
	http.Redirect(w,r,"/admin", http.StatusFound)
	displayAdmin(w,r)
}

func createFaculty(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formEmail := r.FormValue("email")
	mu := model.MainUser{}
	db.Where(&model.MainUser{UserEmail: formEmail}).First(&mu)
	facultyType, _ := strconv.Atoi(r.FormValue("faculty-type"))
	println("Faculty type num: ", facultyType)
	department, _ := strconv.ParseUint(r.FormValue("department"), 10, 64)
	faculty := model.Faculty{FacultyID: mu.UserID, FacultyType: facultyType, DepartmentID: uint(department)}
	db.Create(&faculty)
	if (facultyType == 1) {
		facultyFT := model.FullTimeFaculty{FullTimeFacultyID: faculty.FacultyID}
		db.Create(&facultyFT)
	} else if (facultyType == 2) {
		facultyPT := model.PartTimeStudent{PartTimeStudentID: faculty.FacultyID}
		db.Create(&facultyPT)
	}
	http.Redirect(w, r, "/admin", http.StatusFound)
	displayAdmin(w, r)

}

func searchUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Include user type information with user search results
	queryVals := r.URL.Query()
	emailQuery, _ := queryVals["email"]
	firstNameQuery, _ := queryVals["first-name"]
	lastNameQuery, _ := queryVals["last-name"]

	email := "N"
	firstName := "N"
	lastName := "N"

	if len(emailQuery) < 1 {
		println("No email given")
		email = "-"
	} else {
		email = emailQuery[0]
	}

	if len(firstNameQuery) < 1 {
		println("No FirstNae given")
		firstName = "-"
	} else {
		firstName = firstNameQuery[0]
	}

	if len(lastNameQuery) < 1 {
		println("No lastName given")
		lastName = "-"
	} else {
		lastName = lastNameQuery[0]
	}

	println("Query From Form- Email: ", email, " FName: ", firstName, " LName: ", lastName)
	users := []model.MainUser{}
	db.Where("first_name LIKE ? OR last_name LIKE ?",firstName, lastName).Or(model.MainUser{UserEmail: email}).Find(&users)
	for _, v := range users {
		fmt.Println("UserEmail", v.UserEmail)
	}

	data :=  map[string]interface{}{
		"Users": users,
	}
	global.Tpl.ExecuteTemplate(w, "searchUsersAdmin", data)
}


func deleteUser(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	params := mux.Vars(r)
	formEmail,_ := strconv.Atoi(params["userID"])
	//formEmail := r.URL.Query().Get("userID")
	println("UserID coming in is--" + params["userID"])
	mu := model.MainUser{}
	db.Where(&model.MainUser{UserID: uint(formEmail)}).First(&mu)
	println("Fropm DB found user email: " + mu.UserEmail)
	if mu.UserID != 0 {
		userType := mu.UserType
		if userType == 1 {
			student := model.Student{}
			db.First(&student, mu.UserID)
			if student.StudentType == 1 && student.StudentID != 0 {
				studentFT := model.FullTimeStudent{}
				db.First(&studentFT, student.StudentID)
				if studentFT.FullTimeStudentID != 0 {
					println("Deleting Full Time Student")
					db.Delete(&studentFT)
				} else {
					println("FT student not found")
				}
				println("Deleting Student")
				db.Delete(&student)
			} else if student.StudentType == 2 && student.StudentID != 0 {
				studentPT := model.PartTimeStudent{}
				db.First(&studentPT, student.StudentID)
				if studentPT.PartTimeStudentID != 0 {
					println("Deleting Part Time Student")
					db.Delete(&studentPT)
				} else {
					println("Part time student not found")
				}
				println("Deleting Student")
				db.Delete(&student)
			} else {
				println("Student not found")
			}

		} else if userType == 2 {
			faculty := model.Faculty{}
			db.First(&faculty, mu.UserID)
			if faculty.FacultyType == 1 && faculty.FacultyID != 0 {
				studentFT := model.FullTimeStudent{}
				db.First(&studentFT, faculty.FacultyID)
				if studentFT.FullTimeStudentID != 0 {
					println("Deleting Full Time Faculty")
					db.Delete(&studentFT)
				} else {
					println("FT faculty not found")
				}
				println("Deleting faculty")
				db.Delete(&faculty)
			} else if faculty.FacultyType == 2 && faculty.FacultyID != 0 {
				facultyPT := model.PartTimeFaculty{}
				db.First(&facultyPT, faculty.FacultyID)
				if facultyPT.PartTimeFacultyID != 0 {
					println("Deleting Part Time Student")
					db.Delete(&facultyPT)
				} else {
					println("Part time faculty not found")
				}
				println("Deleting faculty")
				db.Delete(&faculty)
			} else {
				println("Faculty not found")
			}

		} else if userType == 3 {
			admin := model.Admin{}
			db.First(&admin, mu.UserID)

			if admin.AdminID != 0 {
				db.Delete(&admin)
			} else {
				println("Admin not found")
			}
			println("Deleting msin ser")
		} else if userType == 4 {
			researcher := model.Researcher{}
			db.First(&researcher, mu.UserID)

			if researcher.ResearcherID != 0 {
				db.Delete(&researcher)
			} else {
				println("Researcher not found")
			}
		}

		println("Deleting main user")
		db.Delete(&mu)
	} else {
		println("Main user not found")
	}
	data :=  map[string]interface{}{
		"deleted": mu,
	}
	global.Tpl.ExecuteTemplate(w, "searchUsersAdmin", data)
}


func searchMasterScheduleForm(w http.ResponseWriter, r *http.Request){
	allDepartments := []model.Department{}
	db.Find(&allDepartments)
	m :=  map[string]interface{}{
		"Departments": allDepartments,
	}
	global.Tpl.ExecuteTemplate(w, "masterScheduleSearch", m)
	}



func searchMasterSchedule(w http.ResponseWriter, r *http.Request){

	println("Inside searchMasterSchedule")

	queryVals := r.URL.Query()

	departmentQuery,_ := queryVals["department"]
	courseNameQuery,_ := queryVals["course-name"]
	professorQuery := queryVals["instructor"]
	day := queryVals["day"]

	depID := departmentQuery[0]
	courseName := courseNameQuery[0]
	professor := professorQuery[0]
	dayID := day[0]

	whereMap := make(map[string]interface{})

	whereStuff := "WHERE "

	numQueries := 0

	if depID != "" {
		println("Department query present: " + depID)
		//depID, _ := strconv.ParseUint(departmentQuery[0], 10, 64)
		whereMap["department_id"] = depID
		whereStuff += "department_id = " + depID
		numQueries++
	}
	if courseName != "" {
		whereMap["course_name"] = courseName
		if numQueries == 0 {
			whereStuff += " course_name = '" + courseName + "'"

		}else {
			whereStuff += " AND course_name = '" + courseName + "'"

		}
		numQueries++
	}

	if dayID != ""{
		if numQueries == 0 {
			whereStuff += " day.day_id = " + dayID
		} else {
			whereStuff += " AND day.day_id = " + dayID
		}
		numQueries++
	}

	if professor != "" {
		prof := strings.Split(professor, " ")
		if numQueries == 0 {
			whereStuff += " first_name = '" + prof[0] + "'"
			whereStuff += " AND last_name = '" + prof[1] + "'"

		}else {
			whereStuff += " AND first_name = '" + prof[0] + "'"
			whereStuff += " AND last_name = '" + prof[1] + "'"
		}
		numQueries++

	}

	//registering for next semester
	if numQueries == 0 {
		whereStuff += " semester.year = 2018 AND semester.season = 'Spring'"
	}else {
		whereStuff += " AND semester.year = 2018 AND semester.season = 'Spring'"
	}

	type IsAdmin struct {
		IsAdmin bool
		User model.MainUser
	}

		type CourseData struct {
			CourseName string
			CourseCredits int
			CourseDescription string
			DepartmentID uint
			SectionID uint
			CourseSectionNumber int
			CourseID uint
			FacultyID uint
			FirstName string
			LastName string
			TimeSlotID uint
			LocationID uint
			DayID uint
			MeetingDay string
			RoomID uint
			RoomNumber string
			RoomType string
			BuildingID uint
			BuildingName string
			Capacity uint
			Time string
			Prerequisites []model.Course
			User IsAdmin
			NumEnrolled int
		}

	//coursesFound := []model.Course{}
	//db.Where(model.Course{CourseName: courseName}).Or(model.Course{DepartmentID: uint(depID)}).Or(model.Course{CourseName: courseName}).Find(&coursesFound)

	queryRes := []CourseData{}


	//rows, err := db.Joins("JOIN course ON course.course_id = section.course_id").Where(whereMap).Rows()
	sql := `SELECT course.course_name, course.course_credits, course.course_description, course.department_id, section.section_id, section.course_section_number,
	section.course_id, section.faculty_id, section.time_slot_id, section.location_id, section.course_section_number,
	main_user.first_name, main_user.last_name,
	day.meeting_day, day.day_id,
	building.building_name,time,
	room.room_number, room.room_type,capacity

	FROM section
	JOIN course ON course.course_id = section.course_id
	JOIN main_user ON main_user.user_id = section.faculty_id
	JOIN location ON section.location_id = location.location_id
	JOIN building ON building.building_id = location.building_id
	JOIN room ON room.room_id = location.room_id
	JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	JOIN semester ON time_slot.semester_id = semester.semester_id
	JOIN day ON time_slot.day_id = day.day_id
	JOIN period ON period.period_id = time_slot.period_id `

	sql += whereStuff

	if strings.Contains(whereStuff,";&"){
		fmt.Println("escape")
	}
	fmt.Println("QUery to be run is", sql)
	db.Raw(sql).Scan(&queryRes)


		//if err == nil{
		//	//rows.Scan(&queryRes)
		//
		//} else {
		//	println(err.Error())
		//}
	/*
	for _, val := range queryRes{
		println(val.CourseName)
	}
	*/
	fmt.Println(queryRes)
	allDepartments := []model.Department{}

	db.Find(&allDepartments)

	coursesFoundPrereqs := make(map[string][]model.Course)

	// get prereqs for every unique course in queryRes
	for qIndex, val := range queryRes{
		if _, present := coursesFoundPrereqs[val.CourseName]; !present {
			course := model.Course{}
			db.First(&course, val.CourseID)
			prereqs := course.FindCoursePrerequisites(db)
			coursesFoundPrereqs[val.CourseName] = prereqs
			queryRes[qIndex].Prerequisites = prereqs
			println("Adding prereqs for: " + course.CourseName)
			for _,c := range val.Prerequisites{
				println("-prereq::" + c.CourseName)
			}
		} else {
			queryRes[qIndex].Prerequisites = coursesFoundPrereqs[val.CourseName]
		}
	}

	//chosenDep := model.Department{}
	//db.First(&chosenDep, depID)
	searchParams := map[string]string{
		//"Department": chosenDep.DepartmentName,
		"Professor": professor,
		"CourseName": courseName,
		//"CourseNum": courseNum,
	}

	_,user := CheckLoginStatus(w,r)



	//courseEnrollment := make(map[int]int)
	for k, _ := range queryRes {
		count := 0
		db.Table("enrollment").Select("*").Where("section_id = ?", queryRes[k].SectionID).Count(&count)
		queryRes[k].NumEnrolled = count
	}


	admin := IsAdmin{}

	if user.UserType == 3 {
		admin.IsAdmin = true
		admin.User = user
	}else{
		admin.IsAdmin = false
	}

	type IsStudent struct {
		IsStudent bool
		User model.MainUser
	}

	student := IsStudent{}

	if user.UserType == 1 {
		student.IsStudent = true
		student.User = user
	}else {
		student.IsStudent = false
	}


	data :=  map[string]interface{}{
		"Results": queryRes,
		"Departments": allDepartments,
		"Params": searchParams,
		"User":admin,
		"Student": student,
	}

	err := global.Tpl.ExecuteTemplate(w, "masterScheduleSearch", data)

	if err != nil {
		fmt.Println(err.Error())
	}

}


func getStudentsReportByGrade(w http.ResponseWriter, r *http.Request){
	allDepartments := []model.Department{}
	db.Find(&allDepartments)
	allCourses := []model.Course{}
	db.Find(&allCourses)
	m := map[string]interface{}{
		"Departments": allDepartments,
		"Courses": allCourses,
	}
	global.Tpl.ExecuteTemplate(w, "researchStudentsByGrade", m)
}

func genReportStudentsByGrade(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	formGradeLow := r.FormValue("letter-grade-boundlow")
	formGradeHigh := r.FormValue("letter-grade-boundhigh")
	courseNum := r.FormValue("department")

	grades := []string{"F","C-","C","C+","B-","B", "B+", "A-", "A"}

	getIndex := func(vs []string, t string) int {
		for i, v := range vs {
			if v == t {
				return i
			}
		}
		return -1
	}

	indexLow := getIndex(grades, formGradeLow)
	indexHigh := getIndex(grades, formGradeHigh)
	println("gradeLow:" + formGradeLow + ", gradeHigh: " + formGradeHigh)
	println("IndexLow:" + strconv.Itoa(indexLow) + ", indexHigh: " + strconv.Itoa(indexHigh))
	gradeSlice := grades[indexLow:indexHigh+1]

	println("Seaching for grades between "+ grades[indexLow] + " and " + grades[indexHigh])
	for g := range gradeSlice {
		println(g)
	}

	contains := func(slice []string, item string) bool {
		set := make(map[string]struct{}, len(slice))
		for _, s := range slice {
			set[s] = struct{}{}
		}
		_, ok := set[item]
		return ok;
	}

	productsSelected := r.Form["filter-options"]

	sql := `SELECT student_history.grade, main_user.first_name, main_user.last_name
	FROM student_history
	JOIN enrollment ON student_history.enrollment_id = enrollment.enrollment_id
	JOIN section ON enrollment.section_id = section.section_id
	JOIN student ON student_history.student_id = student.student_id
	JOIN main_user ON main_user.user_id = student.student_id
	WHERE
	student_history.grade IN (?)
	AND student_history.status = ?
	AND section.course_id = ? `

	//if contains(productsSelected, "major"){
	//	major := r.FormValue("major")
	//	sqlDepartmentFilter := ` AND student.student_id IN
	//	(SELECT student_id FROM student_major
	//	JOIN major ON student_major.major_id = major.major_id
	//	WHERE major.department_id = `+ major + " )"
	//
	//	sql += sqlDepartmentFilter
	//}

	if contains(productsSelected, "student-type"){
		stuType := r.FormValue("full-or-part")
		println("Filtering for student type: " + stuType)
		sqlDepartmentFilter := ` AND student.student_type = `+ stuType
		sql += sqlDepartmentFilter
	}

	type StudentData struct {
		FirstName string
		LastName string
		Grade string
	}
	records := []StudentData{}
	db.Raw(sql, gradeSlice, "Complete", courseNum).Scan(&records)

	modtime := time.Now()
	filepath := "Record" + strconv.Itoa(modtime.Nanosecond()) + ".csv"
	outfile, err := os.Create("./"+filepath)

	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer outfile.Close()
	writer := struct2csv.NewWriter(outfile)

	for _, record := range records {
		if err := writer.WriteStruct(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		} else {
			println("FNAME: " + record.FirstName + ", LNAME: " + record.LastName + ", GRADE: " + record.Grade)
		}
	}

	writer.Flush()
	// tell the browser the returned content should be downloaded
	w.Header().Set("Content-Type", "text/csv")
	path := "./"+filepath
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath)

	http.ServeFile(w,r,path)
	global.Tpl.ExecuteTemplate(w, "researchStudentsByGrade", nil)
	os.Remove(path)
}
