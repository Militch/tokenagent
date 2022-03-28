package router

import "tokenagent/core/router/ipfscli"

type RouterGroup struct {
	IPFSRouter ipfscli.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
