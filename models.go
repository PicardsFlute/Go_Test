package main

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	UserID uint `gnorm:"primary_key"`
	UserEmail string `gnorm:"type:varchar(20);unique"`
	UserPassword string `gnorm:"type:varchar(300)"`
	FirstName string `gnorm:"type:varchar(50)"`
	LastName string `gnorm:"type:varchar(50)"`
	UserType int
}

type Student struct {
	StudentID uint `gnorm:"primary_key"`
	User  User `gorm:"ForeignKey:UserRefer"`
	UserRefer uint
}