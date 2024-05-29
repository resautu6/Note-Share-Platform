package main

import (
	"fmt"
	"strconv"
	"os"

	log "github.com/sirupsen/logrus"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

type DataBase struct {
	UName     string
	UPassword string
	Host string
	Port int
	DataBaseName string
	db 	  *gorm.DB
}

var (
	db *DataBase
)

func initDb() {
	db = &DataBase{
		UName: Config.DataBase_uname,
		UPassword: Config.DataBase_password,
		Host: Config.DataBase_host,
		Port: Config.DataBase_port,
		DataBaseName: Config.DataBase_name,
		db: nil,
	}
	db.connect()
}

func (db *DataBase) connect() {
	// connect to database
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", db.UName, db.Host, db.Port, db.DataBaseName)
	log.Info(dsn)
	// dsn := "root@tcp(127.0.0.1:3306)/nsptest?charset=utf8&parseTime=True&loc=Local"
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

func (db *DataBase) addArticle(aritcle Article) {
	db.db.Create(&aritcle)

	articlePath := "res/" + strconv.FormatInt(int64(aritcle.ArticleUid), 10) + "_" +  strconv.FormatInt(int64(aritcle.ArticleId), 16)
	err := os.Mkdir(articlePath, 0755)
	if err != nil {
		log.Error("Create article directory failed: ", err)
	}

	db.db.Where("article_id = ?", aritcle.ArticleId).Update("image_path", articlePath)
}

func (db *DataBase) getArticles() []Article {
	var articles []Article
	err := db.db.Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) getArticleById(id int) Article {
	var article Article
	err := db.db.First(&article, id).Error
	if err != nil {
		log.Warn(err)
	}
	return article
}

func (db *DataBase) getArticleByUid(uid int) []Article {
	var articles []Article
	err := db.db.Where("uid = ?", uid).Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) getArticleByTitle(title string) []Article {
	var articles []Article
	err := db.db.Where("title = ?", title).Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) getArticleByContent(content string) []Article {
	var articles []Article
	err := db.db.Where("content = ?", content).Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

