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

	admin1 := model.Admin{AdminID:user8.UserID}
	admin2 := model.Admin{AdminID:user9.UserID}
	db.Create(&admin1)
	db.Create(&admin2)

	researcher1 := model.Researcher{ResearcherID:user10.UserID}
	db.Create(researcher1)
}



