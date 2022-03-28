package nftservice

import (
	"fmt"
	"tokenagent/global"
	"tokenagent/model/model"
	"tokenagent/utils"
)

type NFT interface {
	NftPreMintCreate(args model.NftPreMintRequest) (*model.ContractCodeResponse, error)
	NftPreApproveCreate(args model.NftPreApproveRequest) (*model.ContractCodeResponse, error)
	NftPreBurnCreate(args model.NftPreBurnRequest) (*model.ContractCodeResponse, error)
	NftPreSetApprovalForAllCreate(args model.NftPreSetApprovalForAllRequest) (*model.ContractCodeResponse, error)
	NftPreTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error)
	NftPreSafeTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error)
	NftPreSafeTransferFromExCreate(args model.NftPreTransferFromRequestEx) (*model.ContractCodeResponse, error)
	NftPreTransferOwnershipCreate(args model.NftPreTransferOwnershipRequest) (*model.ContractCodeResponse, error)
	OwnerOf(args model.NftOwnerOfRequest) (string, error)
	BalanceOf(args model.NftBalanceOfRequest) (uint64, error)
	TokenUri(args model.NftTokenInfoRequest) (string, error)
	GetApproved(args model.NftTokenInfoRequest) (string, error)
}

type NFTProxy struct {
}

//@author: [libofeng]
//@function: NftPreMintCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreMintRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreMintCreate(args model.NftPreMintRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreMintCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreApproveCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreApproveRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreApproveCreate(args model.NftPreApproveRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreApproveCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreBurnCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreBurnRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreBurnCreate(args model.NftPreBurnRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreBurnCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreSetApprovalForAllCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreSetApprovalForAllRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreSetApprovalForAllCreate(args model.NftPreSetApprovalForAllRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreSetApprovalForAllCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreTransferFromCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreTransferFromRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreTransferFromCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (proxy *NFTProxy) NftPreSafeTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreSafeTransferFromCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func (proxy *NFTProxy) NftPreSafeTransferFromExCreate(args model.NftPreTransferFromRequestEx) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v`is invalid", args.BlockChain))
	}
	response, err := cli.NftPreSafeTransferFromExCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: NftPreTransferOwnershipCreate
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftPreTransferOwnershipRequest
//@return: *model.ContractCodeResponse
func (proxy *NFTProxy) NftPreTransferOwnershipCreate(args model.NftPreTransferOwnershipRequest) (*model.ContractCodeResponse, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return nil, global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.NftPreTransferOwnershipCreate(args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//@author: [libofeng]
//@function: OwnerOf
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftOwnerOfRequest
//@return: string
func (proxy *NFTProxy) OwnerOf(args model.NftOwnerOfRequest) (string, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.OwnerOf(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: BalanceOf
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftBalanceOfRequest
//@return: uint64
func (proxy *NFTProxy) BalanceOf(args model.NftBalanceOfRequest) (uint64, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return uint64(0), global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.BalanceOf(args)
	if err != nil {
		return uint64(0), err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenUri
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftTokenUriRequest
//@return: string
func (proxy *NFTProxy) TokenUri(args model.NftTokenInfoRequest) (string, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.TokenUri(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: TokenUri
//@description: 根据前端请求的BlockChain切换到相应链，处理RPC请求
//@param: args model.NftTokenUriRequest
//@return: string
func (proxy *NFTProxy) GetApproved(args model.NftTokenInfoRequest) (string, error) {
	cli := getNFTProxyInstance(args.BlockChain)
	bExit := utils.IsNil(cli)
	if bExit {
		return "", global.NewRPCError(global.InvalidParamsErrorCode, fmt.Sprintf("blockchain:`%v` is invalid", args.BlockChain))
	}
	response, err := cli.GetApproved(args)
	if err != nil {
		return "", err
	}
	return response, nil
}

//@author: [libofeng]
//@function: getNFTProxyInstance
//@description: 获取链对象
//@param: chain string
//@return: RPCClient接口对象
func getNFTProxyInstance(chain string) NFT {
	switch chain {
	case global.EthMain:
		return &ETHNft{}
	case global.EthPolygon:
		return &ETHNft{}
	case global.EthPolygonMumbai:
		return &ETHNft{}
	case global.EthRinkeby:
		return &ETHNft{}
	case global.XfsMain:
		return &XFSNft{}
	default:
		return nil
	}
}
