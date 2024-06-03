package main

import (
	"time"
	"strconv"
	// "os"

	gorm "gorm.io/gorm"
	// log "github.com/sirupsen/logrus"
)

type Article struct {
	ArticleId int `gorm:"column:article_id;primaryKey;autoIncrement"`
	ArticleTitle string `gorm:"column:title;type:varchar(32);not null;index"`
	ArticleUid int `gorm:"column:uid;type:int;not null;index"`
	ArticleContent string `gorm:"column:content;type:varchar(8196);not null;index"`
	ArticleModifyTime time.Time `gorm:"column:modify_time;type:TIMESTAMP;not null;index"`
	ArticleImagePath string `gorm:"column:image_path;type:varchar(64);not null;"`
	ArticleImageNum int `gorm:"column:image_num;type:int;not null;"`
	ArticleViewNum int `gorm:"column:view_num;type:int;not null;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func makeArticle(title string, content string, uid int, imageNum int, viewNum int) Article {
	timeStamp := time.Now()
	millisTimestamp := timeStamp.UnixMilli()

	articlePath := "res/" + strconv.FormatInt(int64(uid), 10) + "_" +  strconv.FormatInt(millisTimestamp, 16)

	ret := Article {
		ArticleTitle: title,
		ArticleUid: uid,
		ArticleContent: content,
		ArticleModifyTime: time.Now(),
		ArticleImagePath: articlePath,
		ArticleImageNum: imageNum,
		ArticleViewNum: viewNum,
	}

	return ret
}