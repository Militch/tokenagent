package chaincli

import (
	"context"
	"tokenagent/global"
	"tokenagent/model/model"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
)

type EthClient struct {
	*ethclient.Client
	RpcCli *rpc.Client
}

//@author: [libofeng]
//@function: TxPreExtra
//@description: 所有上链的操作需要将nonce,gasprice,gaslimit信息返给前端作为交易的参考
//@param: msg ethereum.CallMsg
//@return: *ethmodel.TxPreExtra
func (eth *EthClient) TxPreExtra(msg ethereum.CallMsg) (*model.TxPreExtra, error) {
	nonce, err := eth.PendingNonceAt(context.Background(), msg.From)
	if err != nil {
		global.MARKET_LOG.Debug("TxPreExtra->PendingNonceAt", zap.Error(err))
		return nil, err
	}
	gasPrice, err := eth.SuggestGasPrice(context.Background())
	if err != nil {
		global.MARKET_LOG.Debug("TxPreExtra->SuggestGasPrice", zap.Error(err))
		return nil, err
	}
	gas, err := eth.EstimateGas(context.Background(), msg)
	if err != nil {
		global.MARKET_LOG.Debug("TxPreExtra->EstimateGas", zap.Error(err))
		return nil, err
	}
	response := &model.TxPreExtra{
		GasPrice: gasPrice.Uint64(),
		GasLimit: gas,
		Nonce:    nonce,
	}
	return response, nil
}

func GetEthClient(chain string) *EthClient {
	switch chain {
	case global.EthMain:
		return GetEthMainInstance()
	case global.EthPolygon:
		return GetEthPolygonInstance()
	case global.EthRinkeby:
		return GetRinkebyInstance()
	case global.EthPolygonMumbai:
		return GetolygonMumbaiInstance()
	default:
		return nil
	}
}
