package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
	//"Starfleet/model"

	"github.com/gorilla/mux"
	"strconv"
)

func ViewSchedule(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]
	idInt , _ := strconv.Atoi(id)
	ss := []StudentSchedule{}


	db.Raw(`SELECT course_name,course_credits,building_name,room_number,meeting_day,time
	FROM enrollment
	NATURAL JOIN Section
	NATURAL JOIN time_slot
	NATURAL JOIN course
	NATURAL JOIN period
	NATURAL JOIN day NATURAL
	JOIN location NATURAL JOIN building
	NATURAL JOIN room WHERE enrollment.student_id =?`,idInt).Scan(&ss)
	fmt.Println(ss)

	global.Tpl.ExecuteTemplate(w, "ViewStudentScheduleDetails", ss)

}
