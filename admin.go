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
)



/* Student Actions */

func ViewStudentSchedulePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "viewStudentScheduleAdmin", user)
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
	errorMessage string
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
	} else {
		db.Where(&model.MainUser{UserEmail: email}).Find(&user).Count(&count)
	}

	if count > 0 {
		fmt.Println("You have a user", user.FirstName)
	} else {//TODO: Render error correctly, try passing in map with admin user
		global.Tpl.ExecuteTemplate(w,"adminViewStudentScheduleDetails", "No user found")
		return
	}

	//Successfully gets the student

	//email := user.UserEmail

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
	/*
	db.Table("enrollment").Select("course_name,course_credits," +
		//"room_number,building_name,start_time,end_time,meeting_day" +
		"").Joins("JOIN section on section.section_id = enrollment.section_id AND student_id = ?" +
		"", id).Joins("JOIN section.course_id = course.course_id").First(&ss)
	*/
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

	//fmt.Println("Results from query are", results)

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

	/*
	results := model.StudentHolds{} //should be student-holds

	db.Table("student").Select("*").Joins("join student_holds on student.student_id = holds.student_id WHERE student_id = ?", intId).Scan(&results)

	fmt.Println("Hold id is:",results.HoldID, "Student id:", results.StudentID)
	hold := model.Hold{}

	db.Where(model.Hold{HoldID:results.HoldID}).First(&hold)

	hd := HoldDetail{StudentName:user.FirstName, HoldName:hold.HoldName}
	global.Tpl.ExecuteTemplate(w, "adminStudentHold" , hd)
	*/
	//hd := []HoldDetail{}
	hs := []model.Hold{}

	db.Raw("SELECT hold_name FROM student NATURAL JOIN student_holds NATURAL JOIN hold WHERE student.student_id =?", user.UserID).Scan(&hs)
	//fmt.Println(hd)

	m := map[string]interface{}{
		"User": user,
		"Holds": hs,
	}

	global.Tpl.ExecuteTemplate(w, "adminStudentHold", m)


}



func AdminAddCoursePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "addCourseAdmin", user)
	}
}