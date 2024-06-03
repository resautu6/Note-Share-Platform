package main

import (
	"time"

	gorm "gorm.io/gorm"
)

type Favourite struct {
	FavouriteId int `gorm:"column:favourite_id;primaryKey;autoIncrement"`
	FavouriteUid int `gorm:"column:uid;type:int;not null;index"`
	FavouriteArticleId int `gorm:"column:article_id;type:int;not null;index"`
	FavouritesModifytime time.Time `gorm:"column:modify_time;type:TIMESTAMP;not null;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func makeFavourite(uid int, articleId int) Favourite {
	return Favourite{
		FavouriteUid: uid,
		FavouriteArticleId: articleId,
		FavouritesModifytime: time.Now(),
	}
}