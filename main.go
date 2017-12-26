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


	//dbPassword := os.Getenv("PG_DATABASE_PW")
	//db, err = gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password="+dbPassword)

	//for heroku
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))


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
	routes.HandleFunc("/logout", logout)

	/* Master schedule searching, public functions */
	routes.HandleFunc("/course", searchMasterScheduleForm).Methods("GET")
	routes.HandleFunc("/course/", searchMasterSchedule).Methods("GET")


	/*Admin routes */
	
	routes.Handle("/admin",  checkAdmin(displayAdmin)).Methods("GET")
	routes.HandleFunc("/admin/student" , ViewStudentSchedulePage).Methods("GET")
	routes.HandleFunc("/admin/student/{student}", ViewStudentSchedule).Methods("GET")

	routes.HandleFunc("/admin/transcript", viewStudentTranscriptPage).Methods("GET")
	routes.HandleFunc("/admin/transcript/{student}", viewStudentTranscript).Methods("GET")

	routes.Handle("/admin/holds", checkAdmin(ViewStudentHoldsPage))
	routes.Handle("/admin/holds/{user}", checkAdmin(ViewStudentHolds)).Methods("GET")
	routes.Handle("/admin/holds/{user}/{id}", checkAdmin(AdminDeleteHold)).Methods("POST")

	routes.Handle("/admin/course",checkAdmin(AdminAddCoursePage))
	routes.Handle("/admin/course/{course}",checkAdmin(AdminAddCourse)).Methods("POST")
	routes.Handle("/admin/courses/prereq",checkAdmin(AddCoursePreRequisit)).Methods("POST")

	routes.Handle("/admin/course/search", checkAdmin(AdminSearchCoursePage)).Methods("GET")
	routes.Handle("/admin/course/search/{course}", checkAdmin(AdminSearchCourse)).Methods("GET")
	routes.Handle("/course/update", checkAdmin(UpdateCourse)).Methods("POST")
	//routes.HandleFunc("/admin/course/", AdminDeleteCourse)
	routes.HandleFunc("/admin/section", AdminAddSectionPage)
	routes.HandleFunc("/section/update", AdminUpdateSectionForm).Methods("GET")
	routes.HandleFunc("/section/update", AdminUpdateSection).Methods("POST")

	routes.Handle("/admin/section/{section}", checkAdmin(AdminAddSection))
	routes.Handle("/admin/section/room/{id}", checkAdmin(GetRoomsForBuilding))
	routes.Handle("/admin/section/department/{id}", checkAdmin(GetDepartmentsForSections)).Methods("GET")
	routes.Handle("/admin/user", checkAdmin(newUserForm)).Methods("GET")
	routes.Handle("/admin/user", checkAdmin(createUser)).Methods("POST")
	routes.Handle("/admin/user/student", checkAdmin(createStudent)).Methods("POST")
	routes.Handle("/admin/user/faculty", checkAdmin(createFaculty)).Methods("POST")
	routes.Handle("/admin/user/search" , checkAdmin(searchUser)).Methods("GET")
	routes.Handle("/admin/user/{userID}/delete", checkAdmin(deleteUser)).Methods("POST")
	routes.Handle("/admin/semester" , checkAdmin(changeSemesterStatusForm)).Methods("GET")
	routes.Handle("/admin/semester" , checkAdmin(changeSemesterStatus)).Methods("POST")
	routes.Handle("/admin/time", checkAdmin(addSectionTime)).Methods("POST")


	/* Student Routes*/
	routes.Handle("/student",  checkStudent(displayStudent)).Methods("GET")
	routes.Handle("/student/schedule", checkStudent(ViewSchedule)).Methods("GET")
	routes.Handle("/student/registered", checkStudent(ViewRegisteredCourses)).Methods("GET")
	routes.Handle("/student/registered", checkStudent(DropRegisteredCourse)).Methods("POST")
	routes.Handle("/student/holds", checkStudent(ViewHolds)).Methods("GET")
	routes.Handle("/student/advisor", checkStudent(ViewAdvisor)).Methods("GET")
	routes.Handle("/student/transcript", checkStudent(ViewTranscript)).Methods("GET")
	routes.Handle("/student/search", checkStudent(AddCoursePage)).Methods("GET")
	routes.Handle("/student/register", checkStudent(RegisterForSection)).Methods("POST")

	/*Faculty Routes */
	routes.Handle("/faculty",  checkSessionWrapper(displayFaculty)).Methods("GET")
	routes.Handle("/faculty/schedule", checkFaculty(facultyViewSchedule)).Methods("GET")
	routes.Handle("/faculty/grades", checkFaculty(giveStudentGradesPage)).Methods("GET")
	routes.Handle("/faculty/grades/{sectionID}", checkFaculty(giveStudentGradesForm)).Methods("GET")
	routes.Handle("/faculty/grades/{sectionID}", checkFaculty(submitGrades)).Methods("POST")


	routes.Handle("/researcher", checkSessionWrapper(displayResearcher)).Methods("GET")
	routes.Handle("/researcher/students/grades", checkResearcher(getStudentsReportByGrade)).Methods("GET")
	routes.Handle("/researcher/students/grades", checkResearcher(genReportStudentsByGrade)).Methods("POST")




	// USED FOR HEROKU
	http.ListenAndServe(":" + os.Getenv("PORT"),handlers.LoggingHandler(os.Stdout,routes))
	//http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout,routes))

	//defer db.Close(), want to keep db connectioT"), routes)

	//USED FOR LOCAL, only use onen open


}


// routes for site


