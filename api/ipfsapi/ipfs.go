package ipfsapi

import (
	"tokenagent/global"
	"tokenagent/model"
	"tokenagent/service/ipfsservice"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IPFSAPIHandler struct {
	ipfsService *ipfsservice.FileUploadAndDownloadService
}

//@author: [libofeng]
//@function: Publish2IPFS
//@description: 上传文件到IPFS网络
//@param: form-data
//@return: 文件hash

func (ipfs *IPFSAPIHandler) Publish2IPFS(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.MARKET_LOG.Error("接收文件失败!", zap.Error(err))
		return
	}
	id, err := ipfs.ipfsService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		global.MARKET_LOG.Error("发布文件到ipfs!", zap.Error(err))
		model.FailWithMessage(err.Error(), ctx)
		return
	}
	// model.OkWithDetailed(gin.H{"ipfsurl": global.MARKET_CONFIG.IpfsOSS.Path + url}, "发布成功!", ctx)
	model.OkWithDetailed(gin.H{"ipfs": id}, "发布成功!", ctx)
}
