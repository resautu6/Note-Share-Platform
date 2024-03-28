package main

import (
	//"fmt"

	log "github.com/sirupsen/logrus"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

type DataBase struct {
	UName     string
	UPassword string
	Host string
	Port int
	db 	  *gorm.DB
}

var (
	db *DataBase
)

func initDb() {
	db = &DataBase{
		UName: Config.DataBase_name,
		UPassword: Config.DataBase_password,
		Host: Config.DataBase_host,
		Port: Config.DataBase_port,
		db: nil,
	}
	db.connect()
}

func (db *DataBase) connect() {
	// connect to database
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", db.UName, db.UPassword, Config.Ipaddr, Config.Port, "users")
	dsn := "root@tcp(127.0.0.1:3306)/nsptest?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	db.db = conn
}

func (db *DataBase) close() {

}

func (db *DataBase) addUser(user User) {
	db.db.Create(&user)
}

func (db *DataBase) getUserByName(name string) User {
	var user User
	err := db.db.Where("uname = ?", name).First(&user).Error
	if err != nil {
		log.Warn(err)
	}
	return user
}

func (db *DataBase) getUserById(id int) User {
	var user User
	err := db.db.First(&user, id).Error
	if err != nil {
		log.Warn(err)
	}
	return user
}

func (db *DataBase) getUserByNameAndPassword(name string, password string) User {
	var user User
	err := db.db.Where("uname = ? AND password = ?", name, password).First(&user).Error
	if err != nil {
		log.Warn(err)
	}
	return user
}
