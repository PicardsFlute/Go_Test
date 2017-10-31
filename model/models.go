package model


import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"time"
)

type MainUser struct {
	UserID uint `gorm:"primary_key"`
	UserEmail string `gorm:"type:varchar(50);unique"`
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
	Student  Student `gorm:"ForeignKey:StudentID;AssociationForeignKey:PartTimeStudentID"`
	NumCredits int `gorm:"not null"`
}

type FullTimeStudent struct {
	FullTimeStudentID uint `gorm:"primary_key"`
	Student  Student `gorm:"ForeignKey:StudentID;AssociationForeignKey:FullTimeStudentID"`
	NumCredits int `gorm:"not null"`
}


type Department struct {
	DepartmentID uint `gorm:"primary_key" `
	DepartmentName string `gorm:"type:varchar(30);not null"`
	DepartmentChair string `gorm:"type:varchar(50);not null"`
	DepartmentBuilding string `gorm:"type:varchar(50);not null"`
	DepartmentPhoneNumber  string `gorm:"type:varchar(15);not null"`
	DepartmentRoomNumber string `gorm:"type:varchar(10);not null"`
	Faculty []Faculty
}


type Faculty struct {
	FacultyID uint `gorm:"primary_key" `
	FacultyType int `gorm:"not null"`
	MainUser  MainUser `gorm:"ForeignKey:UserID; AssociationForeignKey:FacultyID"`
	DepartmentID uint `gorm:"not null"`
	Department Department `gorm:"ForeignKey:DepartmentID; AssociationForeignKey:DepartmentID"`

}

type PartTimeFaculty struct {
	PartTimeFacultyID uint `gorm:"primary_key" `
	Faculty Faculty `gorm:"ForeignKey:FacultyID; AssociationForeignKey:PartTimeFacultyID"`
}

type FullTimeFaculty struct {
	FullTimeFacultyID uint `gorm:"primary_key" `
	Faculty Faculty `gorm:"ForeignKey:FacultyID; AssociationForeignKey:FullTimeFacultyID"`
}


type Admin struct {
	AdminID uint `gorm:"primary_key"`
	MainUser MainUser `gorm:"ForeignKey:UserID;AssociationForeignKey:AdminID"`
}

type Researcher struct {
	ResearcherID uint `gorm:"primary_key"`
	MainUser MainUser `gorm:"ForeignKey:UserID;AssociationForeignKey:ResearcherID"`
}

type Major struct {
	MajorID uint `gorm:"primary_key"`
	DepartmentID uint `gorm:"not null"`
	MajorName string `gorm:"type:varchar(59)"`
	Department Department `gorm:"ForeignKey:DepartmentID; AssociationForeignKey:DepartmentID"`

}

type Minor struct {
	MinorID uint `gorm:"primary_key"`
	MinorName string `gorm:"type:varchar(50)"`
	DepartmentID uint `gorm:"not null"`
	Department Department `gorm:"ForeignKey:DepartmentID; AssociationForeignKey:DepartmentID"`
}

type StudentMajor struct {
	StudentID uint `gorm:"primary_key "`
	MajorID uint `gorm:"primary_key "`
	Major Major `gorm:"ForeignKey:MajorID; AssociationForeignKey:MajorID"`
	Student Student `gorm:"ForeignKey:StudentID; AssociationForeignKey:StudentID"`
}

type StudentMinor struct {
	StudentID uint `gorm:"primary_key "`
	MinorID uint `gorm:"primary_key "`
	Minor Minor `gorm:"ForeignKey:MinorID; AssociationForeignKey:MinorID"`
	Student Student `gorm:"ForeignKey:StudentID; AssociationForeignKey:StudentID"`
}


type Course struct {
	CourseID uint `gorm:"primary_key"`
	CourseName string `gorm:"type:varchar(50)"`
	CourseCredits int `gorm:"type:integer"`
	CourseDescription string `gorm:"type:text"`
	DepartmentID uint `gorm:"not null"`
	Department Department `gorm:"ForeignKey:DepartmentID; AssociationForeignKey:DepartmentID"`
}


type Prerequisite struct {
	CourseRequiredBy uint `gorm:"primary_key"`
	CourseRequirement uint `gorm:"primary_key"`
	Course Course  `gorm:"ForeignKey:CourseID; AssociationForeignKey:CourseRequiredBy"`
	CourseRequired Course  `gorm:"ForeignKey:CourseID; AssociationForeignKey:CourseRequirement"`
}

type Day struct{
	DayID uint `gorm:"primary_key"`
	MeetingDay string `gorm:"type:varchar(50)"`
}

type Semester struct {
	SemesterID uint `gorm:"primary_key"`
	Year int `gorm:"type:integer"`
	Season string `gorm:"type:varchar(50)"`
}

type Period struct {
	PeriodID uint `gorm:"primary_key"`
	StartTime time.Time
	EndTime time.Time
}

type TimeSlot struct{
	TimeSlotID uint `gorm:"primary_key"`
	PeriodID uint `gorm:"not null"`
	SemesterID uint `gorm:"not null"`
	DayID uint `gorm:"not null"`
	Period Period  `gorm:"ForeignKey:PeriodID; AssociationForeignKey:PeriodID"`
	Semester Semester  `gorm:"ForeignKey:SemesterID; AssociationForeignKey:SemesterID"`
	Day Day  `gorm:"ForeignKey:DayID; AssociationForeignKey:DayID"`
}


