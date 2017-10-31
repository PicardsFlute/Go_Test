package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"Starfleet/model"
	"os"
	"fmt"
)

func main(){
	dbPassword := os.Getenv("PG_DATABASE_PW")
	db, err := gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password="+dbPassword)
	if err != nil {
		fmt.Println("Cannot connect to database...")
		fmt.Println("DB Error: ", err)
	}
	db.SingularTable(true)

	db.DropTable(&model.Prerequisite{})
	db.DropTable(&model.Course{})

	db.DropTable(&model.TimeSlot{})
	db.DropTable(&model.Day{})
	db.DropTable(&model.Semester{})
	db.DropTable(&model.Period{})

	db.DropTable(&model.StudentMajor{})
	db.DropTable(&model.StudentMinor{})
	db.DropTable(&model.Major{})
	db.DropTable(&model.Major{})
	db.DropTable(&model.FullTimeStudent{})
	db.DropTable(&model.PartTimeStudent{})
	db.DropTable(&model.Student{})
	db.DropTable(&model.PartTimeFaculty{})
	db.DropTable(&model.FullTimeFaculty{})
	db.DropTable(&model.Faculty{})
	db.DropTable(&model.Researcher{})
	db.DropTable(&model.Admin{})
	db.DropTable(&model.MainUser{})

	db.DropTable(&model.Department{})


	db.AutoMigrate(
		&model.MainUser{},
		&model.Student{},
		&model.PartTimeStudent{},
		&model.FullTimeStudent{},
		&model.Department{},
		&model.Faculty{},
		&model.PartTimeFaculty{},
		&model.FullTimeFaculty{},
		&model.Admin{},
		&model.Researcher{},
		&model.Major{},
		&model.Minor{},
		&model.StudentMajor{},
		&model.StudentMinor{},

		&model.Course{},
		&model.Prerequisite{},

		&model.Day{},
		&model.Semester{},
		&model.Period{},
		&model.TimeSlot{},

	)

	db.Model(&model.Student{}).AddForeignKey("student_id", "main_user(user_id)", "CASCADE", "CASCADE")

	db.Model(&model.PartTimeStudent{}).AddForeignKey("part_time_student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.FullTimeStudent{}).AddForeignKey("full_time_student_id", "student(student_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Faculty{}).AddForeignKey("faculty_id", "main_user(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Faculty{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.PartTimeFaculty{}).AddForeignKey("part_time_faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.FullTimeFaculty{}).AddForeignKey("full_time_faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Admin{}).AddForeignKey("admin_id", "main_user(user_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Researcher{}).AddForeignKey("researcher_id", "main_user(user_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Major{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Minor{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.StudentMajor{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentMajor{}).AddForeignKey("major_id", "major(major_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.StudentMinor{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentMinor{}).AddForeignKey("major_id", "major(major_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.TimeSlot{}).AddForeignKey("day_id", "day(day_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("semester_id", "semester(semester_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("period_id", "period(period_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Course{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Prerequisite{}).AddForeignKey("course_required_by", "course(course_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Prerequisite{}).AddForeignKey("course_requirement", "course(course_id)", "RESTRICT", "RESTRICT")

	user1 := model.MainUser{FirstName: "Pat", LastName:"Lagat", UserEmail:"plagat@yahoo.com", UserPassword:"pl12345", UserType:1}
	user2 := model.MainUser{FirstName: "Irish", LastName:"James", UserEmail:"jirish@yahoo.com", UserPassword:"ij12345", UserType:1}
	user3 := model.MainUser{FirstName: "Trey", LastName:"Gorkin", UserEmail:"gork@yahoo.com", UserPassword:"tg12345", UserType:1}
	user4 := model.MainUser{FirstName: "Testy", LastName:"McTest", UserEmail:"test@test.test", UserPassword:"testPW", UserType:1}
	user5 := model.MainUser{FirstName: "Faculty", LastName:"McFaculton", UserEmail:"faculty@test.test", UserPassword:"testPW", UserType:2}
	user6 := model.MainUser{FirstName: "Ima", LastName:"Faculty", UserEmail:"ifaculty@starfleet.edu", UserPassword:"testPW", UserType:2}
	user7 := model.MainUser{FirstName: "Jordi", LastName:"LaForge", UserEmail:"laforge@starfleet.edu", UserPassword:"testPW", UserType:2}
	user8 := model.MainUser{FirstName: "Admin", LastName:"McAdminton", UserEmail:"admin@test.test", UserPassword:"testPW", UserType:3}
	user9 := model.MainUser{FirstName: "George", LastName:"Admintonson", UserEmail:"adminson@starfleet.edu", UserPassword:"testPW", UserType:3}
	user10 := model.MainUser{FirstName: "Lesdo", LastName:"SomeResearch", UserEmail:"research@starfleet.edu", UserPassword:"testPW", UserType:4}

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)
	db.Create(&user5)
	db.Create(&user6)
	db.Create(&user7)
	db.Create(&user8)
	db.Create(&user9)
	db.Create(&user10)

	student2 := model.Student{StudentID: user2.UserID, StudentType:1}
	student1 := model.Student{StudentID: user1.UserID, StudentType:1}
	student3 := model.Student{StudentID: user3.UserID, StudentType:2}
	student4 := model.Student{StudentID: user4.UserID, StudentType:2}
	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&student4)

	// example of finding related models
	u := model.MainUser{}
	db.Model(&student2).Association("MainUser").Find(&u)
	fmt.Println("For the student2, the user email is: ", u.UserEmail)

	fullTimeStudent1 := model.FullTimeStudent{FullTimeStudentID: student1.StudentID, NumCredits:16}
	fullTimeStudent2 := model.FullTimeStudent{FullTimeStudentID: student2.StudentID, NumCredits:18}
	db.Create(&fullTimeStudent1)
	db.Create(&fullTimeStudent2)
	partTimeStudent1 := model.PartTimeStudent{PartTimeStudentID: student3.StudentID, NumCredits:12}
	partTimeStudent2 := model.PartTimeStudent{PartTimeStudentID: student4.StudentID, NumCredits:8}
	db.Create(&partTimeStudent1)
	db.Create(&partTimeStudent2)

	// example of lookup back from part_time_student to main_user
	uLookup := model.MainUser{}
	stuLookup := model.Student{}
	db.Model(&partTimeStudent1).Association("Student").Find(&stuLookup)
	db.Model(&stuLookup).Association("MainUser").Find(&uLookup)
	fmt.Println("For the partTimeStudent1, the user email is: ", uLookup.UserEmail)

	department1 := model.Department{DepartmentName:"Math", DepartmentBuilding:"MathBuilding", DepartmentRoomNumber:"302", DepartmentChair:"Mr Math", DepartmentPhoneNumber:"111-222-3333"}
	department2 := model.Department{DepartmentName:"Computer Science", DepartmentBuilding:"CSBuilding", DepartmentRoomNumber:"100", DepartmentChair:"Mr Computer", DepartmentPhoneNumber:"123-456-7899"}
	db.Create(&department1)
	db.Create(&department2)

	faculty1 := model.Faculty{FacultyID:user5.UserID, FacultyType:1, DepartmentID:department1.DepartmentID}
	faculty2 := model.Faculty{FacultyID:user6.UserID, FacultyType:2, DepartmentID:department2.DepartmentID}
	faculty3 := model.Faculty{FacultyID:user7.UserID, FacultyType:1, DepartmentID:department1.DepartmentID}
	db.Create(&faculty1)
	db.Create(&faculty2)
	db.Create(&faculty3)

	// example of finding a many-one lookup
	dep := model.Department{}
	db.Model(&faculty1).Association("Department").Find(&dep)
	fmt.Println("For the faculty1, the department is: ", dep.DepartmentName)
	facMembers := []model.Faculty{}
	// search faculty by department (.Association and .Related doesn't seem to work)
	//db.Model(&department1).Association("Faculty").Find(&facMembers)
	db.Where(model.Faculty{DepartmentID: department1.DepartmentID}).Find(&facMembers)
	fmt.Println("For the departent1 (Math), the faculty is: ")
	for _, v := range facMembers {
		fmt.Println("FacultyID", v.FacultyID)
	}


	admin1 := model.Admin{AdminID:user8.UserID}
	admin2 := model.Admin{AdminID:user9.UserID}
	db.Create(&admin1)
	db.Create(&admin2)

	researcher1 := model.Researcher{ResearcherID:user10.UserID}
	db.Create(researcher1)


	major1 := model.Major{DepartmentID: department1.DepartmentID, MajorName:"MATH"}
	major2 := model.Major{DepartmentID: department1.DepartmentID, MajorName:"SUPERMATH"}
	major3 := model.Major{DepartmentID: department2.DepartmentID, MajorName:"CIS"}
	major4 := model.Major{DepartmentID: department2.DepartmentID, MajorName:"MIS"}
	db.Create(&major1)
	db.Create(&major2)
	db.Create(&major3)
	db.Create(&major4)

	minor1 := model.Minor{DepartmentID: department1.DepartmentID, MinorName:"APPLIED MATH"}
	minor2 := model.Minor{DepartmentID: department1.DepartmentID, MinorName:"MINOR MATH"}
	minor3 := model.Minor{DepartmentID: department2.DepartmentID, MinorName:"LITTLE COMPUTERS"}
	minor4 := model.Minor{DepartmentID: department2.DepartmentID, MinorName:"COMPUTER STUFF"}
	db.Create(&minor1)
	db.Create(&minor2)
	db.Create(&minor3)
	db.Create(&minor4)

}




