package ipfsservice

import (
	"errors"
	"mime/multipart"
	"tokenagent/global"

	ipfsshell "github.com/ipfs/go-ipfs-api"
	"go.uber.org/zap"
)

func (*FileUploadAndDownloadService) UploadFile(file *multipart.FileHeader) (string, error) {
	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.MARKET_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	shell := ipfsshell.NewShell(global.MARKET_CONFIG.IpfsOSS.EndPoint)

	cid, err := shell.Add(f)
	if err != nil {
		global.MARKET_LOG.Error("publish file to ipfs Filed", zap.Any("err", err.Error()))
		return "", errors.New("publish file to ipfs Filed, err:" + err.Error())
	}

	return cid, nil
}
