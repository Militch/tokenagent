package chainservice

import (
	"fmt"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils"
)

type Chain interface {
	// 发送签名的交易上链
	SendRawTransaction(args model.TxRequest) (*model.TxResponse, error)
	// 获取交易回执
	GetTransactionReceipt(args model.TxRecRequest) (*model.Receipt, error)
}

type ChainProxy struct {
}

//@author: [libofeng]
//@function: SendRawTransaction
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TxRequest
//@return: resp **model.TxResponse

func (proxy *ChainProxy) SendRawTransaction(args model.TxRequest) (*model.TxResponse, error) {
	cli := getChainProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.SendRawTransaction(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: SendRawTransaction
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TxRequest
//@return: resp **model.TxResponse
func (proxy *ChainProxy) GetTransactionReceipt(args model.TxRecRequest) (*model.Receipt, error) {
	cli := getChainProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.GetTransactionReceipt(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: getChainProxyInstance
//@description: 获取链对象
//@param: chain string
//@return: RPCClient接口对象
func getChainProxyInstance(chain string) Chain {
	switch chain {
	case global.EthMain:
		return &ETHCommand{}
	case global.EthPolygon:
		return &ETHCommand{}
	case global.EthPolygonMumbai:
		return &ETHCommand{}
	case global.EthRinkeby:
		return &ETHCommand{}
	case global.XfsMain:
		return &XFSCommand{}
	default:
		return nil
	}
}
