package ipfscli

import (
	"tokenagent/api/ipfsapi"

	"github.com/gin-gonic/gin"
)

type IPFSCliRouter struct {
}

func (ipfs *IPFSCliRouter) InitIPFSRouter(Router *gin.RouterGroup) {
	ipfsRouter := Router.Group("api")
	ipfsApi := &ipfsapi.IPFSAPIHandler{}
	{
		ipfsRouter.POST("publish2ipfs", ipfsApi.Publish2IPFS)

	}
}
