package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
	"github.com/gorilla/mux"
	"strconv"
	"strings"
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
	WHERE section.faculty_id = ? AND semester.year = ? AND semester.season = ?`, user.UserID,2017,"Fall").Scan(&facultySchedule)
	//TODO: these are hardcoded for current semester, fix these at some point


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
	WHERE section.faculty_id = ? AND semester.year = ? AND semester.season = ?`, user.UserID,2017,"Fall").Scan(&facultySchedule)
	fmt.Println(facultySchedule)
	//TODO: these are hardcoded for current semester, fix these at some point
	err := global.Tpl.ExecuteTemplate(w, "FacultyGiveGrades", facultySchedule)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func giveStudentGradesForm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sectionID := vars["sectionID"]
	sectionIDint,_ := strconv.Atoi(sectionID)

	type GradingPeriod struct {
		SectionID uint
		SemesterID uint
		SemesterStatus string
	}

	gp := GradingPeriod{}

	db.Raw(`
		SELECT section.section_id, semester.semester_id, semester_status
		FROM section
		JOIN time_slot ON section.time_slot_id = time_slot.time_slot_id
		JOIN semester ON time_slot.semester_id = semester.semester_id
		WHERE section.section_id = ?
	`, sectionIDint).Scan(&gp)

	if strings.Compare(gp.SemesterStatus,"Grading") != 0 {
		fmt.Println("Error it is not currently a grading period")
		global.Tpl.ExecuteTemplate(w, "facultySuccess", "Error it is not currently a grading period")
		return
	}

	type CourseInfo struct{
		CourseName string
		SectionID uint
	}

	courseDetail := CourseInfo{}

	db.Raw(`
		SELECT course_name,section_id
		FROM section
		JOIN course ON course.course_id = section.course_id
		WHERE section.section_id = ?`,sectionIDint).Scan(&courseDetail)

	fmt.Println("Course Detail is", courseDetail)

	type GradingForm struct{
		StudentID uint
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

	m := map[string]interface{}{
		"Course":courseDetail,
		"Roster":gradeForm,
	}

	err := global.Tpl.ExecuteTemplate(w, "FacultyGiveGradesForm", m)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func submitGrades(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	sectionID := vars["sectionID"]
	sectionIDint,_ := strconv.Atoi(sectionID)
	//grades := r.FormValue("grade")
	r.ParseForm()
	id := r.Form["studentId"]
	grades := r.Form["grade"]


	type Err struct {
		error string
	}

	e := Err{}

	for k,_ := range grades{
		db.Raw(`
			UPDATE student_history
			SET grade = ?,
			status = 'Complete'
			FROM enrollment
			WHERE enrollment.enrollment_id = student_history.enrollment_id
			AND enrollment.section_id = ?
			AND student_history.student_id = ?
		`,grades[k],sectionIDint,id[k]).Scan(&e)
		fmt.Println("Updating grade with info sec, grade, stud_id", sectionIDint,grades[k], id[k])
	}

	fmt.Println(e)

	global.Tpl.ExecuteTemplate(w, "facultySuccess", "Grades sucessfully submitted")


}