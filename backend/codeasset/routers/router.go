package routers

import (
	"backend/codeasset/common/middleware/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

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
