package initialize

import (
	"tokenagent/core/middleware"
	"tokenagent/core/router"
	"tokenagent/global"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	gin.SetMode("release")
	Router.Use(gin.Recovery())
	Router.Use(middleware.CorsByRules())
	ipfs := router.RouterGroupApp.IPFSRouter
	PublibGroup := Router.Group("")
	{
		ipfs.InitIPFSRouter(PublibGroup)
	}

	global.MARKET_LOG.Info("router register success")
	return Router
}
