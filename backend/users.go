package main

import (
	gorm "gorm.io/gorm"
)


type User struct {
	Uid	   int    `gorm:"column:uid;primaryKey;autoIncrement"`
	UName     string `gorm:"column:uname;type:varchar(16);not null;unique;index"`
	UPassword string `gorm:"column:password;type:varchar(32);not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// ULoginTime string `json:"ulogintime"`
}