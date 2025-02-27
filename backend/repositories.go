package main

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

type DataBase struct {
	UName        string
	UPassword    string
	Host         string
	Port         int
	DataBaseName string
	db           *gorm.DB
}

var (
	db *DataBase
)

func initDb() {
	db = &DataBase{
		UName:        Config.DataBase_uname,
		UPassword:    Config.DataBase_password,
		Host:         Config.DataBase_host,
		Port:         Config.DataBase_port,
		DataBaseName: Config.DataBase_name,
		db:           nil,
	}
	db.connect()
}

func (db *DataBase) connect() {
	// connect to database
	dsn := fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", db.UName, db.Host, db.Port, db.DataBaseName)
	// log.Info(dsn)
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

func (db *DataBase) addArticle(aritcle Article) Article {
	result := db.db.Create(&aritcle)

	if result.Error != nil {
		aritcle.ArticleImagePath = "error"
		return aritcle
	}

	articlePath := "res/" + strconv.FormatInt(int64(aritcle.ArticleUid), 10) + "_" + strconv.FormatInt(int64(aritcle.ArticleId), 16)
	err := os.Mkdir(articlePath, 0755)
	if err != nil {
		log.Error("Create article directory failed: ", err)
	}
	aritcle.ArticleImagePath = articlePath
	result = db.db.Model(&aritcle).Update("image_path", articlePath)
	// result = db.db.Where("article_id = ?", aritcle.ArticleId).Update("image_path", articlePath)
	if result.Error != nil {
		aritcle.ArticleImagePath = "error"
		return aritcle
	}

	return aritcle
}

func (db *DataBase) getArticles(lmt int) []Article {
	if lmt == 0 {
		lmt = 16
	}
	var articles []Article
	err := db.db.Find(&articles).Limit(lmt).Error
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
		article.ArticleId = -1
	}
	return article
}

func (db *DataBase) deleteArticleById(id int) {
	db.db.Delete(&Article{}, id)
}

func (db *DataBase) getArticlesByUid(uid int) []Article {
	var articles []Article
	err := db.db.Where("uid = ?", uid).Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) updateViewNumByAid(aid int) {
	var article Article
	article.ArticleId = aid
	result := db.db.Table("articles").Where("article_id = ?", aid).Update("view_num", gorm.Expr("view_num + ?", 1))

	if result.Error != nil {
		log.Warn(result)
	}
}

func (db *DataBase) getArticleByTitle(title string) []Article {
	var articles []Article
	err := db.db.Table("articles").Where("MATCH(title) AGAINST(?)", title).Find(&articles).Error

	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) getArticleByContent(content string) []Article {
	var articles []Article
	err := db.db.Table("articles").Where("MATCH(content) AGAINST(?)", content).Find(&articles).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) getArticleByTitleAndContent(content string, lmt int) []Article {
	var articles []Article
	if lmt == 0 {
		lmt = 114514
	}

	err := db.db.Table("articles").Where("MATCH(title, content) AGAINST(?)", content).Find(&articles).Limit(lmt).Error
	if err != nil {
		log.Warn(err)
	}
	return articles
}

func (db *DataBase) deleteArticleByAid(aid int) {
	db.db.Delete(&Article{}, "article_id = ?", aid)
}

func (db *DataBase) addFavorite(favourite Favourite) error {
	result := db.db.Create(&favourite)
	return result.Error
}

func (db *DataBase) getFavouritesByUid(uid int) []Favourite {
	var favourites []Favourite
	err := db.db.Where("uid = ?", uid).Find(&favourites).Error
	if err != nil {
		log.Warn(err)
	}
	return favourites
}

func (db *DataBase) getFavouritesByFid(fid int) []Favourite {
	var favourites []Favourite
	err := db.db.Where("favourite_id = ?", fid).Find(&favourites).Error
	if err != nil {
		log.Warn(err)
	}
	return favourites
}

func (db *DataBase) getFavouritesByUidAndAid(uid int, aid int) []Favourite {
	var favourites []Favourite
	err := db.db.Where("uid = ? AND article_id = ?", uid, aid).Find(&favourites).Error
	if err != nil {
		log.Warn(err)
	}
	return favourites
}

func (db *DataBase) deleteFavouriteByAidAndUid(aid int, uid int) {
	db.db.Delete(&Favourite{}, "article_id = ? AND uid = ?", aid, uid)
}
