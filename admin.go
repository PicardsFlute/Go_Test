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
	"github.com/gorilla/mux"
	"encoding/json"
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




type NoUser struct {
	ErrorMessage string
}

func viewStudentTranscriptPage(w http.ResponseWriter, r *http.Request){
	isLogged, user := CheckLoginStatus(w,r)

	if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "viewStudentTranscriptAdmin", user)
	}else {
		http.Redirect(w,r,"/",http.StatusForbidden)
	}
}



func viewStudentTranscript(w http.ResponseWriter, r *http.Request){

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
		global.Tpl.ExecuteTemplate(w,"viewStudentTranscriptAdmin", m)
		return
	}
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
	fmt.Println(st)
	m := map[string]interface{}{
		"User":user,
		"Transcript":st,
	}

	errTemp := global.Tpl.ExecuteTemplate(w, "adminViewStudentTranscriptDetails", m)
	if errTemp != nil {
		fmt.Println(errTemp.Error())
	}

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
	WHERE enrollment.student_id = ? AND student_history.status = 'In progress'`,user.UserID).Scan(&ss)
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
	studentHold := model.StudentHolds{}
	//db.Where("student_id = ? AND hold_id = ?", userInt,holdInt).First(&studentHold)
	//db.Delete(&hold)
	db.Raw("SELECT * FROM student_holds WHERE student_id = ? AND hold_id = ?", userInt,holdInt).Scan(&studentHold)
	fmt.Println("Hold found", studentHold)
	db.Delete(&studentHold)
	//fmt.Println("Hold deleted sucessfully")


}



func AdminAddCoursePage(w http.ResponseWriter, r *http.Request){
	//isLogged, user := CheckLoginStatus(w,r)
	//if isLogged && user.UserType == 3 {
	departments := []model.Department{}

		db.Table("department").Select("*").Scan(&departments)
		err :=global.Tpl.ExecuteTemplate(w, "addCourseAdmin", departments)
		if err != nil{
			fmt.Println(err.Error())
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
	errors := db.Create(&course)
	if errors.Error != nil {
		fmt.Println(errors.Error)
		return
	}

	courses := []model.Course{}
	db.Table("course").Select("course_name, course_id").Scan(&courses)


	m := map[string]interface{}{
		"Courses": courses,
		"CurrentCourse":course,
	}

	global.Tpl.ExecuteTemplate(w, "coursePreReq", m)

}

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func AddCoursePreRequisit(w http.ResponseWriter, r *http.Request){
	fmt.Println("inside admin add course pre-req")

	jsonVal := r.FormValue("prereqs")
	courseID := r.FormValue("course")

	courseIDint,_ := strconv.Atoi(courseID)

	fmt.Println("Course prereqs are",jsonVal)
	fmt.Println("Course id is",courseID)

	courses := []Course{}

	bytes := []byte(jsonVal)
	err := json.Unmarshal(bytes,&courses)
	if err != nil{
		fmt.Print(err.Error())
	}

	fmt.Println("JSON data is",courses)


	for _, course := range courses {
		courseRequirementIDint, _ := strconv.Atoi(course.Id)
		prereq := model.Prerequisite{CourseRequiredBy:uint(courseIDint), CourseRequirement:uint(courseRequirementIDint)}
		db.Create(&prereq)
		fmt.Println("Course added", prereq)
	}

	//TODO: JUst add success page, must change view with js after ajax call
	//TODO other options involve using a form instead of ajax

	global.Tpl.ExecuteTemplate(w, "adminSuccess", nil)

}

func AdminSearchCoursePage(w http.ResponseWriter, r *http.Request){
	departments := []model.Department{}
	db.Table("department").Select("*").Scan(&departments)

	//isLogged, user := CheckLoginStatus(w,r)
	//if isLogged && user.UserType == 3 {
		global.Tpl.ExecuteTemplate(w, "searchCourseAdmin", departments)
	//}
}
/*
type AdminViewSection struct {
	Name string
	Description string

}
*/


func AdminSearchCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Inside admin search course")
	vars := mux.Vars(r)
	id := vars["course"]
	idInt, _ := strconv.Atoi(id)
	fmt.Println("the id is", idInt)
	course := model.Course{}
	//db.Table("course").Select("*").Where("course_id")
	db.Where(model.Course{CourseID:uint(idInt)}).Find(&course)
	prereqs := course.FindCoursePrerequisites(db)
	fmt.Println("Course is", course)
	fmt.Println("Pre-Reqs are", prereqs)


	courseDetail := model.Course{}
	db.Raw(`SELECT course_name,course_credits,course_description
	FROM course
	WHERE course.course_id =?`,id).Scan(&courseDetail)
	fmt.Println(courseDetail)

	//Todo:: must render view on client side if you send an ajax request
	//TODO: just show the course with all the pre-requisits

	m := map[string]interface{}{
		"Course": courseDetail,
		"PreReqs": prereqs,
	}

	err := global.Tpl.ExecuteTemplate(w, "adminViewCourseSection", m)
	if err != nil {
		fmt.Println(err.Error())
	}
	//return
}





func AdminAddSectionPage(w http.ResponseWriter, r *http.Request){
	type CourseOptions struct {
		CourseName string
		CourseID uint
	}
	courses := []CourseOptions{}
	fac := []model.MainUser{}
	periods := []model.Period{}
	buildings := []model.Building{}
	departments := []model.Department{}
	//timeSlot := []model.TimeSlot{}

	//db.Raw("SELECT user_id,first_name,last_name FROM main_user WHERE user_type= ?", 2).Scan(&fac)
	//TODO: see if we can do all this with 1 query and place into multiple structs
	//TODO: or parse through a custom struct
	db.Table("course").Select("course_name, course_id").Scan(&courses)

	db.Table("main_user").Select("*").Where("user_type = ?",2).Scan(&fac)

	db.Table("period").Select("*").Scan(&periods)

	db.Table("building").Select("*").Scan(&buildings)

	db.Table("department").Select("*").Scan(&departments)


	fmt.Println("Courses are", courses, "Faculty are", fac)
	fmt.Println("Periods are", periods)

	m := map[string]interface{}{
		"Courses": courses,
		"Faculty": fac,
		"Period": periods,
		"Building":buildings,
		"Department":departments,
	}

	errTpl := global.Tpl.ExecuteTemplate(w, "addSectionAdmin", m)
	if errTpl != nil {
		fmt.Println(errTpl.Error())
	}
}

func GetRoomsForBuilding(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	buildingId := vars["id"]
	rooms := []model.Room{}

	//if they select all then change the query

	buildingInt, _ := strconv.Atoi(buildingId)
	db.Raw("select room_id,room_type,room_number from location natural join " +
			"room natural join building where building_id = ?",buildingInt).Scan(&rooms)


	//encode the rooms to a slice of bytes in json form
	data , err := json.Marshal(rooms)
	if err != nil {
		fmt.Println(err.Error())
	}

	//write those bytes to the response
	w.Write(data)
	fmt.Println(data)
}


func GetDepartmentsForSections(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	departmentID := vars["id"]
	courses := []model.Course{}
	fmt.Println(departmentID)
	if departmentID == "*"{
		db.Raw("SELECT course_name,course_id FROM course NATURAL JOIN department").Scan(&courses)

	}else {
		departmentInt, _ := strconv.Atoi(departmentID)
		db.Raw("SELECT course_name,course_id FROM course NATURAL JOIN department WHERE department_id = ?", departmentInt).Scan(&courses)

	}


	data, err := json.Marshal(courses)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(data)
	fmt.Println(string(data))

}

/*
type SectionInfo struct {
	Year int
	Season string
	Day string
}
*/



func AdminAddSection(w http.ResponseWriter, r *http.Request){

	sectionNum := r.FormValue("section-number")
	courseSubject := r.FormValue("course-subject")
	courseName := r.FormValue("course")
	faculty := r.FormValue("faculty-name")
	time := r.FormValue("time")
	buildingNum := r.FormValue("building")
	roomNum := r.FormValue("room")
	semester := r.FormValue("semester")
	day := r.FormValue("day")

	fmt.Println("Section Num:",sectionNum)
	fmt.Println("Course Name:",courseName)
	fmt.Println("Course Subject:",courseSubject)
	fmt.Println("Faculty",faculty)
	fmt.Println("Time",time)
	fmt.Println("Building",buildingNum)
	fmt.Println("Room Num:",roomNum)
	fmt.Println("Semester:",semester)
	fmt.Println("Day",day)


	semesterInt, _ := strconv.Atoi(semester)
	dayInt , _ := strconv.Atoi(day)
	timeInt , _ := strconv.Atoi(time)
	timeSlot := model.TimeSlot{PeriodID:uint(timeInt),SemesterID:uint(semesterInt),DayID:uint(dayInt)}
	buildingInt , _ := strconv.Atoi(buildingNum)
	roomInt, _ := strconv.Atoi(roomNum)
	facultyInt, _ := strconv.Atoi(faculty)

	//TODO: Rework this, shouldn't be creating, should be searching
	db.Create(&timeSlot)

	location := model.Location{}
	db.Where(model.Location{BuildingID:uint(buildingInt),RoomID:uint(roomInt)}).First(&location)

	timeSlotID := timeSlot.TimeSlotID

	sectionInt, _ := strconv.Atoi(sectionNum)
	courseInt, _ := strconv.Atoi(courseName)
	facultyID, _ := strconv.Atoi(faculty)

	newCourseSection := model.Section{CourseSectionNumber:sectionInt,CourseID:uint(courseInt), FacultyID:uint(facultyID),
	TimeSlotID:timeSlotID,LocationID:location.LocationID}

	//TODO: Complete, now test

	type RoomCheck struct{
		Section_id int
		Location_id int
		Building_id int
		Building_name string
		Room_id int
		Room_number string
	}
	rc := []RoomCheck{}
	db.Raw(`
		 SELECT section.section_id, location.location_id, building.building_id,room.room_id,room_number,building_name,period.period_id,time,day.day_id,meeting_day
		 FROM section
		 JOIN location ON section.location_id = location.location_id
		 JOIN building ON building.building_id = location.building_id
		 JOIN room ON room.room_id = location.room_id
		 JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
		 JOIN day ON time_slot.day_id = day.day_id
		 JOIN period ON period.period_id = time_slot.period_id;
	 	 WHERE location.room_id = ? AND building.building_id = ? AND period.period_id = ?
	 	 AND day.day_id = ?`,location.RoomID, location.BuildingID,timeInt,dayInt).Scan(&rc)


	if len(rc) > 0 {
		fmt.Println("Cant add room, because is already ocupied at this time")
		return
	}


	type ProfessorCheck struct {
		UserID int
		UserEmail string
		FirstName string
		section_id int
		Period_id int
		Time string
		Day_id int
		Meeting_day string
	}

	cc := []ProfessorCheck{}
	//TODO: Complete, now test
	db.Raw(`
		SELECT user_id,user_email, first_name, section_id,period.period_id,time,day.day_id,meeting_day FROM main_user
		 JOIN faculty ON main_user.user_id = faculty.faculty_id
		 JOIN section ON faculty.faculty_id = section.section_id
		 JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
		 JOIN day ON time_slot.day_id = day.day_id
		 JOIN period ON period.period_id = time_slot.period_id
		 WHERE period.period_id = ? AND day.day_id = ? AND user_id = ?`, time,day,facultyInt).Scan(&cc)

	fmt.Println(cc)

	if len(cc) > 0{
		fmt.Println("Cant add section,teacher is already teachinga  course at this time slot exit function")
		return
	}
	db.Create(&newCourseSection)
	global.Tpl.ExecuteTemplate(w, "admin", nil)
}

func openGradingPeriod(w http.ResponseWriter, r *http.Request){


}
