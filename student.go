package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
	"strings"
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
	db.Raw(`SELECT student_history.student_id,course_name,course_credits,building_name,room.room_number,meeting_day, first_name, last_name, time,student_history.status
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


func ViewAdvisor(w http.ResponseWriter, r *http.Request){

	_ ,user := CheckLoginStatus(w,r)

	type AdvisingData struct {
		FacultyID uint
		FirstName string
		LastName string
		RoomNumber string
		DepartmentName string
		DepartmentBuilding string
		DepartmentPhoneNumber string
	}

	ad := AdvisingData{}

	db.Raw(`
		SELECT faculty.faculty_id, first_name, last_name, room_number,department_building,department_name,department_phone_number
		FROM advises
		JOIN faculty ON advises.faculty_id = faculty.faculty_id
		JOIN main_user ON faculty.faculty_id = main_user.user_id
		JOIN department ON department.department_id = faculty.department_id
		WHERE advises.student_id = ?
		`,user.UserID).Scan(&ad)


	global.Tpl.ExecuteTemplate(w, "ViewAdvisor", ad)

}

func ViewTranscript(w http.ResponseWriter, r *http.Request){

	_ ,user := CheckLoginStatus(w,r)


	type Transcript struct {
		StudentID uint
		Grade string
		Status string
		Year int
		Season string
		CourseName string
		CourseCredits int
	}

	st := []Transcript{}
	db.Raw(`
 	 SELECT enrollment.student_id,grade,status,year,season,course_name,course_credits
	 FROM student_history
	 JOIN enrollment ON student_history.enrollment_id = enrollment.enrollment_id
	 JOIN section ON enrollment.section_id = section.section_id
	 JOIN course on course.course_id = section.course_id
	 JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	 JOIN semester ON time_slot.semester_id = semester.semester_id
	 WHERE enrollment.student_id = ?`,user.UserID).Scan(&st)

	m := map[string]interface{}{
		"User":user,
		"Transcript":st,
	}


	global.Tpl.ExecuteTemplate(w, "ViewTranscript", m)

}

func AddCoursePage(w http.ResponseWriter, r *http.Request){
	_ ,user := CheckLoginStatus(w,r)
	allDepartments := []model.Department{}
	db.Find(&allDepartments)
	m :=  map[string]interface{}{
		"User": user,
		"Departments": allDepartments,
	}

	errTemp := global.Tpl.ExecuteTemplate(w, "studentRegisterCourse", m)
	if errTemp != nil {
		fmt.Println(errTemp.Error())
	}

}

func StudentSearchCourseResults(w http.ResponseWriter, r *http.Request){

	println("Inside searchMasterSchedule")

	queryVals := r.URL.Query()

	departmentQuery,_ := queryVals["department"]
	courseNameQuery,_ := queryVals["course-name"]
	courseNumQuery := queryVals["course-number"]
	professorQuery := queryVals["instructor"]

	depID := departmentQuery[0]
	courseName := courseNameQuery[0]
	courseNum  := courseNumQuery[0]
	professor := professorQuery[0]

	whereMap := make(map[string]interface{})

	whereStuff := "WHERE "

	if depID != "" {
	println("Department query present: " + depID)
	//depID, _ := strconv.ParseUint(departmentQuery[0], 10, 64)
	whereMap["department_id"] = depID
	whereStuff += "department_id = " + depID
	}
	if courseName != "" {
	whereMap["course_name"] = courseName
	whereStuff += " AND course_name = '" + courseName + "'"
	}
	if courseNum != "" {
	whereStuff += " AND course_num = " + courseNum
	}
	if professor != "" {
	prof := strings.Split(professor, " ")
	whereStuff += " AND first_name = '" + prof[0] + "'"
	whereStuff += " AND last_name = '" + prof[1] + "'"
	}

	//registering for next semester
	whereStuff += " AND semester.year = 2018 AND semester.season = 'Spring'"

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
	Time string
	}

	//coursesFound := []model.Course{}
	//db.Where(model.Course{CourseName: courseName}).Or(model.Course{DepartmentID: uint(depID)}).Or(model.Course{CourseName: courseName}).Find(&coursesFound)

	queryRes := []CourseData{}

	//TODO: Phil add prerequisits to query and display the rest of the data in MS search

	//rows, err := db.Joins("JOIN course ON course.course_id = section.course_id").Where(whereMap).Rows()
	sql := `SELECT course.course_name, course.course_credits, course.course_description, course.department_id, section.section_id, section.course_section_number,
	section.course_id, section.faculty_id, section.time_slot_id, section.location_id, section.course_section_number,
	main_user.first_name, main_user.last_name,
	day.meeting_day, day.day_id,

	building.building_name,
	room.room_number, room.room_type
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

	fmt.Println("QUery to be run is", sql)
	db.Raw(sql).Scan(&queryRes)


	fmt.Println(queryRes)
	allDepartments := []model.Department{}

	db.Find(&allDepartments)

	data :=  map[string]interface{}{
	"Results": queryRes,
	"Departments": allDepartments,
	}

	global.Tpl.ExecuteTemplate(w, "studentRegisterCourse", data)

}

func RegisterForSection(w http.ResponseWriter, r *http.Request){

}

