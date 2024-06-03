package main

// import (
// 	log "github.com/sirupsen/logrus"
// )

type Cache struct {
	cache map[string]interface{}

}

func (cch *Cache) hasContent(key string) bool {
	_, ok := cch.cache[key]
	return ok
}

func (cch *Cache) getContent(key string) interface{}{
	return cch.cache[key]
}

func (cch *Cache) setContent(key string, value interface{}) {
	cch.cache[key] = value
}


type ArticleCache struct {
	Cache
}

func (artch *ArticleCache) init() {
	artch.cache = make(map[string]interface{})
}

type UserCache struct {
	Cache
}

