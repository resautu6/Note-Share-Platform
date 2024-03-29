package main

import (
	"fmt"
	"sync"
	"net/http"

	log "github.com/sirupsen/logrus"
	gin "github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
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
		check_user := db.getUserByName(c.PostForm("name"))
		if check_user.UName != "" {
			c.JSON(401, gin.H{
				"message": "User name has exist",
			})
			return
		}

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

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // provide secretkey
            return []byte(Config.SecretKey), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        // 将用户信息存储在请求上下文中
        c.Set("user", claims)
        c.Next()
    }
}