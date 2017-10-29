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
	db.DropTable(&model.Student{})
	db.DropTable(&model.FullTimeStudent{})
	db.DropTable(&model.PartTimeStudent{})
	db.DropTable(&model.Faculty{})
	db.DropTable(&model.MainUser{})

	db.DropTable(&model.Department{})


	db.AutoMigrate(
		&model.MainUser{},
		&model.Student{},
		&model.PartTimeStudent{},
		&model.FullTimeStudent{},
		&model.Department{},
		&model.Faculty{},
	)
	db.Model(&model.Student{}).AddForeignKey("student_id", "main_user(user_id)", "CASCADE", "CASCADE")

	db.Model(&model.Faculty{}).AddForeignKey("faculty_id", "main_user(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Faculty{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	user1 := model.MainUser{FirstName: "Pat", LastName:"Lagat", UserEmail:"plagat@yahoo.com", UserPassword:"pl12345"}
	user2 := model.MainUser{FirstName: "Irish", LastName:"James", UserEmail:"jirish@yahoo.com", UserPassword:"ij12345"}
	user3 := model.MainUser{FirstName: "Trey", LastName:"Gorkin", UserEmail:"gork@yahoo.com", UserPassword:"tg12345", UserType:1}
	user4 := model.MainUser{FirstName: "Testy", LastName:"McTest", UserEmail:"test@test.test", UserPassword:"testPW", UserType:1}
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)
	
	student2 := model.Student{StudentID: user3.UserID, StudentType:2}
	db.Create(&student2)

	student1 := model.Student{StudentID: user1.UserID, StudentType:1}
	db.Create(&student1)

	// example of finding related models
	u := model.MainUser{}
	db.Model(&student2).Association("MainUser").Find(&u)
	fmt.Println("For the student2, the user email is: ", u.UserEmail)



	department1 := model.Department{DepartmentName:"Math", DepartmentBuilding:"MathBuilding", DepartmentRoomNumber:"302", DepartmentChair:"Mr Math", DepartmentPhoneNumber:"111-222-3333"}
	department2 := model.Department{DepartmentName:"Computer Science", DepartmentBuilding:"CSBuilding", DepartmentRoomNumber:"100", DepartmentChair:"Mr Computer", DepartmentPhoneNumber:"123-456-7899"}
	db.Create(&department1)
	db.Create(&department2)

	faculty1 := model.Faculty{FacultyID:user2.UserID, FacultyType:1, DepartmentID:department1.DepartmentID}
	faculty2 := model.Faculty{FacultyID:user4.UserID, FacultyType:2, DepartmentID:department2.DepartmentID}
	db.Create(&faculty1)
	db.Create(&faculty2)


}



