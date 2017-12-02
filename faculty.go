package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
	"github.com/gorilla/mux"
	"strconv"
)

func facultyViewSchedule(w http.ResponseWriter,r *http.Request){
	_, user := CheckLoginStatus(w,r)
	type FacultySchedule struct{
		CourseName string
		BuildingName string
		RoomNumber string
		MeetingDay string
		Time string
		Year int
		Season string
	}

	facultySchedule := []FacultySchedule{}
	db.Raw(`
	SELECT course_name, building_name,room_number,meeting_day,time,year,season
	FROM section
	JOIN course ON course.course_id = section.course_id
	JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	JOIN semester ON time_slot.semester_id = semester.semester_id
	JOIN period ON time_slot.period_id = period.period_id
	JOIN day ON time_slot.day_id = day.day_id
	JOIN location ON section.location_id = location.location_id
	JOIN building ON location.building_id = building.building_id
	JOIN room ON location.room_id = room.room_id
	WHERE section.faculty_id = ? AND semester.year = 2018 AND semester.season = 'Spring'`, user.UserID).Scan(&facultySchedule)
	fmt.Println(facultySchedule)

	err := global.Tpl.ExecuteTemplate(w, "ViewFacultyScheduleDetails", facultySchedule)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func giveStudentGradesPage(w http.ResponseWriter, r *http.Request){
	_, user := CheckLoginStatus(w,r)
	type FacultyScheduleGrade struct{
		CourseName string
		SectionID uint
		BuildingName string
		RoomNumber string
		MeetingDay string
		Time string
		Year int
		Season string
	}

	facultySchedule := []FacultyScheduleGrade{}
	db.Raw(`
	SELECT course_name,section.section_id,building_name,room_number,meeting_day,time,year,season
	FROM section
	JOIN course ON course.course_id = section.course_id
	JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	JOIN semester ON time_slot.semester_id = semester.semester_id
	JOIN period ON time_slot.period_id = period.period_id
	JOIN day ON time_slot.day_id = day.day_id
	JOIN location ON section.location_id = location.location_id
	JOIN building ON location.building_id = building.building_id
	JOIN room ON location.room_id = room.room_id
	WHERE section.faculty_id = ? AND semester.year = 2018 AND semester.season = 'Spring'`, user.UserID).Scan(&facultySchedule)
	fmt.Println(facultySchedule)

	err := global.Tpl.ExecuteTemplate(w, "FacultyGiveGrades", facultySchedule)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func giveStudentGrades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sectionID := vars["sectionID"]
	fmt.Println("Inside give student grades")
	fmt.Println("section id is", sectionID)
	sectionIDint,_ := strconv.Atoi(sectionID)
	type GradingForm struct{
		FirstName string
		LastName string
		Status string
	}

	gradeForm := []GradingForm{}

	db.Raw(`
	SELECT student_history.student_id,first_name,last_name,student_history.status
	FROM student_history
	JOIN enrollment ON student_history.enrollment_id = enrollment.enrollment_id
	JOIN section ON enrollment.section_id = section.section_id
	JOIN student ON enrollment.student_id = student.student_id
	JOIN main_user ON main_user.user_id = student.student_id
	WHERE section.section_id = ? AND student_history.status = 'In progress'`,sectionIDint).Scan(&gradeForm)

	fmt.Println(gradeForm)

	err := global.Tpl.ExecuteTemplate(w, "FacultyGiveGradesForm", gradeForm)
	if err != nil {
		fmt.Println(err.Error())
	}

	//on form submit, loop through and do an update for each row in the previous query on grade


}