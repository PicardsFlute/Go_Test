package main

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	UserID uint `gnorm:"primary_key"`
	UserEmail string `gnorm:"type:varchar(20);unique"`


}