package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"html/template"
	"Starfleet/model"
	"Starfleet/global"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request){
	global.Tpl.ExecuteTemplate(w, "index", nil)

}

func loginPage(w http.ResponseWriter, r *http.Request){
	// use session information to determine the user, and if they are already logged in
	logged, u := CheckLoginStatus(w, r)
	// a user already logged in will be sent to the page of their respective role
	if logged{
		checkUserType(u,w,r)
	}else {
		global.Tpl.ExecuteTemplate(w, "login", nil)
	}
}

func redirectPost(w http.ResponseWriter, r *http.Request){
	req, err := http.NewRequest("DELETE", "/admin/holds/{id}", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	respBody , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP RESPONSE FROM DELETE IS", string(respBody))

}



func loginUser(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		//is a POST
		formEmail := r.FormValue("email")
		formPassword :=	r.FormValue("password")
		// Try to find user in DB
		user := model.MainUser{}
		db.Where(&model.MainUser{UserEmail: formEmail}).First(&user)

		if user.UserEmail != "" {
			dbPassword := user.UserPassword

			if user.CheckPasswordMatch(formPassword) {
				fmt.Println("User found in DB with email:", formEmail, " and password: ", dbPassword)
				sess.Set("username", r.Form["username"])
				sess.Set("UserID", user.UserID)
				//http.Redirect(w,r,"/user/" + strconv.Itoa(int(user.UserID)), http.StatusFound)

				//Tpl.ExecuteTemplate(w,"user",user)
				checkUserType(user, w, r)
			} else {
				global.Tpl.ExecuteTemplate(w,"login","Error, username or password does not match.")
			}

		} else {
			fmt.Println()
			global.Tpl.ExecuteTemplate(w,"login","User not found")
		}
	}
}


func checkUserType(user model.MainUser, w http.ResponseWriter, r *http.Request){

	switch user.UserType {

	case 1:
		fmt.Println("You're a student")
		fmt.Println("User data is", user.FirstName)
		http.Redirect(w,r,"/student", http.StatusFound)

		// The data is lost after redirect because it's a new request,
		// now I need to get the student data and render the template, which is a different request
		//since http is stateless, you l;ose the data structure after the first request.
	case 2:
		fmt.Println("Youre a faculty")
		http.Redirect(w,r,"/faculty", http.StatusFound)


	case 3:
		fmt.Println("Youre an admin")
		http.Redirect(w,r,"/admin", http.StatusFound)
		//Tpl.ExecuteTemplate(w,"admin", "administrative user!")

	case 4:
		fmt.Println("Youre a researcher")
		http.Redirect(w,r,"/researcher", http.StatusFound)
		//Tpl.ExecuteTemplate(w,"admin", "administrative user!")

	default:
		fmt.Println("Not sure your type")
		http.Redirect(w,r,"/", http.StatusFound)
		global.Tpl.ExecuteTemplate(w,"index",nil)
		//return user,user.UserType
	}

}




func CheckLoginStatus(w http.ResponseWriter, r *http.Request) (bool,model.MainUser){
	sess := globalSessions.SessionStart(w,r)
	sess_uid := sess.Get("UserID")
	u := model.MainUser{}
	if sess_uid == nil {
		//http.Redirect(w,r, "/", http.StatusForbidden)
		//Tpl.ExecuteTemplate(w,"index", "You can't access this page")
		return false, u
	} else {
		uID := sess_uid
		db.First(&u, uID)
		fmt.Println("Logged in User, ", uID)
		//Tpl.ExecuteTemplate(w, "user", nil)
		return true, u
	}
}

/*
In this snippet we're placing our handler logic in an anonymous function
 and closing-over the message variable to form a closure.
 We're then converting this closure to a handler by using the http.HandlerFunc adapter and returning it.
 */
func checkSessionWrapper(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware")
		isLogged, _ := CheckLoginStatus(w, r)
		if isLogged { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkStudent(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for student")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType ==1 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkFaculty(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for faculty")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 2 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkAdmin(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for admin")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 3 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}

func checkResearcher(handle http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewware for researcher")
		isLogged, user := CheckLoginStatus(w, r)
		if isLogged && user.UserType == 5 { //if check user is true, execute the handle that's inside
			handle.ServeHTTP(w,r)
		} else{ //otherwise deny request
			//Tpl.ExecuteTemplate(w,"index", "You can't access that page")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)           // redirects route and gives unauthorized link
			global.Tpl.ExecuteTemplate(w,"login", "You must login first.") //this renders the index template right under it
		}

	})
}





func displayStudent(w http.ResponseWriter, r *http.Request){

	_, user := CheckLoginStatus(w, r)
	if user.UserType == 1 {
		global.Tpl.ExecuteTemplate(w, "student", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func displayFaculty(w http.ResponseWriter, r *http.Request){
	_, user := CheckLoginStatus(w,r)

	if user.UserType == 2 {
		global.Tpl.ExecuteTemplate(w, "faculty", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func displayAdmin(w http.ResponseWriter, r *http.Request){
	_, user := CheckLoginStatus(w,r)

	if user.UserType == 3 {
		m := map[string]interface{}{
			"User":user,
		}
		global.Tpl.ExecuteTemplate(w, "admin", m)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}



func displayResearcher(w http.ResponseWriter, r *http.Request){

	_, user := CheckLoginStatus(w,r)

	if user.UserType == 4 {
		global.Tpl.ExecuteTemplate(w, "researcher", user)
	}else {
		http.Redirect(w,r,"/", http.StatusForbidden)
		index(w,r)
	}
}


func logout(w http.ResponseWriter, r *http.Request){
	sess := globalSessions.SessionStart(w, r)
	//sid := sess.SessionID()
	sess.Delete("UserID")
	sess.Delete("username")
	http.Redirect(w,r,"/login", http.StatusSeeOther)
}



func searchMasterScheduleForm(w http.ResponseWriter, r *http.Request){
	allDepartments := []model.Department{}
	db.Find(&allDepartments)
	m :=  map[string]interface{}{
		"Departments": allDepartments,
	}
	global.Tpl.ExecuteTemplate(w, "masterScheduleSearch", m)
}



func searchMasterSchedule(w http.ResponseWriter, r *http.Request){

	println("Inside searchMasterSchedule")

	queryVals := r.URL.Query()

	departmentQuery,_ := queryVals["department"]
	courseNameQuery,_ := queryVals["course-name"]
	professorQuery := queryVals["instructor"]
	day := queryVals["day"]

	depID := departmentQuery[0]
	courseName := courseNameQuery[0]
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

	type IsAdmin struct {
		IsAdmin bool
		User model.MainUser
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
		Capacity uint
		Time string
		Prerequisites []model.Course
		User IsAdmin
		NumEnrolled int
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
	room.room_number, room.room_type,capacity

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

	//chosenDep := model.Department{}
	//db.First(&chosenDep, depID)
	searchParams := map[string]string{
		//"Department": chosenDep.DepartmentName,
		"Professor": professor,
		"CourseName": courseName,
		//"CourseNum": courseNum,
	}

	_,user := CheckLoginStatus(w,r)



	//courseEnrollment := make(map[int]int)
	for k, _ := range queryRes {
		count := 0
		db.Table("enrollment").Select("*").Where("section_id = ?", queryRes[k].SectionID).Count(&count)
		queryRes[k].NumEnrolled = count
	}


	admin := IsAdmin{}

	if user.UserType == 3 {
		admin.IsAdmin = true
		admin.User = user
	}else{
		admin.IsAdmin = false
	}

	type IsStudent struct {
		IsStudent bool
		User model.MainUser
	}

	student := IsStudent{}

	if user.UserType == 1 {
		student.IsStudent = true
		student.User = user
	}else {
		student.IsStudent = false
	}


	data :=  map[string]interface{}{
		"Results": queryRes,
		"Departments": allDepartments,
		"Params": searchParams,
		"User":admin,
		"Student": student,
	}

	err := global.Tpl.ExecuteTemplate(w, "masterScheduleSearch", data)

	if err != nil {
		fmt.Println(err.Error())
	}

}
