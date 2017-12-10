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

func ViewRegisteredCourses(w http.ResponseWriter, r *http.Request){

	type StudentSchedule struct {
		CourseName string
		CourseCredits string
		RoomNumber string
		BuildingName string
		Time string
		MeetingDay string
		FirstName string
		LastName string
		SectionID uint
		StudentID uint
	}


	_, user := CheckLoginStatus(w,r)
	ss := []StudentSchedule{}
	db.Raw(`SELECT student_history.student_id,course_name,section.section_id,course_credits,building_name,room.room_number,meeting_day, first_name, last_name, time,student_history.status
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
	WHERE enrollment.student_id = ? AND student_history.status = ?`, user.UserID,"Registered").Scan(&ss)
	fmt.Println(ss)

	err := global.Tpl.ExecuteTemplate(w, "ViewStudentRegistered", ss)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DropRegisteredCourse(w http.ResponseWriter, r *http.Request){
	sec := r.FormValue("section")
	stud := r.FormValue("student")

	//TODO first check if its an open registration period otherwise its too late to drop

	fmt.Println("sec id", sec)
	fmt.Println("stud id", stud)

	enrollment := model.Enrollment{}
	db.Table("enrollment").Select("*").Where("student_id = ? AND section_id = ?",stud,sec).Scan(&enrollment)

	hist := model.StudentHistory{}
	db.Table("student_history").Select("*").Where("enrollment_id = ? AND student_id = ?", enrollment.EnrollmentID,stud).Scan(&hist)

	db.Delete(&hist) //delete the hist
	db.Delete(&enrollment) //delete the enrollment

	global.Tpl.ExecuteTemplate(w, "studentSuccess", "The course has been dropped")
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

	println("Inside searchMasterSchedule")

	queryVals := r.URL.Query()

	departmentQuery,_ := queryVals["department"]
	courseNameQuery,_ := queryVals["course-name"]
	courseNumQuery := queryVals["course-number"]
	professorQuery := queryVals["instructor"]
	day := queryVals["day"]

	depID := departmentQuery[0]
	courseName := courseNameQuery[0]
	courseNum  := courseNumQuery[0]
	professor := professorQuery[0]
	dayID := day[0]

	whereMap := make(map[string]interface{})

	whereStuff := "WHERE "

	numQueries := 0

	if depID != "" {
		println("Department query present: " + depID)
		//depID, _ := strconv.ParseUint(departmentQuery[0], 10, 64)
		whereMap["department_id"] = depID
		whereStuff += "department_id = " + depID
		numQueries++
	}
	if courseName != "" {
		whereMap["course_name"] = courseName
		if numQueries == 0 {
			whereStuff += " course_name = '" + courseName + "'"

		}else {
			whereStuff += " AND course_name = '" + courseName + "'"

		}
		numQueries++
	}

	if dayID != ""{
		if numQueries == 0 {
			whereStuff += " day.day_id = " + dayID
		} else {
			whereStuff += " AND day.day_id = " + dayID
		}
		numQueries++
	}

	if courseNum != "" {

		if numQueries == 0 {
			whereStuff += " course_num = " + courseNum
		}else {
			whereStuff += " AND course_num = " + courseNum
		}
		numQueries++

	}
	if professor != "" {
		prof := strings.Split(professor, " ")
		if numQueries == 0 {
			whereStuff += " first_name = '" + prof[0] + "'"
			whereStuff += " AND last_name = '" + prof[1] + "'"

		}else {
			whereStuff += " AND first_name = '" + prof[0] + "'"
			whereStuff += " AND last_name = '" + prof[1] + "'"
		}
		numQueries++

	}

	//registering for next semester
	if numQueries == 0 {
		whereStuff += " semester.year = 2018 AND semester.season = 'Spring'"
	}else {
		whereStuff += " AND semester.year = 2018 AND semester.season = 'Spring'"
	}

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
		Prerequisites []model.Course
	}

	//coursesFound := []model.Course{}
	//db.Where(model.Course{CourseName: courseName}).Or(model.Course{DepartmentID: uint(depID)}).Or(model.Course{CourseName: courseName}).Find(&coursesFound)

	queryRes := []CourseData{}


	//rows, err := db.Joins("JOIN course ON course.course_id = section.course_id").Where(whereMap).Rows()
	sql := `SELECT course.course_name, course.course_credits, course.course_description, course.department_id, section.section_id, section.course_section_number,
	section.course_id, section.faculty_id, section.time_slot_id, section.location_id, section.course_section_number,
	main_user.first_name, main_user.last_name,
	day.meeting_day, day.day_id,
	building.building_name,time,
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

	if strings.Contains(whereStuff,";&"){
		fmt.Println("escape")
	}
	fmt.Println("QUery to be run is", sql)
	db.Raw(sql).Scan(&queryRes)

	//if err == nil{
	//	//rows.Scan(&queryRes)
	//
	//} else {
	//	println(err.Error())
	//}
	/*
	for _, val := range queryRes{
		println(val.CourseName)
	}
	*/
	fmt.Println(queryRes)
	allDepartments := []model.Department{}

	db.Find(&allDepartments)

	coursesFoundPrereqs := make(map[string][]model.Course)

	// get prereqs for every unique course in queryRes
	for qIndex, val := range queryRes{
		if _, present := coursesFoundPrereqs[val.CourseName]; !present {
			course := model.Course{}
			db.First(&course, val.CourseID)
			prereqs := course.FindCoursePrerequisites(db)
			coursesFoundPrereqs[val.CourseName] = prereqs
			queryRes[qIndex].Prerequisites = prereqs
			println("Adding prereqs for: " + course.CourseName)
			for _,c := range val.Prerequisites{
				println("-prereq::" + c.CourseName)
			}
		} else {
			queryRes[qIndex].Prerequisites = coursesFoundPrereqs[val.CourseName]
		}
	}
	//
	//
	//for _,c := range queryRes{
	//	println("Sent to template: course-" + c.CourseName + ", Prereqs:")
	//	for _,p := range c.Prerequisites{
	//		println(p.CourseName)
	//	}
	//}

	//chosenDep := model.Department{}
	//db.First(&chosenDep, depID)
	searchParams := map[string]string{
		//"Department": chosenDep.DepartmentName,
		"Professor": professor,
		"CourseName": courseName,
		"CourseNum": courseNum,
	}

	data :=  map[string]interface{}{
		"Results": queryRes,
		"Departments": allDepartments,
		"Params": searchParams,
	}

	global.Tpl.ExecuteTemplate(w, "studentRegisterCourse", data)

}

func RegisterForSection(w http.ResponseWriter, r *http.Request){
	section := r.FormValue("section")
	_,user := CheckLoginStatus(w,r)
	fmt.Println("Section is" ,section)

	//TODO check if its a registration period

	holds := 0

	db.Table("student_holds").Select("*").Where("student_id = ?",user.UserID).Count(&holds)
	//check holds
	if holds > 0 {
		fmt.Println("Error can't regester because you have holds")
		global.Tpl.ExecuteTemplate(w, "studentSuccess", "Error can't regester because you have holds" )
		return
	}else {
		fmt.Println("No holds, continue on") //holds check works , add view
	}

	student := model.Student{}
	maxCredits := 0

	//db.Where(model.Student{StudentID:user.UserID}).Scan(&student)
	db.First(&student,user.UserID)
	fmt.Println("Found student", student)

	if student.StudentType == 1{ //if full time student limit is 16
		maxCredits = 16
	}else { //part time limit is 8
		maxCredits = 8
	}

	type CourseRegistering struct {
		SectionID uint
		CourseID uint
		CourseCredits int
		MeetingDay string
		Time string
	}
	//course attempting to register
	courseRegistering := CourseRegistering{}
	totalCredits := 0

	db.Raw(`
		SELECT section.section_id, course.course_id, course_credits, meeting_day,time
		FROM section
		JOIN course on course.course_id = section.course_id
		JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
		JOIN semester ON time_slot.semester_id = semester.semester_id
		JOIN period ON time_slot.period_id = period.period_id
		JOIN day ON time_slot.day_id = day.day_id
		WHERE section.section_id = ?

	`,section).Scan(&courseRegistering)

	type RegistrationCheck struct {
		SectionID uint
		CourseID uint
		CourseCredits int
		MeetingDay string
		Time string
		Grade string
		Status string
	}

	registrationCheck := []RegistrationCheck{}
	//already enrolled courses
	db.Raw(`
			SELECT section.section_id,section.course_id,course_credits, meeting_day,time, grade, status
			FROM section
			JOIN course ON course.course_id = section.course_id
			JOIN  enrollment ON enrollment.section_id = section.section_id
			JOIN student_history ON enrollment.enrollment_id = student_history.enrollment_id
			JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
			JOIN semester ON time_slot.semester_id = semester.semester_id
			JOIN period ON time_slot.period_id = period.period_id
			JOIN day ON time_slot.day_id = day.day_id
			WHERE enrollment.student_id = ? AND status = ?
	`,user.UserID,"Registered").Scan(&registrationCheck)

	fmt.Println("Registered courses are" , registrationCheck)

	//check course attempting to register vs courses already registered
	for i := 0; i < len(registrationCheck); i++{
		totalCredits += registrationCheck[i].CourseCredits
		if totalCredits >= maxCredits { //check max credits
			fmt.Println("Error you are over the credit limit for your student type")
			global.Tpl.ExecuteTemplate(w, "studentSuccess", "Error you are over the credit limit for your student type" )
			return
		}

		//check if student is already registered for that course this semester
		if courseRegistering.CourseID == registrationCheck[i].CourseID {
			fmt.Println("Error, you are already registered for that same course")
			global.Tpl.ExecuteTemplate(w, "studentSuccess" , "Error, you are already registered for that same courseß")
			return
		}

		//check if time and day are equal
		if strings.Compare(courseRegistering.Time,registrationCheck[i].Time) == 0 &&
		strings.Compare(courseRegistering.MeetingDay,registrationCheck[i].MeetingDay) == 0 {
			fmt.Println("Error you are already registrered for that time slot at that day")
			global.Tpl.ExecuteTemplate(w, "studentSuccess", "Error you are already registrered for that time slot at that day " )
			return
		}


	}

	course := model.Course{}

	db.Where(model.Course{CourseID:courseRegistering.CourseID}).Find(&course)

	//* prereq works
	prereqs := course.FindCoursePrerequisites(db)
	fmt.Println("Prereqs for this course are ", prereqs)
	//if the course has prereqs check to make sure the student has all the prereqs
	if len(prereqs) > 0 {
		numPreReqs := len(prereqs)
		coursesTaken := []model.Course{}
		//var hasTakenList [numPreReqs]bool
		hasTakenList := make([]bool, numPreReqs)
		//set courses taken to a boolean
		db.Raw(`
			SELECT course.course_id
			FROM section
			JOIN course ON course.course_id = section.course_id
			JOIN  enrollment ON enrollment.section_id = section.section_id
			JOIN student_history ON enrollment.enrollment_id = student_history.enrollment_id
			WHERE enrollment.student_id = ? AND status != ?;
		`,user.UserID,"Registered").Scan(&coursesTaken)


		fmt.Println("Before looping")
		fmt.Println("Courses Taken", coursesTaken)
		fmt.Println("Pre Reqs needed", prereqs)
		fmt.Println("Boolean array", hasTakenList)
		//first check for each of their preques, they have taken the course
		for i := 0; i < len(prereqs); i++{
			for j := 0; j < len(coursesTaken); j++{
				if prereqs[i].CourseID == coursesTaken[j].CourseID {
					hasTakenList[i] = true
					fmt.Println("Change is has taken array", hasTakenList)
				}
			}

		}

		fmt.Println("After looping through courses taken and checking with prereqs")
		fmt.Println("Boolean array", hasTakenList)


		for i := 0 ; i < len(hasTakenList); i++{
			if hasTakenList[i] == false {
				fmt.Println("Error, missing pre-req") //Todo test this more
				global.Tpl.ExecuteTemplate(w, "studentSuccess", "Error, you are missing pre-reqs" )
				return
			}
		}
	}
	fmt.Println("Made it to bottom of function, attempting to insert enrollment & history")
	//IF we make it here everything is checked for and we can insert the enrollment
	enrollment := model.Enrollment{StudentID:user.UserID, SectionID:courseRegistering.SectionID}

	db.Create(&enrollment) //enroll
	fmt.Println("created enrollment, ", enrollment)


	history := model.StudentHistory{StudentID:user.UserID, EnrollmentID:enrollment.EnrollmentID,Status:"Registered", Grade:"-"}
	db.Create(&history)
	fmt.Println("Created history,", history)

	global.Tpl.ExecuteTemplate(w, "studentSuccess", "Registration successful." )


}

