package collectionservice

import (
	"fmt"
	"math/big"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils"
)

type Collection interface {
	// 生成合集的合约码，返回给前端进行签名
	CollectionPreCreate(args model.CollectionPreCreateRequest) (*model.ContractCodeResponse, error)
	// 修改Collection的Token的URI前缀
	CollectionPreSetTokenURIPrefixCreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error)
	// 获取合集的名称，直接调用无需签名
	Name(args model.CollectionRequest) (string, error)
	//  获取合集的符号，直接调用无需签名
	Symbol(args model.CollectionRequest) (string, error)
	// 获取NFT Token的URL前缀，直接调用无需签名
	TokenURIPrefix(args model.CollectionRequest) (string, error)
	// 获取合集的拥有者，直接调用无需签名
	Owner(args model.CollectionRequest) (string, error)
	// 获取合集下Token的总额，直接调用无需签名
	TotalSupply(args model.CollectionRequest) (*big.Int, error)
	// 前端将私钥发过来，直接签名交易上链
	CollectionCreateWithPrikey(args model.CollectionCreateWithPrikeyRequest) (*model.TxResponse, error)
}

type CollectionProxy struct {
}

//@author: [libofeng]
//@function: CollectionPreCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionPreCreateRequest
//@return: *model.ContractCodeResponse

func (proxy *CollectionProxy) CollectionPreCreate(args model.CollectionPreCreateRequest) (*model.ContractCodeResponse, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.CollectionPreCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreSetTokenURIPrefixCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionPreSetURI
//@return: *model.ContractCodeResponse
func (proxy *CollectionProxy) CollectionPreSetTokenURIPrefixCreate(args model.CollectionPreSetURI) (*model.ContractCodeResponse, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.CollectionPreSetTokenURIPrefixCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: Name
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionRequest
//@return: string
func (proxy *CollectionProxy) Name(args model.CollectionRequest) (string, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.Name(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: Symbol
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionRequest
//@return: string
func (proxy *CollectionProxy) Symbol(args model.CollectionRequest) (string, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.Symbol(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenURIPrefix
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionRequest
//@return: string
func (proxy *CollectionProxy) TokenURIPrefix(args model.CollectionRequest) (string, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenURIPrefix(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: Owner
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionRequest
//@return: string
func (proxy *CollectionProxy) Owner(args model.CollectionRequest) (string, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.Owner(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TotalSupply
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionRequest
//@return: uint64
func (proxy *CollectionProxy) TotalSupply(args model.CollectionRequest) (*big.Int, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return *new(*big.Int), global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TotalSupply(args)
	if err != nil {
		return *new(*big.Int), err
	}
	return response, nil
}

//@author: [libofeng]
//@function: CollectionCreateWithPrikey
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.CollectionCreateWithPrikeyRequest
//@return: *model.TxResponse

func (proxy *CollectionProxy) CollectionCreateWithPrikey(args model.CollectionCreateWithPrikeyRequest) (*model.TxResponse, error) {
	cli := getCollectionProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.CollectionCreateWithPrikey(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: getCollectionProxyInstance
//@description: 获取链对象
//@param: chain string
//@return: RPCClient接口对象
func getCollectionProxyInstance(chain string) Collection {
	switch chain {
	case global.EthMain:
		return &ETHCollection{}
	case global.EthPolygon:
		return &ETHCollection{}
	case global.EthPolygonMumbai:
		return &ETHCollection{}
	case global.EthRinkeby:
		return &ETHCollection{}
	case global.XfsMain:
		return &XFSCollection{}
	default:
		return nil
	}
}
