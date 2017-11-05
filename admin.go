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
)



/* Student Actions */

func ViewStudentSchedulePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "viewStudentScheduleAdmin", user)
	}
}

func ViewStudentSchedule(w http.ResponseWriter, r *http.Request){

	//vars := mux.Vars(r) //returns a mapping responses
	//personId := vars["student"] //get map with key id number

	user := model.MainUser{}

	//name := r.FormValue("name")
	id := r.FormValue("id")
	//major := r.FormValue("major")

	//db.Where("id = ?", id).Find(&model.Enrollment{})
	intId, err := strconv.Atoi(id)
	if err != nil {
		err.Error()
	}

	//first check if they entered an ID
	if id != "" {
		db.Where(&model.MainUser{UserID: uint(intId)}).Find(&user)
	}

	if user.FirstName != "" {
		fmt.Println("You have a user", user.FirstName)
	} else {
		fmt.Println("Error searching user", user)
	}

	//Successfully gets the student

	//email := user.UserEmail

	//TODO: Check their enrollment by joining enrollment ID and secttion ID
	//TODO: Then get section details for day, semester, and period info
	//TODO: Then join course ID in  section and course to get the course name,description, and credits

	history := model.StudentHistory{}

	//results := model.Enrollment{}


	//listHistory := make([]model.StudentHistory, 10)

	db.Where(&model.StudentHistory{StudentID: uint(intId)}).Find(&history)

	/*

	for _,v := range listHistory {
		fmt.Println("Student history is", v.StudentID)
	}
	*/

	//this correctly joins enrollment and section on section_id, could possible store all the data in an empty interface
	rows, err := db.Table("enrollment").Select("*").Joins("join section on section.section_id = enrollment.section_id  AND student_id = ?", id).Rows()
	if err != nil{
		fmt.Println(err.Error())
	}else {
		for rows.Next(){
			fmt.Println("found rows", rows)
		}
	}


	//fmt.Println("Results from query are", results)

	//global.Tpl.ExecuteTemplate(w, "studentHistory", history)

}


func ViewStudentHoldsPage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "viewStudentHoldsAdmin", user)
	}
}


func ViewStudentHolds (w http.ResponseWriter,r *http.Request) {

	user := model.MainUser{}

	//name := r.FormValue("name")
	id := r.FormValue("id")
	//major := r.FormValue("major")

	//db.Where("id = ?", id).Find(&model.Enrollment{})
	intId, err := strconv.Atoi(id)
	if err != nil {
		err.Error()
	}

	//first check if they entered an ID
	if id != "" {
		db.Where(&model.MainUser{UserID: uint(intId)}).Find(&user)
	}

	if user.FirstName != "" {
		fmt.Println("You have a user", user.FirstName)
	} else {
		fmt.Println("Error searching user", user)
	}

	results := model.Student{} //should be student-holds

	db.Table("student").Select("*").Joins("join student_holds on student.studentID = holds.studentID WHERE student.id = ?", intId).Scan(&results)
	//TODO: No holds table in db or model in project
}



func AdminAddCoursePage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)
	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "addCourseAdmin", user)
	}
}