package main

import (
	"github.com/mohae/struct2csv"
	"net/http"
	"Starfleet/model"
	"strconv"
	"time"
	"os"
	"log"
	"Starfleet/global"
)

func getStudentsReportByGrade(w http.ResponseWriter, r *http.Request){
	allDepartments := []model.Department{}
	db.Find(&allDepartments)
	allCourses := []model.Course{}
	db.Find(&allCourses)
	m := map[string]interface{}{
		"Departments": allDepartments,
		"Courses": allCourses,
	}
	global.Tpl.ExecuteTemplate(w, "researchStudentsByGrade", m)
}

func genReportStudentsByGrade(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	formGradeLow := r.FormValue("letter-grade-boundlow")
	formGradeHigh := r.FormValue("letter-grade-boundhigh")
	courseNum := r.FormValue("department")

	grades := []string{"F","C-","C","C+","B-","B", "B+", "A-", "A"}

	getIndex := func(vs []string, t string) int {
		for i, v := range vs {
			if v == t {
				return i
			}
		}
		return -1
	}

	indexLow := getIndex(grades, formGradeLow)
	indexHigh := getIndex(grades, formGradeHigh)
	if indexLow > indexHigh {
		temp := indexHigh
		indexHigh = indexLow
		indexLow = temp
	}
	println("gradeLow:" + formGradeLow + ", gradeHigh: " + formGradeHigh)
	println("IndexLow:" + strconv.Itoa(indexLow) + ", indexHigh: " + strconv.Itoa(indexHigh))
	gradeSlice := grades[indexLow:indexHigh+1]

	println("Seaching for grades between "+ grades[indexLow] + " and " + grades[indexHigh])
	for g := range gradeSlice {
		println(g)
	}

	contains := func(slice []string, item string) bool {
		set := make(map[string]struct{}, len(slice))
		for _, s := range slice {
			set[s] = struct{}{}
		}
		_, ok := set[item]
		return ok;
	}

	productsSelected := r.Form["filter-options"]

	sql := `SELECT student_history.grade, main_user.first_name, main_user.last_name
	FROM student_history
	JOIN enrollment ON student_history.enrollment_id = enrollment.enrollment_id
	JOIN section ON enrollment.section_id = section.section_id
	JOIN student ON student_history.student_id = student.student_id
	JOIN main_user ON main_user.user_id = student.student_id
	WHERE
	student_history.grade IN (?)
	AND student_history.status = ?
	AND section.course_id = ? `

	//if contains(productsSelected, "major"){
	//	major := r.FormValue("major")
	//	sqlDepartmentFilter := ` AND student.student_id IN
	//	(SELECT student_id FROM student_major
	//	JOIN major ON student_major.major_id = major.major_id
	//	WHERE major.department_id = `+ major + " )"
	//
	//	sql += sqlDepartmentFilter
	//}

	if contains(productsSelected, "student-type"){
		stuType := r.FormValue("full-or-part")
		println("Filtering for student type: " + stuType)
		sqlDepartmentFilter := ` AND student.student_type = `+ stuType
		sql += sqlDepartmentFilter
	}

	type StudentData struct {
		FirstName string
		LastName string
		Grade string
	}
	records := []StudentData{}
	db.Raw(sql, gradeSlice, "Complete", courseNum).Scan(&records)

	modtime := time.Now()
	filepath := "Record" + strconv.Itoa(modtime.Nanosecond()) + ".csv"
	outfile, err := os.Create("./"+filepath)

	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer outfile.Close()
	writer := struct2csv.NewWriter(outfile)

	for _, record := range records {
		if err := writer.WriteStruct(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		} else {
			println("FNAME: " + record.FirstName + ", LNAME: " + record.LastName + ", GRADE: " + record.Grade)
		}
	}

	writer.Flush()
	// tell the browser the returned content should be downloaded
	w.Header().Set("Content-Type", "text/csv")
	path := "./"+filepath
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath)

	http.ServeFile(w,r,path)
	global.Tpl.ExecuteTemplate(w, "researchStudentsByGrade", nil)
	os.Remove(path)
}