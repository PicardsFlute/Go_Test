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
	//db.DropTable(&model.Student{})
	db.DropTable(&model.FullTimeStudent{})
	db.DropTable(&model.PartTimeStudent{})
	//db.DropTable(&model.MainUser{})
	//db.DropTable(&model.Faculty{})
	db.DropTable(&model.Department{})


	db.AutoMigrate(
		&model.MainUser{},
		&model.Student{},
		&model.PartTimeStudent{},
		&model.FullTimeStudent{},
		&model.Department{},
		&model.Faculty{},
	)
	db.Model(&model.Student{}).AddForeignKey("student_id", "main_user(user_id)", "RESTRICT", "RESTRICT")

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

	//dbUser := model.User{}
	//db.Where(&model.User{UserEmail: user1.UserEmail}).First(&dbUser)
	student1 := model.Student{MainUser: user1}
	db.Create(&student1)
	////
	//////db.Where(&model.User{UserEmail: user2.UserEmail}).First(&dbUser)
	student2 := model.Student{MainUser: user3}
	db.Create(&student2)

	//var user model.User
	//db.Model(&user).Association("Students")
	//db.Model(&).Related().Find(&student1)
	fmt.Println("For the student1, the user email is: ", student1.MainUser.UserEmail)

}



