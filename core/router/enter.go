package router

import (
	"tokenagent/core/router/ipfscli"
	"tokenagent/core/router/websocketproxy"
)

type RouterGroup struct {
	IPFSRouter    ipfscli.RouterGroup
	WSProxyRouter websocketproxy.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
