package main

import (
	gorm "gorm.io/gorm"
)

type Article struct {
	ArticleId int `gorm:"column:article_id;primaryKey;autoIncrement"`
	ArticleTitle string `gorm:"column:title;type:varchar(16);not null;index"`
	AticleUid int `gorm:"column:uid;type:int;not null;index"`
	AticleContent string `gorm:"column:content;type:varchar(8196);not null;index"`
	AticleModifyTime string `gorm:"column:modify_time;type:TIMESTAMP;not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}