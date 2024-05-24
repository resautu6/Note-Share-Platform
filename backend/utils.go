package main

import (
	"encoding/json"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Port      int    `json:"Port"`
	Ipaddr    string `json:"Ipaddr"`
	SecretKey string `json:"Secret_Key"`
    DataBase_host string `json:"DataBase_host"`
    DataBase_port int `json:"DataBase_port"`
    DataBase_uname string `json:"DataBase_uname"`
    DataBase_password string `json:"DataBase_password"`
	DataBase_name string `json:"DataBase_name"`
}

var(
    Config config
)

func loadConfigFile(path string) {
	if path == "" {
		path = "./config.json"
	}

	file, err := os.Open(path)
    if err != nil {
        log.Error(err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&Config)
    if err != nil {
        log.Error(err)
    }
    
}

func generateJWT(uid int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid" : 	uid,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // set duration=24h
	})

	// signature token with secret key
	tokenString, err := token.SignedString([]byte(Config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
