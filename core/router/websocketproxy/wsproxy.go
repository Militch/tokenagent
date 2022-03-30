package websocketproxy

import (
	"tokenagent/service/wsproxy"

	"github.com/gin-gonic/gin"
)

type WSProxyRouter struct {
}

func (proxy *WSProxyRouter) WSRouter(Router *gin.RouterGroup) {
	Router.GET("/ethmainnet", wsproxy.EthMainNetReverseProxy)
	Router.GET("/ethpolygon", wsproxy.EthPolygonReverseProxy)
	Router.GET("/ethpolygonmumbai", wsproxy.EthPolygonmumbaiProxy)
	Router.GET("/ethrinkeby", wsproxy.EthRinkebyReverseProxy)

}
