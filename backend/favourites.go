package main

import (
	gorm "gorm.io/gorm"
)

type Favourites struct {
	FavouriteId int `gorm:"column:favourite_id;primaryKey;autoIncrement"`
	FavouriteUid int `gorm:"column:uid;type:int;not null;index"`
	FavouriteArticleId int `gorm:"column:article_id;type:int;not null;index"`
	FavouritesModifytime string `gorm:"column:modify_time;type:TIMESTAMP;not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}