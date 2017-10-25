package model


import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

type User struct {
	UserID uint `gorm:"primary_key"`
	UserEmail string `gorm:"type:varchar(20);unique"`
	UserPassword string `gorm:"type:varchar(300)"`
	FirstName string `gorm:"type:varchar(50)"`
	LastName string `gorm:"type:varchar(50)"`
	UserType int `gorm:"not null"`
}

func (u *User)BeforeCreate(){
	// hash the password text, and save it as the password
	println("User object created: ", u.UserEmail )
}

func (u *User)CheckPasswordMatch(plainTextPassword string)(bool){
	// returns true is hashed plainTest matches hashed PW in database
	return false
}


type Student struct {
	StudentID uint `gorm:"primary_key"`
	User  User `gorm:"ForeignKey:UserRefer"`
	UserRefer uint `gorm:"not null"`
	StudentType int `gorm:"not null"`
}

type PartTimeStudent struct {
	PartTimeStudentID uint `gnorm:"primary_key"`
	Student  Student `gorm:"ForeignKey:UserRefer"`
	StudentRefer uint `gorm:"not null"`
	NumCredits int `gorm:"not null"`
}

type FullTimeStudent struct {
	PartTimeStudentID uint `gnorm:"primary_key"`
	Student  Student `gorm:"ForeignKey:StudentRefer"`
	StudentRefer uint `gorm:"not null"`
	NumCredits int `gorm:"not null"`
}


type Department struct {
	DepartmentID uint
	DepartmentName string `gorm:"type:varchar(30);not null"`
	DepartmentChair string `gorm:"type:varchar(30);not null"`
	DepartmentBuilding string `gorm:"type:varchar(30);not null"`
	DepartmentPhoneNumber  string `gorm:"type:varchar(15);not null"`
	departmentRoomNumber string `gorm:"type:varchar(10);not null"`

}


type Faculty struct {
	FacultyID uint `gnorm:"primary_key"`
	FacultyType int `gorm:"not null"`
	User  User `gorm:"ForeignKey:UserRefer"`
	UserRefer uint `gorm:"not null"`
	Department Department `gorm:"ForeignKey:DepartmentRefer"`
	DepartmentRefer uint `gorm:"not null"`
}
