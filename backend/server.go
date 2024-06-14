package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	gin "github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/hashicorp/golang-lru"
)

type Server struct {
	ipaddr         string
	port           int
	router         *gin.Engine

	articleCache *lru.Cache
	userCache *lru.Cache
}

func (s *Server) registerRouter() {
	s.handleRootGet()
	s.handleLoginPost()
	s.handleRegisterPost()
	s.handleGetArticleContent()
	s.handleUploadArticle()
	s.handleGetUserInform()
	s.handleModifyArticle()

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
	var err error
	s.userCache, err = lru.New(5)
	if err != nil {
		log.Error("user cache badly init")
		return
	}

	s.articleCache, err = lru.New(5)
	if err != nil {
		log.Error("article cache badly init")
		return
	}


	lru.New(5)

	s.ipaddr = ipaddr
	s.port = port
	s.router = gin.Default()
	s.registerMiddleware()
	s.registerRouter()
	s.router.Static("/res", "./res")
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

		s.userCache.Add(user.Uid, user)

		jwtToken, _ := generateJWT(user.Uid, user.UName)

		c.JSON(200, gin.H{
			"message": "Login success",
			"uname":   user.UName,
			"token":   jwtToken,
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
	s.router.POST("/upload_article", authMiddleware(), func(c *gin.Context) {

		claimsInterface, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{"message": "Token not found"})
			return
		}
		claims := claimsInterface.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))

		imageNum, _ := strconv.Atoi(c.PostForm("image_num"))
		article := makeArticle(c.PostForm("article_title"), c.PostForm("article_content"), uid, imageNum, 0)
		article = db.addArticle(article)
		imgPath := article.ArticleImagePath

		if imgPath == "error" {
			c.JSON(403, gin.H{"message": "文章上传失败"})
			db.deleteArticleById(article.ArticleId)
			return
		}

		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(403, gin.H{"message": "图片上传失败"})
			db.deleteArticleById(article.ArticleId)
			return
		}
		imageList := form.File["image_list"]
		if len(imageList) != imageNum {
			c.JSON(403, gin.H{"message": "图片数量不匹配"})
			db.deleteArticleById(article.ArticleId)
			return
		}
		for idx, file := range imageList {
			// 读取文件
			src, err := file.Open()
			if err != nil {
				c.String(http.StatusBadRequest, "file open err: %s", err.Error())
				db.deleteArticleById(article.ArticleId)
				return
			}
			defer src.Close()

			// 读取文件内容
			fileBytes, err := io.ReadAll(src)
			if err != nil {
				c.String(http.StatusBadRequest, "file read err: %s", err.Error())
				db.deleteArticleById(article.ArticleId)
				return
			}

			// 创建文件并保存到res目录
			fileName := strconv.Itoa(idx)
			dst := filepath.Join(imgPath, fileName+".png")
			err = os.WriteFile(dst, fileBytes, 0644)
			if err != nil {
				c.String(http.StatusBadRequest, "file write err: %s", err.Error())
				db.deleteArticleById(article.ArticleId)
				return
			}
		}

		

		s.articleCache.Add(article.ArticleId, article)

		c.JSON(200, gin.H{
			"message": "发布成功！",
		})
	})
}

func (s *Server) handleGetArticleContent() {
	s.router.GET("article/list", func(c *gin.Context) {
		articles := db.getArticles(8)
		article_ids := make([]int, 0)
		for _, article := range articles {
			article_ids = append(article_ids, article.ArticleId)
		}
		c.JSON(200, gin.H{
			"item_sum": len(article_ids),
			"items":    article_ids,
		})
	})

	s.router.GET("/article/:id", func(c *gin.Context) {
		var article Article
		article_id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(403, gin.H{
				"message": "Wrong article id!",
			})
			c.Status(403)
			return
		}
		if s.articleCache.Contains(article_id) {
			tmp, _ := s.articleCache.Get(article_id)
			article = tmp.(Article)
		} else {
			article = db.getArticleById(article_id)
		}

		if article.ArticleId == -1 {
			c.JSON(403, gin.H{
				"message": "Article not found!",
			})
			return
		}

		s.articleCache.Add(article_id, article)

		c.JSON(200, gin.H{
			"id":         article.ArticleId,
			"title":      article.ArticleTitle,
			"uname":      article.ArticleUid,
			"image_num":  article.ArticleImageNum,
			"image_path": article.ArticleImagePath,
			"view_num":   article.ArticleViewNum,
			"modify_time": article.ArticleModifyTime,
		})
	})

	s.router.GET("/search/article/:word", func(c *gin.Context) {
		articles := db.getArticleByTitleAndContent(c.Param("word"), 16)
		article_ids := make([]int, 0)
		for _, article := range articles {
			article_ids = append(article_ids, article.ArticleId)
		}

		c.JSON(200, gin.H{
			"item_sum": len(article_ids),
			"items":    article_ids,
		})
	})

	s.router.GET("/article/:id/content", func(c *gin.Context) {
		var article Article
		articleID, _ := strconv.Atoi(c.Param("id"))
		if s.articleCache.Contains(articleID) {
			tmp, _ := s.articleCache.Get(articleID)
			article = tmp.(Article)
		} else {
			article = db.getArticleById(articleID)
		}

		if article.ArticleId == -1 {
			c.JSON(403, gin.H{
				"message": "Article not found!",
			})
			return
		}

		db.updateViewNumByAid(articleID)

		article.ArticleViewNum += 1
		s.articleCache.Add(articleID, article)

		c.JSON(200, gin.H{
			"id":         c.Param("id"),
			"content":    article.ArticleContent,
			"view_num":   article.ArticleViewNum,
			"image_path": article.ArticleImagePath,
		})
	})
}

func (s *Server) handleModifyArticle() {
	s.router.POST("/article/:id/modify", authMiddleware(), func(c *gin.Context) {
		claimsInterface, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{"message": "Token not found"})
			return
		}
		claims := claimsInterface.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))

		articleID, _ := strconv.Atoi(c.Param("id"))

		command := c.PostForm("command")
		if command == "" {
			c.JSON(403, gin.H{"message": "Command not found"})
			return
		}

		if command == "delete" {

			article := db.getArticleById(articleID)
			if article.ArticleId == -1 || article.ArticleUid != uid {
				c.JSON(403, gin.H{"message": "Article not found or you are not the author"})
				return
			} 
			db.deleteArticleByAid(articleID)

			if s.articleCache.Contains(articleID) {
				s.articleCache.Remove(articleID)
			}

			c.JSON(200, gin.H{"message": "Article deleted successfully"})
			return

		}
	})
}

func (s *Server) handleGetUserInform() {
	s.router.GET("/user/favourites", authMiddleware(), func(c *gin.Context) {
		claimsInterface, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{"message": "Token not found"})
			return
		}
		claims := claimsInterface.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))

		favourites := db.getFavouritesByUid(uid)
		favourites_ids := make([]int, 0)
		for _, favourite := range favourites {
			favourites_ids = append(favourites_ids, favourite.FavouriteArticleId)
		}

		c.JSON(200, gin.H{
			"item_sum": len(favourites_ids),
			"items":    favourites_ids,
		})
	})

	s.router.POST("/user/favourites", authMiddleware(), func(c *gin.Context) {
		claimsInterface, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{"message": "Token not found"})
			return
		}
		claims := claimsInterface.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))

		article_id, err := strconv.Atoi(c.PostForm("article_id"))
		command := c.PostForm("command")
		if command == "" {
			c.JSON(403, gin.H{"message": "Command not found"})
			return
		}
		if err != nil {
			c.JSON(403, gin.H{"message": "Wrong article id"})
			return
		}
        
        // log.Info(command)
        
		if command == "delete" {
			db.deleteFavouriteByAidAndUid(article_id, uid)
			c.JSON(200, gin.H{"message": "Delete favourite success"})
			return
		} else if command != "add" {
			c.JSON(403, gin.H{"message": "Command not found"})
			return
		}

		article := db.getArticleById(article_id)

		if article.ArticleId == -1 {
			c.JSON(403, gin.H{"message": "Article not found"})
			return
		}

		testList := db.getFavouritesByUidAndAid(uid, article_id)
		if len(testList) != 0 {
			c.JSON(403, gin.H{"message": "Article already in favourite list"})
			return
		}

		favourite := makeFavourite(uid, article_id)
		err = db.addFavorite(favourite)

		if err != nil {
			c.JSON(403, gin.H{"message": "Add favourite failed"})
			return
		}

		c.JSON(200, gin.H{"message": "Add favourite success"})

	})

	s.router.GET("/user/uname/:id", func(c *gin.Context) {
		uid, _ := strconv.Atoi(c.Param("id"))
		var user User
		if s.userCache.Contains(uid) {
			tmp, _ := s.userCache.Get(uid)
			user = tmp.(User)
		} else {
			user = db.getUserById(uid)
			s.userCache.Add(uid, user)
		}

		if user.Uid != uid {
			c.JSON(403, gin.H{"message": "User not found"})
			return
		}
		c.JSON(200, gin.H{
			"uname": user.UName,
		})
	})

	s.router.GET("/user/article", authMiddleware(), func(c *gin.Context) {
		claimsInterface, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{"message": "Token not found"})
			return
		}
		claims := claimsInterface.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))

		articles := db.getArticlesByUid(uid)
		article_ids := make([]int, 0)
		for _, article := range articles {
			article_ids = append(article_ids, article.ArticleId)
		}
		c.JSON(200, gin.H{
			"item_sum": len(article_ids),
			"items":    article_ids,
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
