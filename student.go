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


	type StudentSchedule struct {
		CourseName string
		CourseCredits string
		RoomNumber string
		BuildingName string
		Time string
		MeetingDay string
		FirstName string
		LastName string
	}


	_, user := CheckLoginStatus(w,r)
	ss := []StudentSchedule{}
	db.Raw(`SELECT student_history.student_id,course_name,course_credits,building_name,room_number,meeting_day, first_name, last_name, time,student_history.status
	FROM student_history
	JOIN enrollment ON student_history.enrollment_id = enrollment.enrollment_id
	JOIN section ON enrollment.section_id = section.section_id
	JOIN course ON course.course_id = section.course_id
	JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	JOIN semester ON time_slot.semester_id = semester.semester_id
	JOIN period ON time_slot.period_id = period.period_id
	JOIN day ON time_slot.day_id = day.day_id
	JOIN location ON section.location_id = location.location_id
	JOIN building ON location.building_id = building.building_id
	JOIN room ON location.room_id = room.room_id
	JOIN faculty ON section.faculty_id = faculty.faculty_id
	JOIN main_user ON faculty.faculty_id = main_user.user_id
	WHERE enrollment.student_id = ? AND student_history.status = 'In progress'`, user.UserID).Scan(&ss)
	fmt.Println(ss)

	err := global.Tpl.ExecuteTemplate(w, "ViewStudentScheduleDetails", ss)
	if err != nil {
		fmt.Println(err.Error())
	}
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
