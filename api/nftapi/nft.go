package nftapi

import (
	"tokenagent/model/model"
	"tokenagent/service"
)

type NftAPIHandler struct {
}

//@author: [libofeng]
//@function: PreMint
//@description: 生成铸造NFT的code码，返回给前端进行签名
//@param: args model.NftPreMintRequest
//@return: resp **model.ContractCodeResponse

func (handler *NftAPIHandler) PreMint(args model.NftPreMintRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreMintCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreApprove
//@description: 生成NFT授权的code码，返回给前端进行签名
//@param: args model.NftPreApproveRequest
//@return: resp **model.ContractCodeResponse

func (handler *NftAPIHandler) PreApprove(args model.NftPreApproveRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreApproveCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreBurn
//@description: 生成注销NFT的code码，返回给前端进行签名
//@param: args model.NftPreBurnRequest
//@return: resp **model.ContractCodeResponse

func (handler *NftAPIHandler) PreBurn(args model.NftPreBurnRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreBurnCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

func (handler *NftAPIHandler) PreSetApprovalForAll(args model.NftPreSetApprovalForAllRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreSetApprovalForAllCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreTransferFrom
//@description: 生成转移NFTToken的code码，返回给前端进行签名
//@param: args model.NftPreTransferFromRequest
//@return: resp **model.ContractCodeResponse
func (handler *NftAPIHandler) PreTransferFrom(args model.NftPreTransferFromRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreTransferFromCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreSafeTransferFrom
//@description: 生成转移NFTToken的code码，返回给前端进行签名
//@param: args model.NftPreTransferFromRequest
//@return: resp **model.ContractCodeResponse
func (handler *NftAPIHandler) PreSafeTransferFrom(args model.NftPreTransferFromRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreSafeTransferFromCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreSafeTransferFromEx
//@description: 生成转移NFTToken的code码，返回给前端进行签名
//@param: args model.NftPreTransferFromRequestEx
//@return: resp **model.ContractCodeResponse
func (handler *NftAPIHandler) PreSafeTransferFromEx(args model.NftPreTransferFromRequestEx, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreSafeTransferFromExCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreTransferOwnership
//@description: 生成转移NFTToken的拥有权code码，返回给前端进行签名
//@param: args model.NftPreTransferFromRequest
//@return: resp **model.ContractCodeResponse
func (handler *NftAPIHandler) PreTransferOwnership(args model.NftPreTransferOwnershipRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.NFTProxy.NftPreTransferOwnershipCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: OwnerOf
//@description: 返回NFT Token的拥有者，直接调用，无需签名
//@param: args model.NftOwnerOfRequest
//@return: resp *string
func (handler *NftAPIHandler) OwnerOf(args model.NftOwnerOfRequest, resp *string) error {
	response, err := service.ServiceGroupApp.NFTProxy.OwnerOf(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: BalanceOf
//@description: 返回NFT拥有者的余额，返回给前端进行签名
//@param: args model.NftBalanceOfRequest
//@return: resp *uint64
func (handler *NftAPIHandler) BalanceOf(args model.NftBalanceOfRequest, resp *uint64) error {
	response, err := service.ServiceGroupApp.NFTProxy.BalanceOf(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: TokenUri
//@description: 返回NFTToken的URI，返回给前端进行签名
//@param: args model.NftTokenUriRequest
//@return: resp *string
func (handler *NftAPIHandler) TokenUri(args model.NftTokenInfoRequest, resp *string) error {
	response, err := service.ServiceGroupApp.NFTProxy.TokenUri(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: GetApproved
//@description: 返回NFTToken的URI，返回给前端进行签名
//@param: args model.NftTokenUriRequest
//@return: resp *string
func (handler *NftAPIHandler) GetApproved(args model.NftTokenInfoRequest, resp *string) error {
	response, err := service.ServiceGroupApp.NFTProxy.GetApproved(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}
