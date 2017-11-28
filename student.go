package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
	//"Starfleet/model"

	_"github.com/gorilla/mux"
	_"strconv"
	"Starfleet/model"
)

func ViewSchedule(w http.ResponseWriter, r *http.Request){
	/*
	vars := mux.Vars(r)
	id := vars["id"]
	idInt , _ := strconv.Atoi(id)
	*/

	_, user := CheckLoginStatus(w,r)
	ss := []StudentSchedule{}
	db.Raw(`SELECT course_name,course_credits,building_name,room_number,meeting_day,time
	FROM enrollment
	NATURAL JOIN Section
	NATURAL JOIN time_slot
	NATURAL JOIN course
	NATURAL JOIN period
	NATURAL JOIN day NATURAL
	JOIN location NATURAL JOIN building
	NATURAL JOIN room WHERE enrollment.student_id =?`,user.UserID).Scan(&ss)
	fmt.Println(ss)

	global.Tpl.ExecuteTemplate(w, "ViewStudentScheduleDetails", ss)

}

func ViewHolds(w http.ResponseWriter, r *http.Request){

	_ ,user := CheckLoginStatus(w,r)
	fmt.Println("Current user is", user)
	hs := []model.Hold{}

	db.Raw("SELECT * FROM student NATURAL JOIN student_holds NATURAL JOIN hold WHERE student.student_id =?", user.UserID).Scan(&hs)

	m := map[string]interface{}{
		"User": user,
		"Holds": hs,
	}
	errTpl := global.Tpl.ExecuteTemplate(w, "StudentHold", m)
	if errTpl != nil {
		fmt.Println(errTpl.Error())
	}

}
