package wsproxy

import (
	"net/http"
	"net/url"
	"tokenagent/global"
	"tokenagent/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func EthMainNetReverseProxy(ctx *gin.Context) {
	target := global.MARKET_CONFIG.Eth.EthMain.WSUrl
	startProxy(target, ctx)
}

func EthPolygonReverseProxy(ctx *gin.Context) {
	target := global.MARKET_CONFIG.Eth.EthPolygon.WSUrl
	startProxy(target, ctx)
}

func EthRinkebyReverseProxy(ctx *gin.Context) {
	target := global.MARKET_CONFIG.Eth.EthRinkeby.WSUrl
	startProxy(target, ctx)
}

func EthPolygonmumbaiProxy(ctx *gin.Context) {
	target := global.MARKET_CONFIG.Eth.EthPolygonMumbai.WSUrl
	startProxy(target, ctx)
}

func startProxy(proxyURL string, ctx *gin.Context) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	uri, err := url.Parse(proxyURL)
	if err != nil {
		return
	}
	proxy := utils.NewProxy(uri)
	proxy.Upgrader = upgrader

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
