package model


import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"

	"fmt"
)

type MainUser struct {
	UserID uint `gorm:"primary_key"`
	UserEmail string `gorm:"type:varchar(20);unique"`
	UserPassword string `gorm:"type:varchar(300)"`
	FirstName string `gorm:"type:varchar(50)"`
	LastName string `gorm:"type:varchar(50)"`
	UserType int `gorm:"not null"`
}

func (u *MainUser)BeforeCreate(){
	// hash the password text, and save it as the password
	println("User object created: ", u.UserEmail )
	hashedPW, err := bcrypt.GenerateFromPassword( []byte(u.UserPassword), 10)
	if (err != nil){
		fmt.Println("Problem hashing...", err)
	}
	u.UserPassword = string(hashedPW)
}

func (u *MainUser)CheckPasswordMatch(plainTextPassword string)(bool){
	// returns true is hashed plainTest matches hashed PW in database
	dbPassword := u.UserPassword
	check := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(plainTextPassword))
	if check != nil {
		fmt.Println("Error comparing passwords: ", check)
		return false
	} else {
		return true
	}

}


type Student struct {
	StudentID uint `gorm:"primary_key"`
	MainUser  MainUser `gorm:"ForeignKey:UserID;AssociationForeignKey:StudentID"`
	StudentType int `gorm:"not null"`
}

type PartTimeStudent struct {
	PartTimeStudentID uint `gorm:"primary_key"`
	Student  Student `gorm:"ForeignKey:StudentID"`
	StudentRefer uint `gorm:"not null"`
	NumCredits int `gorm:"not null"`
}

type FullTimeStudent struct {
	PartTimeStudentID uint `gorm:"primary_key"`
	Student  Student `gorm:"ForeignKey:StudentID"`
	StudentRefer uint `gorm:"not null"`
	NumCredits int `gorm:"not null"`
}


type Department struct {
	DepartmentID uint `gorm:"primary_key" `
	DepartmentName string `gorm:"type:varchar(30);not null"`
	DepartmentChair string `gorm:"type:varchar(30);not null"`
	DepartmentBuilding string `gorm:"type:varchar(30);not null"`
	DepartmentPhoneNumber  string `gorm:"type:varchar(15);not null"`
	DepartmentRoomNumber string `gorm:"type:varchar(10);not null"`
	Faculty []Faculty
}


type Faculty struct {
	FacultyID uint `gorm:"primary_key" `
	FacultyType int `gorm:"not null"`
	MainUser  MainUser `gorm:"ForeignKey:UserID; AssociationForeignKey:FacultyID"`
	DepartmentID uint `gorm:"not null"`
	Department Department `gorm:"ForeignKey:DepartmentID"`

}

type PartTimeFaculty struct {
	FacultyID uint `gorm:"primary_key" `
	Faculty Faculty `gorm:"ForeignKey:FacultyID; AssociationForeignKey:FacultyID"`
}

type FullTimeFaculty struct {
	FacultyID uint `gorm:"primary_key" `
	Faculty Faculty `gorm:"ForeignKey:FacultyID; AssociationForeignKey:FacultyID"`
}






