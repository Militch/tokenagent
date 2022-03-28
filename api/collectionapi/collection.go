package collectionapi

import (
	"tokenagent/model/model"
	"tokenagent/service"
)

type CollectionAPIHandler struct {
}

//@author: [libofeng]
//@function: PreCreate
//@description: 生成合集的合约码，返回给前端进行签名
//@param: args model.CollectionPreCreateRequest
//@return: resp **model.ContractCodeResponse

func (handler *CollectionAPIHandler) PreCreate(args model.CollectionPreCreateRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.CollectionProxy.CollectionPreCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreSetTokenURIPrefix
//@description: 生成设置NFTToken的URI前缀的code码，返回给前端进行签名
//@param: args model.CollectionPreSetURI
//@return: resp **model.ContractCodeResponse
func (handler *CollectionAPIHandler) PreSetTokenURIPrefix(args model.CollectionPreSetURI, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.CollectionProxy.CollectionPreSetTokenURIPrefixCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: Name
//@description: 获取合集的名称，直接调用无需签名
//@param: args model.CollectionRequest
//@return: resp *string

func (handler *CollectionAPIHandler) Name(args model.CollectionRequest, resp *string) error {
	response, err := service.ServiceGroupApp.CollectionProxy.Name(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: Symbol
//@description: 获取合集的符号，直接调用无需签名
//@param: args model.CollectionRequest
//@return: resp *string

func (handler *CollectionAPIHandler) Symbol(args model.CollectionRequest, resp *string) error {
	response, err := service.ServiceGroupApp.CollectionProxy.Symbol(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: TokenURIPrefix
//@description: 获取NFT Token的URL，直接调用无需签名
//@param: args model.CollectionRequest
//@return:  resp *string

func (handler *CollectionAPIHandler) TokenURIPrefix(args model.CollectionRequest, resp *string) error {
	response, err := service.ServiceGroupApp.CollectionProxy.TokenURIPrefix(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: Owner
//@description: 获取合集的拥有者，直接调用无需签名
//@param: args model.CollectionRequest
//@return:  resp *string

func (handler *CollectionAPIHandler) Owner(args model.CollectionRequest, resp *string) error {
	response, err := service.ServiceGroupApp.CollectionProxy.Owner(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: TotalSupply
//@description: 获取合集下Token的总额，直接调用无需签名
//@param: args model.CollectionRequest
//@return:  resp *uint64

func (handler *CollectionAPIHandler) TotalSupply(args model.CollectionRequest, resp *uint64) error {
	response, err := service.ServiceGroupApp.CollectionProxy.TotalSupply(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: CreateWithPrikey
//@description: 前端将私钥发过来，直接签名交易上链
//@param: args model.CollectionRequest
//@return:  resp **model.TxResponse)

func (handler *CollectionAPIHandler) CreateWithPrikey(args model.CollectionCreateWithPrikeyRequest, resp **model.TxResponse) error {
	response, err := service.ServiceGroupApp.CollectionProxy.CollectionCreateWithPrikey(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}
