package tokenservice

import (
	"fmt"
	"math/big"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils"
)

type Token interface {
	TokenPreCreate(args model.TokenPreCreateRequest) (*model.ContractCodeResponse, error)
	Name(args model.TokenCallRequest) (string, error)
	Symbol(args model.TokenCallRequest) (string, error)
	TotalSupply(args model.TokenCallRequest) (*big.Int, error)
	BalanceOf(args model.TokenBalanceOfCallRequest) (*big.Int, error)
	TokenPreMintCreate(args model.TokenPreMintRequest) (*model.ContractCodeResponse, error)
	TokenPreApproveCreate(args model.TokenPreApproveRequest) (*model.ContractCodeResponse, error)
	TokenPreBurnFromCreate(args model.TokenPreBurnRequest) (*model.ContractCodeResponse, error)
	TokenPreIncreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error)
	TokenPreDecreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error)
	TokenPreTransferFromCreate(args model.TokenPreTransferRequest) (*model.ContractCodeResponse, error)
	TokenPrePauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error)
	TokenPreUnPauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error)
}

type TokenProxy struct {
}

//@author: [libofeng]
//@function: TokenPreCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreCreateRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreCreate(args model.TokenPreCreateRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: Name
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenRequest
//@return: string
func (proxy *TokenProxy) Name(args model.TokenCallRequest) (string, error) {
	cli := getTokenProxyInstance(args.BlockChain)
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
//@param: args model.TokenRequest
//@return: string
func (proxy *TokenProxy) Symbol(args model.TokenCallRequest) (string, error) {
	cli := getTokenProxyInstance(args.BlockChain)
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
//@function: TotalSupply
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenRequest
//@return: *big.Int
func (proxy *TokenProxy) TotalSupply(args model.TokenCallRequest) (*big.Int, error) {
	cli := getTokenProxyInstance(args.BlockChain)
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
//@function: BalanceOf
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenRequest
//@return: *big.Int
func (proxy *TokenProxy) BalanceOf(args model.TokenBalanceOfCallRequest) (*big.Int, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return *new(*big.Int), global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.BalanceOf(args)
	if err != nil {
		return *new(*big.Int), err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreMintCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreMintRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreMintCreate(args model.TokenPreMintRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreMintCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreApproveCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreApproveRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreApproveCreate(args model.TokenPreApproveRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreApproveCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreMintCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreMintRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreBurnFromCreate(args model.TokenPreBurnRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreBurnFromCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreIncreaseAllowance
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreAllowanceRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreIncreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreIncreaseAllowance(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreDecreaseAllowance
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreAllowanceRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreDecreaseAllowance(args model.TokenPreAllowanceRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreDecreaseAllowance(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreSafeTransferFromCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreTransferRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreTransferFromCreate(args model.TokenPreTransferRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreTransferFromCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPrePauseCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPrePauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPrePauseCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenPreUnPauseCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.TokenPreRequest
//@return: *model.ContractCodeResponse
func (proxy *TokenProxy) TokenPreUnPauseCreate(args model.TokenPreRequest) (*model.ContractCodeResponse, error) {
	cli := getTokenProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenPreUnPauseCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: getTokenProxyInstance
//@description: 获取链对象
//@param: chain string
//@return: RPCClient接口对象
func getTokenProxyInstance(chain string) Token {
	switch chain {
	case global.EthMain:
		return &ETHToken{}
	case global.EthPolygon:
		return &ETHToken{}
	case global.EthPolygonMumbai:
		return &ETHToken{}
	case global.EthRinkeby:
		return &ETHToken{}
	case global.XfsMain:
		return &ETHToken{}
	default:
		return nil
	}
}
