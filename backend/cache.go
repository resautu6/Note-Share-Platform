package main

import (
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	cache map[int]interface{}

}

func (cch *Cache) getContent(key int) interface{}{
	return cch.cache[key]
}

func (cch *Cache) setContent(key int, value interface{}) {
	cch.cache[key] = value
}


type ArticleCache struct {
	Cache
}

func (artch *ArticleCache) init() {
	artch.cache = make(map[int]interface{})
}

type UserCache struct {
	Cache
}

