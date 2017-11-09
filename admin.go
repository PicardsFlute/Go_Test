package main

import (
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_"Starfleet/memory"
	"net/http"
	"Starfleet/global"
	//"github.com/gorilla/mux"
	"Starfleet/model"
	"strconv"
	"fmt"
	"time"
	"github.com/gorilla/mux"
)



/* Student Actions */

func ViewStudentSchedulePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		m := map[string]interface{}{
			"User":user,
		}
		err := global.Tpl.ExecuteTemplate(w, "viewStudentScheduleAdmin", m)
		//err := global.Tpl.ExecuteTemplate(w, "viewStudentScheduleAdmin", user)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

type StudentSchedule struct {
	CourseName string
	CourseCredits string
	RoomNumber string
	BuildingName string
	StartTime time.Time
	EndTime time.Time
	MeetingDay string
}

type HoldDetail struct {
	HoldName string
}

type NoUser struct {
	ErrorMessage string
}

func ViewStudentSchedule(w http.ResponseWriter, r *http.Request){

	//vars := mux.Vars(r) //returns a mapping responses
	//personId := vars["student"] //get map with key id number

	user := model.MainUser{}

	email := r.FormValue("email")
	id := r.FormValue("id")
	count := 0
	//major := r.FormValue("major")

	//db.Where("id = ?", id).Find(&model.Enrollment{})
	intId, err := strconv.Atoi(id)
	if err != nil {
		err.Error()
	}


	//first check if they entered an ID
	if id != "" {
		db.Where(&model.MainUser{UserID: uint(intId)}).Find(&user).Count(&count)
	}

	if email != ""{
		db.Where(&model.MainUser{UserEmail: email}).Find(&user).Count(&count)
	}
	//nu := NoUser{"Nobody found"}
	if count > 0 {
		fmt.Println("You have a user", user.FirstName)
	} else {//TODO: Render error correctly, try passing in map with admin user

		m := map[string]interface{}{
			"NoUser": "Nobody Home",
		}

		fmt.Println("Showing no users found error")
		global.Tpl.ExecuteTemplate(w,"viewStudentScheduleAdmin", m)
		return
	}



	//this correctly joins enrollment and section on section_id, could possible store all the data in an empty interface
	/*
	rows, err := db.Table("enrollment").Select("*").Joins("join section on section.section_id = enrollment.section_id  AND student_id = ?", id).Rows()
	if err != nil{
		fmt.Println(err.Error())
	}else {
		for rows.Next(){
			fmt.Println("found rows", rows)
		}
	}
	*/
	ss := []StudentSchedule{}
	db.Raw(`SELECT course_name,course_credits,building_name,room_number,meeting_day,start_time,end_time
	FROM enrollment
	NATURAL JOIN Section
	NATURAL JOIN time_slot
	NATURAL JOIN course
	NATURAL JOIN period
	NATURAL JOIN day NATURAL
	JOIN location NATURAL JOIN building
	NATURAL JOIN room WHERE enrollment.student_id =?`,user.UserID).Scan(&ss)
	fmt.Println(ss)

	global.Tpl.ExecuteTemplate(w, "adminViewStudentScheduleDetails", ss)

}


func ViewStudentHoldsPage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)

	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "viewStudentHoldsAdmin", user)
	}else {
		http.Redirect(w,r,"/",http.StatusForbidden)
	}
}


func ViewStudentHolds (w http.ResponseWriter,r *http.Request) {



	user := model.MainUser{}

	email := r.FormValue("email")
	id := r.FormValue("id")
	count := 0
	//major := r.FormValue("major")

	//db.Where("id = ?", id).Find(&model.Enrollment{})
	intId, err := strconv.Atoi(id)
	if err != nil {
		err.Error()
	}


	//first check if they entered an ID
	if id != "" {
		db.Where(&model.MainUser{UserID: uint(intId)}).Find(&user).Count(&count)
	} else {
		db.Where(&model.MainUser{UserEmail: email}).Find(&user).Count(&count)
	}

	if count > 0 {
		fmt.Println("You have a user", user.FirstName)
	} else {
		global.Tpl.ExecuteTemplate(w,"viewStudentHoldsAdmin", "No user found")
		return
	}

	hs := []model.Hold{}

	db.Raw("SELECT * FROM student NATURAL JOIN student_holds NATURAL JOIN hold WHERE student.student_id =?", user.UserID).Scan(&hs)
	//fmt.Println(hd)

	m := map[string]interface{}{
		"User": user,
		"Holds": hs,
	}
	errTpl := global.Tpl.ExecuteTemplate(w, "adminStudentHold", m)
	if errTpl != nil {
		fmt.Println(errTpl.Error())
	}


}


func AdminDeleteHold(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	holdId := vars["id"]
	user := vars["user"]

	holdInt , err := strconv.Atoi(holdId)
	if err != nil {
		fmt.Println(err.Error())
	}
	userInt , err := strconv.Atoi(user)
	if err != nil {
		fmt.Println(err.Error())
	}

		fmt.Println("StudentID =", userInt, "HoldID =", holdInt, "User =", user)
	//TODO: use ajax to delete the hold, get the id of the hold, and get the student id from a different portion of the page by using attribute
	studentHold := model.StudentHolds{}
	//db.Where("student_id = ? AND hold_id = ?", userInt,holdInt).First(&studentHold)
	//db.Delete(&hold)
	db.Raw("SELECT * FROM student_holds WHERE student_id = ? AND hold_id = ?", userInt,holdInt).Scan(&studentHold)
	fmt.Println("Hold found", studentHold)
	db.Delete(&studentHold)
	//fmt.Println("Hold deleted sucessfully")


}



func AdminAddCoursePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "addCourseAdmin", user)
	}
}

func AdminAddCourse(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if !isLogged || user.UserType != 3{
		http.Redirect(w, r, "/", http.StatusForbidden)
	}

	courseName := r.FormValue("name")
	courseCredits := r.FormValue("credits")
	courseDescription := r.FormValue("description")
	courseDepartment := r.FormValue("department")


	intCredits,err := strconv.Atoi(courseCredits)
	if err != nil {
		fmt.Println(err.Error())
	}

	intDepartment, err := strconv.Atoi(courseDepartment)
	if err != nil {
		fmt.Println(err.Error())
	}

	course := model.Course{CourseName:courseName, CourseCredits:intCredits,
	CourseDescription:courseDescription, DepartmentID:uint(intDepartment)}

	fmt.Println("Course info is", course)
	db.Create(&course)

}

func AdminSearchCoursePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "addCourseAdmin", user)
	}
}

type CourseOptions struct {
	CourseName string
	CourseID uint
}



func AdminAddSectionPage(w http.ResponseWriter, r *http.Request){
	courses := []CourseOptions{}
	fac := []model.MainUser{}

	//db.Raw("SELECT user_id,first_name,last_name FROM main_user WHERE user_type= ?", 2).Scan(&fac)

	db.Table("course").Select("course_name, course_id").Scan(&courses)

	db.Table("main_user").Select("*").Where("user_type = ?",2).Scan(&fac)

	fmt.Println("Courses are", courses, "Faculty are", fac)

	m := map[string]interface{}{
		"Courses": courses,
		"Faculty": fac,
	}
	errTpl := global.Tpl.ExecuteTemplate(w, "addSectionAdmin", m)
	if errTpl != nil {
		fmt.Println(errTpl.Error())
	}

}