package main

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	gin "github.com/gin-gonic/gin"
)

type Server struct{
	ipaddr string
	port int
	router *gin.Engine
	usersMap map[string]User
	usersMapRWLock sync.RWMutex
}

func (s *Server) registerRouter() {
	s.handleRootGet()
	s.handleLoginPost()
	s.handleRegisterPost()

}

func (s *Server) start(ipaddr string, port int) {
	log.Info("Server started")
	initDb()

	s.ipaddr = ipaddr
	s.port = port
	s.router = gin.Default()
	s.usersMap = make(map[string]User)
	s.registerRouter()
	s.router.Run(fmt.Sprintf("%s:%d", s.ipaddr, s.port))
}

func (s *Server) handleRootGet() {
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the server",
		})
	})
}

func (s *Server) handleLoginPost() {
	s.router.POST("/login", func(c *gin.Context) {
		user := db.getUserByNameAndPassword(c.PostForm("name"), c.PostForm("password"))
		if user.UName == "" {
			c.JSON(401, gin.H{
				"message": "Login failed",
			})
			return
		}

		s.usersMapRWLock.Lock()
		s.usersMap[user.UName] = user
		s.usersMapRWLock.Unlock()

		c.JSON(200, gin.H{
			"message": "Login success",
		})
	})
}

func (s *Server) handleRegisterPost() {
	s.router.POST("/register", func(c *gin.Context) {
		user := User{
			UName: c.PostForm("name"),
			UPassword: c.PostForm("password"),
		}

		db.addUser(user)
		c.JSON(200, gin.H{
			"message": "Register success",
		})
	})
}