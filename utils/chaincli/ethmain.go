package chaincli

import (
	"sync"
	"tokenagent/global"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
)

var ethMainIns *EthClient
var ethMainOnce sync.Once

func GetEthMainInstance() *EthClient {
	ethMainOnce.Do(func() {
		ethMainIns = NewEthMainClient()
	})
	return ethMainIns
}

//@author: [libofeng]
//@function: NewClient
//@description: 根据前端发送过来的链字典(ganache|rinkeby|...)找到配置文件tokenagent.yaml中的RPCURL建立客户端请求
//@param: chain string
//@return: EthClient对象
func NewEthMainClient() *EthClient {
	var err error
	cli := &EthClient{}
	cli.RpcCli, err = rpc.Dial(global.MARKET_CONFIG.Eth.EthMain.HttpUrl)
	if err != nil {
		global.MARKET_LOG.Error("ethmainnet client connect err...", zap.Error(err))
		return nil
	}
	cli.Client = ethclient.NewClient(cli.RpcCli)
	return cli
}
