package routers

import (
	"backend/codeasset/common/middleware/logger"
	"embed"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"backend/codeasset/controls/my"
	"backend/codeasset/controls/softinfo"
)

func Register(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, "OK")
	})
	router.GET("/enableqps", func(c *gin.Context) {
		logger.OptionQps(true)
		c.JSON(200, "OK")
	})
	router.GET("/disableqps", func(c *gin.Context) {
		logger.OptionQps(false)
		c.JSON(200, "OK")
	})

	// gr := router.Group("/")
	// gr.Use(jwtwrap.NewJWT().JWTAuth())

	apiRoute := router.Group("/api")

	userApi(apiRoute)

	router.Use(feMw("/")) //替换nginx serve 前端HTML代码
	logger.InitUserAndPathLimit(router)
}

func userApi(router *gin.RouterGroup) {

	// api := router.Group("/rbacapi")
	// api.Use(jwtwrap.NewJWT().JWTAuth())

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", my.Login)
	}
	router.GET("/version", softinfo.VersionInfo)
	/*
		myRoute := router.Group("/my")
		myRoute.Use(jwtwrap.NewJWT().JWTAuth())
		{
			myRoute.GET("/userinfo", my.UserInfo)
			myRoute.GET("/consume", my.UserConsumerHistory)
		}
	*/
}

func RegisterSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte("2abd2bd9663a799dcab3"))
	return sessions.Sessions("session", store)
}

//go:embed dist/*
var fs embed.FS

const fsBase = "dist" //和 embed一样

// feMw 使用go.16新的特性embed 到包前端编译后的代码. 替代nginx.   one binary rules them all
func feMw(urlPrefix string) gin.HandlerFunc {
	const indexHtml = "index.html"

	return func(c *gin.Context) {
		urlPath := strings.TrimSpace(c.Request.URL.Path)
		if urlPath == urlPrefix {
			urlPath = path.Join(urlPrefix, indexHtml)
		}
		urlPath = filepath.Join(fsBase, urlPath)
		fmt.Println(urlPath)
		fmt.Println(urlPath)
		fmt.Println(urlPath)

		f, err := fs.Open(urlPath)
		if err != nil {
			// 把 index.html 赋值过去
			urlPath = path.Join(urlPrefix, indexHtml)
			urlPath = filepath.Join(fsBase, urlPath)
			f, _ = fs.Open(urlPath)
			// return
		}
		fi, err := f.Stat()
		if strings.HasSuffix(urlPath, ".html") {
			c.Header("Cache-Control", "no-cache")
			c.Header("Content-Type", "text/html; charset=utf-8")
		}

		if strings.HasSuffix(urlPath, ".js") {
			c.Header("Content-Type", "text/javascript; charset=utf-8")
		}
		if strings.HasSuffix(urlPath, ".css") {
			c.Header("Content-Type", "text/css; charset=utf-8")
		}

		if err != nil || !fi.IsDir() {
			bs, err := fs.ReadFile(urlPath)
			if err != nil {
				logrus.WithError(err).Error("embed fs")
				return
			}
			c.Status(200)
			c.Writer.Write(bs)
			c.Abort()
		}
	}
}
