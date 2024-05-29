package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	ipaddr         string
	port           int
	router         *gin.Engine
	usersMap       map[string]User
	usersMapRWLock sync.RWMutex

	articleCache ArticleCache 
}

func (s *Server) registerRouter() {
	s.handleRootGet()
	s.handleLoginPost()
	s.handleRegisterPost()
	s.handleGetArticleContent()
}

func (s *Server) registerMiddleware() {
	s.router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func (s *Server) start(ipaddr string, port int) {
	log.Info("Server started")
	initDb()

	s.ipaddr = ipaddr
	s.port = port
	s.router = gin.Default()
	s.usersMap = make(map[string]User)
	s.registerMiddleware()
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
			"uname": user.UName,
			"token": "123",
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
			UName:     c.PostForm("name"),
			UPassword: c.PostForm("password"),
		}

		db.addUser(user)
		c.JSON(200, gin.H{
			"message": "Register success",
		})
	})
}

func (s *Server) handleUploadArticle() {
	s.router.POST("/upload_article", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "发布成功！",
		})
	})
}



func (s *Server) handleGetArticleContent() {
	s.router.GET("article/list", func(c *gin.Context) {
		articles := db.getArticles()
		c.JSON(200, gin.H{
			"item_sum" : 0,
			"items": articles,
		})
	})

	s.router.GET("/article/:id", func(c *gin.Context) {
		var article Article
		article_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(403, gin.H{
				"message" : "Wrong article id!",
			})
			c.Status(403)
			return
		}
 		if s.articleCache.hasContent(c.Param("id")) {
			article = s.articleCache.getContent(c.Param("id")).(Article)
		} else {
			article = db.getArticleById(article_id)
		}

		c.JSON(200, gin.H{
			"id": article.ArticleId,
			"title": article.ArticleTitle,
			"uname": article.ArticleUid,
			"image_num": article.ArticleImageNum,
			"image_path": article.ArticleImagePath,
			"view_num": article.ArticleViewNum,
		})
	})

	s.router.GET("/article/:id/content", func(c *gin.Context) {
		if s.articleCache.hasContent(c.Param("id")) {

		}

		c.JSON(200, gin.H{
			"id" : c.Param("id"),
			"content": "This is a test content",
			"image_path": "/",
		})
	})
}

func (s *Server) handleGetUserInform() {
	s.router.GET("/user/favourites/:id", func(c *gin.Context) {
		user := db.getUserById(c.Param("id"))
		c.JSON(200, gin.H{
			"user": user,
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
