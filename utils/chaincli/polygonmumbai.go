package chaincli

import (
	"sync"
	"tokenagent/global"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
)

var golygonMumbaiIns *EthClient
var olygonMumbaiOnce sync.Once

func GetolygonMumbaiInstance() *EthClient {
	olygonMumbaiOnce.Do(func() {
		golygonMumbaiIns = NewolygonMumbaiClient()
	})
	return golygonMumbaiIns
}

//@author: [libofeng]
//@function: NewolygonMumbaiClient
//@description: 根据前端发送过来的链字典(ganache|rinkeby|...)找到配置文件tokenagent.yaml中的RPCURL建立客户端请求
//@param: chain string
//@return: EthClient对象
func NewolygonMumbaiClient() *EthClient {
	var err error
	cli := &EthClient{}
	cli.RpcCli, err = rpc.Dial(global.MARKET_CONFIG.Eth.EthPolygonMumbai.HttpUrl)
	if err != nil {
		global.MARKET_LOG.Error("PolygonMumbai client connect err...", zap.Error(err))
		return nil
	}
	cli.Client = ethclient.NewClient(cli.RpcCli)
	return cli
}
