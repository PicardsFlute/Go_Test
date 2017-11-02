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
	DepartmentPhoneNumber  string `gorm:"type:varchar(25);not null"`
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



type Room struct {
	RoomID uint `gorm:"primary_key"`
	RoomType string `gorm:"type:varchar(50)"`
	RoomNumber string `gorm:"type:varchar(50)"`
}

type Building struct {
	BuildingID uint `gorm:"primary_key"`
	BuildingName string `gorm:"type:varchar(150)"`
	BuildingAddress string `gorm:"type:varchar(200)"`
}

type Location struct {
	LocationID uint `gorm:"primary_key"`
	BuildingID uint `gorm:"not null"`
	RoomID uint `gorm:"not null"`
	Building Building `gorm:"ForeignKey:BuildingID; AssociationForeignKey:BuildingID"`
	Room Room `gorm:"ForeignKey:RoomID; AssociationForeignKey:RoomID"`
}

type Section struct {
	SectionID uint `gorm:"primary_key"`
	CourseSectionNumber int `gorm:"not null"`
	CourseID uint `gorm:"not null"`
	FacultyID uint `gorm:"not null"`
	TimeSlotID uint `gorm:"not null"`
	LocationID uint `gorm:"not null"`
	Course Course `gorm:"ForeignKey:CourseID; AssociationForeignKey:CourseID"`
	TimeSlot TimeSlot `gorm:"ForeignKey:TimeSlotID; AssociationForeignKey:TimeSlotID"`
	Faculty Faculty `gorm:"ForeignKey:FacultyID; AssociationForeignKey:FacultyID"`
	Location Location `gorm:"ForeignKey:LocationID; AssociationForeignKey:LocationID"`

}

type Enrollment struct {
	EnrollmentID uint `gorm:"primary_key"`
	Grade uint `gorm:"type:varchar(5)"`
	StudentID uint `gorm:"not null"`
	SectionID uint `gorm:"not null"`
	Student Student  `gorm:"ForeignKey:StudentID; AssociationForeignKey:StudentID"`
	Section Section  `gorm:"ForeignKey:SectionID; AssociationForeignKey:SectionID"`
}

type Attends struct {
	EnrollmentID uint `gorm:"primary_key"`
	StudentID uint `gorm:"primary_key"`
	AttendsDate time.Time `gorm:"primary_key"`
	Present bool `gorm:"primary_key"`
	Enrollment Enrollment  `gorm:"ForeignKey:EnrollmentID; AssociationForeignKey:EnrollmentID"`
	Student Student  `gorm:"ForeignKey:StudentID; AssociationForeignKey:StudentID"`
}

type StudentHistory struct {
	StudentID uint `gorm:"primary_key"`
	EnrollmentID uint `gorm:"primary_key"`
	Status string `gorm:"primary_key"`
	Student Student  `gorm:"ForeignKey:StudentID; AssociationForeignKey:StudentID"`
	Enrollment Enrollment  `gorm:"ForeignKey:EnrollmentID; AssociationForeignKey:EnrollmentID"`
}

type Reports struct {
	ReportID uint `gorm:"primary_key"`
	DateCreated time.Time
	Description string  `gorm:"type:text"`
	ReportPath string  `gorm:"type:text"`
	ResearcherID uint `gorm:"not null"`
	Researcher Researcher `gorm:"ForeignKey:ResearcherID; AssociationForeignKey:ResearcherID"`
}
